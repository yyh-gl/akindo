package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAkindo(t *testing.T) {
	t.Run("アカウント情報を取得できること", func(t *testing.T) {
		a := NewAkindo(nil, os.Getenv("API_ACCESS_TOKEN"))
		ac, err := a.GetAccount(context.TODO())

		assert.Equal(t, `{"accounts":[{"id":"101-009-15951441-001","tags":[]}]}`, ac)
		assert.Nil(t, err)
	})

	t.Run("ろうそく足情報を取得できること", func(t *testing.T) {
		a := NewAkindo(nil, os.Getenv("API_ACCESS_TOKEN"))
		c, err := a.GetCandle(context.TODO())

		// ろうそく足情報はリアルタイムに変わっていくので、空でないことを確認
		assert.NotEmpty(t, c)
		assert.Nil(t, err)
	})
}
