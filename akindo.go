package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const host = "https://api-fxpractice.oanda.com/v3"

// Akindo : Akindoクライアントの構造体
type Akindo struct {
	httpClient  *http.Client
	accessToken string
	accountID   string
}

// NewAkindo : Akindoクライアントを生成
func NewAkindo(httpClient *http.Client, accessToken, accountID string) (*Akindo, error) {
	hc := http.DefaultClient
	if httpClient != nil {
		hc = httpClient
	}

	return &Akindo{
		httpClient:  hc,
		accessToken: accessToken,
		accountID:   accountID,
	}, nil
}

// sendRequest : OANDA APIへのリクエスト用共通処理
func (a Akindo) sendRequest(ctx context.Context, path string) (*http.Response, error) {
	url := host + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成に失敗: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+a.accessToken)

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("リクエスト実行に失敗: %w", err)
	}
	return resp, nil
}

// GetAccount : アカウント情報を取得
func (a Akindo) GetAccount(ctx context.Context) (string, error) {
	resp, err := a.sendRequest(ctx, "/accounts")
	if err != nil {
		return "", fmt.Errorf("リクエスト実行に失敗: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("レスポンス読込に失敗: %w", err)
	}
	return string(b), nil
}

type (
	// CandleStick : ローソク足を表す構造体
	// https://developer.oanda.com/rest-live-v20/instrument-df/#Candlestick
	CandleStick struct {
		Complete bool      `json:"complete"`
		Volume   int       `json:"volume"`
		Time     time.Time `json:"time"`
		Prices   Prices    `json:"mid"`
	}

	// Prices : ローソク足情報内の価格情報を表す構造体
	Prices struct {
		Open    string `json:"o"`
		Highest string `json:"h"`
		Lowest  string `json:"l"`
		Closing string `json:"c"`
	}

	// CandleSticks : ローソク足の集合を表す構造体
	CandleSticks []*CandleStick
)

// GetCandle : ローソク足情報を取得
func (a Akindo) GetCandles(ctx context.Context, instrument string) (*CandleSticks, error) {
	type response struct {
		Instrument  string       `json:"instrument"`
		Granularity string       `json:"granularity"`
		Candles     CandleSticks `json:"candles"`
	}

	path := fmt.Sprintf("/accounts/%s/instruments/%s/candles", a.accountID, instrument)
	resp, err := a.sendRequest(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("リクエスト実行に失敗: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンス読込に失敗: %w", err)
	}

	var r response
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, fmt.Errorf("json.Unmarshal()に失敗: %w", err)
	}
	return &r.Candles, nil
}
