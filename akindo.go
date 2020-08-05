package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
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

// GetCandle : ろうそく足情報を取得
func (a Akindo) GetCandle(ctx context.Context) (string, error) {
	path := fmt.Sprintf("accounts/%s/candles/latest?candleSpecifications=USD_JPY:M1:BM", a.accountID)
	resp, err := a.sendRequest(ctx, path)
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
