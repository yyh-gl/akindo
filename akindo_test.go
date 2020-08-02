package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAkindo(t *testing.T) {
	t.Run("アカウント情報を取得できること", func(t *testing.T) {
		ac := NewAkindoClient()
		a, err := ac.GetAccount()

		assert.Equal(t, "", a)
		assert.Nil(t, err)
	})
}
