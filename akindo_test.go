package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAkindo(t *testing.T) {
	t.Run("アカウント情報を取得できること", func(t *testing.T) {
		ac := NewAkindoClient()
		a, err := ac.GetAccount()

		assert.Equal(t, `{"accounts":[{"id":"101-009-15951441-001","tags":[]}]}`, a)
		assert.Nil(t, err)
	})
}
