package oanda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/yyh-gl/fx-auto-trader/candle"
)

const host = "https://api-fxpractice.oanda.com/v3"

// Client : OANDAクライアントの構造体
type Client struct {
	httpClient  *http.Client
	accessToken string
	accountID   string
}

// NewClient : OANDAクライアントを生成
func NewClient(accessToken, accountID string) *Client {
	return &Client{
		httpClient:  http.DefaultClient,
		accessToken: accessToken,
		accountID:   accountID,
	}
}

// sendRequest : OANDA APIへのリクエスト用共通処理
func (c *Client) sendRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var b []byte
	if body != nil {
		var err error
		b, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal(Body: %+v) > %w", body, err)
		}
	}

	url := host + path
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(ctx, %s, %s, nil) > %w", http.MethodGet, url, err)
	}
	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Client.Do(%+v) > %w", req, err)
	}
	return resp, nil
}

// getAccount : アカウント情報を取得
func (c *Client) getAccount(ctx context.Context) (string, error) {
	resp, err := c.sendRequest(ctx, http.MethodGet, "/accounts", nil)
	if err != nil {
		return "", fmt.Errorf("sendRequest(Path: /accounts, Body: nil) > %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll() > %w", err)
	}
	return string(b), nil
}

type (
	// candleStickDTO : OANDA APIからもらったローソク足情報を受け取るためのDTO
	// https://developer.oanda.com/rest-live-v20/instrument-df/#Candlestick
	candleStickDTO struct {
		Complete bool      `json:"complete"`
		Volume   int       `json:"volume"`
		Time     time.Time `json:"time"`
		Prices   struct {
			Open    string `json:"o"`
			Highest string `json:"h"`
			Lowest  string `json:"l"`
			Closing string `json:"c"`
		} `json:"mid"`
	}

	// candleSticksDTO : OANDA APIからもらったローソク足情報の集合を受け取るためのDTO
	candleSticksDTO []*candleStickDTO
)

// getCandle : ローソク足情報を取得
func (c *Client) getCandles(ctx context.Context, instrument string) (candle.CandleSticks, error) {
	type response struct {
		Instrument  string          `json:"instrument"`
		Granularity string          `json:"granularity"`
		Candles     candleSticksDTO `json:"candles"`
	}

	path := fmt.Sprintf("/accounts/%s/instruments/%s/candles?granularity=M2", c.accountID, instrument)
	resp, err := c.sendRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("sendRequest(Path: %s, Body: nil) > %w", path, err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll() > %w", err)
	}

	var r response
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s, response{}) > %w", string(b), err)
	}

	var cs candle.CandleSticks
	for _, c := range r.Candles {
		o, err := strconv.ParseFloat(c.Prices.Open, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv.ParseFloat(OpenPrice: %s) > %w", c.Prices.Open, err)
		}
		h, err := strconv.ParseFloat(c.Prices.Highest, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv.ParseFloat(HighestPrice: %s) > %w", c.Prices.Highest, err)
		}
		l, err := strconv.ParseFloat(c.Prices.Lowest, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv.ParseFloat(LowestPrice: %s) > %w", c.Prices.Lowest, err)
		}
		cl, err := strconv.ParseFloat(c.Prices.Closing, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv.ParseFloat(ClosingPrice: %s) > %w", c.Prices.Closing, err)
		}

		cs = append(cs, &candle.CandleStick{
			Complete: c.Complete,
			Open:     o,
			Highest:  h,
			Lowest:   l,
			Closing:  cl,
		})
	}
	return cs, nil
}

// getLatestCandle : 最新のローソク足情報を1件取得
func (c *Client) GetLatestCandle(ctx context.Context, instrument string) (*candle.CandleStick, error) {
	cs, err := c.getCandles(ctx, instrument)
	if err != nil {
		return nil, fmt.Errorf("getCandles(Instrument: %s) > %w", instrument, err)
	}

	// 最後に取得したローソク足情報がfixしている場合はその情報を返却
	// fixしていない場合は2つ前の情報を返却
	if latest := cs[len(cs)-1]; latest.Complete {
		return latest, nil
	}
	return cs[len(cs)-2], nil
}

// buy : 外貨を購入
func (c *Client) Buy(ctx context.Context, instrument string, units int) error {
	type request struct {
		Order struct {
			Units        string `json:"units"`
			Instrument   string `json:"instrument"`
			Type         string `json:"type"`
			PositionFill string `json:"positionFill"`
		} `json:"order"`
	}
	type response struct {
		OrderCreateTransaction struct {
			AccountID    string    `json:"accountID"`
			BatchID      string    `json:"batchID"`
			ID           string    `json:"id"`
			Instrument   string    `json:"instrument"`
			PositionFill string    `json:"positionFill"`
			Reason       string    `json:"reason"`
			Time         time.Time `json:"time"`
			TimeInForce  string    `json:"timeInForce"`
			Type         string    `json:"type"`
			Units        string    `json:"units"`
			UserID       string    `json:"userID"`
		} `json:"orderCreateTransaction"`
		ErrorMessage string `json:"errorMessage"`
	}

	req := request{
		Order: struct {
			Units        string `json:"units"`
			Instrument   string `json:"instrument"`
			Type         string `json:"type"`
			PositionFill string `json:"positionFill"`
		}{
			Units:        strconv.Itoa(units),
			Instrument:   instrument,
			Type:         "MARKET",
			PositionFill: "DEFAULT",
		},
	}
	path := fmt.Sprintf("/accounts/%s/orders", c.accountID)
	resp, err := c.sendRequest(ctx, http.MethodPost, path, &req)
	if err != nil {
		return fmt.Errorf("sendRequest(Path: %s, Body: %+v) > %w", path, req, err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll() > %w", err)
	}

	var r response
	if err := json.Unmarshal(b, &r); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, response{}) > %w", string(b), err)
	}

	if r.ErrorMessage != "" {
		return fmt.Errorf("failed to request > %s", r.ErrorMessage)
	}
	return nil
}

// sell : 外貨を売却
func (c *Client) Sell(ctx context.Context, instrument string, units int) error {
	return c.Buy(ctx, instrument, -units)
}
