package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var accessToken = os.Getenv("API_ACCESS_TOKEN")
var accountID = os.Getenv("ACCOUNT_ID")

func TestAkindo(t *testing.T) {
	t.Run("アカウント情報を取得できる", func(t *testing.T) {
		a, _ := NewAkindo(nil, accessToken, accountID)
		ac, err := a.GetAccount(context.TODO())

		assert.Equal(t, `{"accounts":[{"id":"`+accountID+`","tags":[]}]}`, ac)
		assert.Nil(t, err)
	})

	t.Run("ろうそく足情報を取得できる", func(t *testing.T) {
		a, _ := NewAkindo(nil, accessToken, accountID)
		c, err := a.GetCandle(context.TODO())

		// ろうそく足情報はリアルタイムに変わっていくので、空でないことを確認
		assert.NotEmpty(t, c)
		assert.Nil(t, err)
	})
}
