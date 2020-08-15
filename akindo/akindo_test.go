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
		a := New(oc, nil, "USD_JPY", 10)

		assert.NotNil(t, a)
	})

	t.Run("judge()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, nil, "USD_JPY", 10)
		gotResult, gotCandle := a.judge(context.Background())

		// 判定結果は毎回変わるのでとりあえず空じゃないことだけ確認
		assert.NotEmpty(t, gotResult)
		assert.NotEmpty(t, gotCandle)
	})

	t.Run("buy()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, nil, "USD_JPY", 10)

		a.buy(context.Background())
	})

	t.Run("sell()仮テスト", func(t *testing.T) {
		oc := oanda.NewClient(at, id)
		a := New(oc, nil, "USD_JPY", 10)

		a.sell(context.Background())
	})
}
