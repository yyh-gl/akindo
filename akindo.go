package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type AkindoClient struct{}

func NewAkindoClient() *AkindoClient {
	return &AkindoClient{}
}

func (ac AkindoClient) GetAccount() (string, error) {
	url := "https://api-fxpractice.oanda.com/v3/accounts"
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("リクエスト作成に失敗: %w", err)
	}

	accessToken := os.Getenv("API_ACCESS_TOKEN")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	c := http.DefaultClient
	resp, err := c.Do(req)
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

func (ac AkindoClient) GetCandle() (string, error) {
	url := "https://api-fxpractice.oanda.com/v3/accounts/101-009-15951441-001/candles/latest?candleSpecifications=USD_JPY:M1:BM"
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("リクエスト作成に失敗: %w", err)
	}

	accessToken := os.Getenv("API_ACCESS_TOKEN")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	c := http.DefaultClient
	resp, err := c.Do(req)
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
