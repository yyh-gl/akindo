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

func TestOANDA(t *testing.T) {
	t.Run("OANDAクライアントを取得できる", func(t *testing.T) {
		t.Run("HTTPクライアントを渡さずにOANDAクライアントを取得できる", func(t *testing.T) {
			a, err := newOANDAClient(nil, accessToken, accountID)

			assert.NotNil(t, a)
			assert.Nil(t, err)
		})

		t.Run("HTTPクライアントを渡してOANDAクライアントを取得できる", func(t *testing.T) {
			a, err := newOANDAClient(http.DefaultClient, accessToken, accountID)

			assert.NotNil(t, a)
			assert.Nil(t, err)
		})
	})

	t.Run("アカウント情報を取得できる", func(t *testing.T) {
		a, _ := newOANDAClient(nil, accessToken, accountID)
		ac, err := a.getAccount(context.TODO())

		assert.Equal(t, `{"accounts":[{"id":"`+accountID+`","tags":[]}]}`, ac)
		assert.Nil(t, err)
	})

	t.Run("ローソク足情報を取得できる", func(t *testing.T) {
		a, _ := newOANDAClient(nil, accessToken, accountID)
		cs, err := a.getCandles(context.TODO(), "USD_JPY")

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
