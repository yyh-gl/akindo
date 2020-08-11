package akindo

import (
	"context"
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

	t.Run("judge()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, "USD_JPY")

		got := a.judge(context.Background())
		assert.Equal(t, got, judgeResultWait)
	})

	t.Run("buy()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, "USD_JPY")

		a.buy(context.Background())
	})

	t.Run("sell()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, "USD_JPY")

		a.sell(context.Background())
	})
}
