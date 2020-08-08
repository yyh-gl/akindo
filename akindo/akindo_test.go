package akindo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yyh-gl/fx-auto-trader/oanda"
)

func TestAkindo(t *testing.T) {
	at := os.Getenv("API_ACCESS_TOKEN")
	id := os.Getenv("ACCOUNT_ID")

	t.Run("商売人を取得できる", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, "USD_JPY")

		assert.NotNil(t, a)
	})

	t.Run("check()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, "USD_JPY")

		a.check()
	})

	t.Run("buy()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, "USD_JPY")

		a.buy()
	})

	t.Run("sell()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, "USD_JPY")

		a.sell()
	})
}
