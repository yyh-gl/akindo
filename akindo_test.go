package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var accessToken = os.Getenv("API_ACCESS_TOKEN")
var accountID = os.Getenv("ACCOUNT_ID")

func TestAkindo(t *testing.T) {
	t.Run("Akindoクライアントを取得できる", func(t *testing.T) {
		t.Run("HTTPクライアントを渡さずにAkindoクライアントを取得できる", func(t *testing.T) {
			a, err := NewAkindo(nil, accessToken, accountID)

			assert.NotNil(t, a)
			assert.Nil(t, err)
		})

		t.Run("HTTPクライアントを渡してAkindoクライアントを取得できる", func(t *testing.T) {
			a, err := NewAkindo(http.DefaultClient, accessToken, accountID)

			assert.NotNil(t, a)
			assert.Nil(t, err)
		})
	})

	t.Run("アカウント情報を取得できる", func(t *testing.T) {
		a, _ := NewAkindo(nil, accessToken, accountID)
		ac, err := a.GetAccount(context.TODO())

		assert.Equal(t, `{"accounts":[{"id":"`+accountID+`","tags":[]}]}`, ac)
		assert.Nil(t, err)
	})

	t.Run("ろうそく足情報を取得できる", func(t *testing.T) {
		a, _ := NewAkindo(nil, accessToken, accountID)
		c, err := a.GetCandles(context.TODO(), "USD_JPY")

		// ろうそく足情報はリアルタイムに変わっていくので、レスポンスの一部分だけを比較
		assert.True(t, strings.Contains(c, `{"instrument":"USD_JPY","granularity":"S5","candles"`))
		assert.Nil(t, err)
	})
}
