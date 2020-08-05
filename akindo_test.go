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

	t.Run("ろうそく足情報を取得できる", func(t *testing.T) {
		a, _ := NewAkindo(nil, accessToken, accountID)
		c, err := a.GetCandle(context.TODO())

		// ろうそく足情報はリアルタイムに変わっていくので、空でないことを確認
		assert.NotEmpty(t, c)
		assert.Nil(t, err)
	})
}
