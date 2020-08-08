package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOANDA(t *testing.T) {
	at := os.Getenv("API_ACCESS_TOKEN")
	id := os.Getenv("ACCOUNT_ID")

	t.Run("OANDAクライアントを取得できる", func(t *testing.T) {
		oc := newOANDAClient(at, id)

		assert.NotNil(t, oc)
	})

	t.Run("アカウント情報を取得できる", func(t *testing.T) {
		oc := newOANDAClient(at, id)
		ac, err := oc.getAccount(context.TODO())

		assert.Equal(t, `{"accounts":[{"id":"`+id+`","tags":[]}]}`, ac)
		assert.Nil(t, err)
	})

	t.Run("ローソク足情報を取得できる", func(t *testing.T) {
		oc := newOANDAClient(at, id)
		cs, err := oc.getCandles(context.TODO(), "USD_JPY")

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
