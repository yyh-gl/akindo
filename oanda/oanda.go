package oanda

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
func (a Client) sendRequest(ctx context.Context, path string) (*http.Response, error) {
	url := host + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(ctx, %s, %s, nil) > %w", http.MethodGet, url, err)
	}
	req.Header.Add("Authorization", "Bearer "+a.accessToken)

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Client.Do(%+v) > %w", req, err)
	}
	return resp, nil
}

// getAccount : アカウント情報を取得
func (a Client) getAccount(ctx context.Context) (string, error) {
	resp, err := a.sendRequest(ctx, "/accounts")
	if err != nil {
		return "", fmt.Errorf("sendRequest(ctx, \"/accounts\") > %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll() > %w", err)
	}
	return string(b), nil
}

type (
	// candleStick : ローソク足を表す構造体
	// https://developer.oanda.com/rest-live-v20/instrument-df/#Candlestick
	candleStick struct {
		Complete bool      `json:"complete"`
		Volume   int       `json:"volume"`
		Time     time.Time `json:"time"`
		Prices   prices    `json:"mid"`
	}

	// prices : ローソク足情報内の価格情報を表す構造体
	prices struct {
		Open    string `json:"o"`
		Highest string `json:"h"`
		Lowest  string `json:"l"`
		Closing string `json:"c"`
	}

	// candleSticks : ローソク足の集合を表す構造体
	candleSticks []*candleStick
)

// getCandle : ローソク足情報を取得
func (a Client) getCandles(ctx context.Context, instrument string) (*candleSticks, error) {
	type response struct {
		Instrument  string       `json:"instrument"`
		Granularity string       `json:"granularity"`
		Candles     candleSticks `json:"candles"`
	}

	path := fmt.Sprintf("/accounts/%s/instruments/%s/candles", a.accountID, instrument)
	resp, err := a.sendRequest(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("sendRequest(ctx, \"%s\") > %w", path, err)
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
	return &r.Candles, nil
}
