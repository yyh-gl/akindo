package main

import (
	"context"
	"net/http"
	"os"
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

	t.Run("ローソク足情報を取得できる", func(t *testing.T) {
		a, _ := NewAkindo(nil, accessToken, accountID)
		cs, err := a.GetCandles(context.TODO(), "USD_JPY")

		// ローソク足情報はリアルタイムに変わっていくので、
		// エラーが出ていないことおよび各フィールドが空でないことだけを確認
		assert.Nil(t, err)
		for _, c := range *cs {
			//assert.NotEmpty(t, c.Complete) // Completeにはfalse（ゼロ値）が入ることがあるため確認しない
			assert.NotEmpty(t, c.Volume)
			assert.NotEmpty(t, c.Time)
			assert.NotEmpty(t, c.Prices.Open)
			assert.NotEmpty(t, c.Prices.Highest)
			assert.NotEmpty(t, c.Prices.Lowest)
			assert.NotEmpty(t, c.Prices.Closing)
		}
	})
}
