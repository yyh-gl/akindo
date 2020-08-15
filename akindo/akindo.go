package akindo

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/yyh-gl/fx-auto-trader/candle"

	"github.com/yyh-gl/fx-auto-trader/oanda"
)

// Akindo : 商売人を表す構造体
type Akindo struct {
	*csv.Writer
	oandaClient *oanda.Client

	// 扱う通貨
	instrument string
	// 1回の取引でやりとりするユニット数
	unitsPerTrade int
	// 最後のアクション
	lastAction action
}

// New : 商売人召喚
// 冒険の書のcloseを忘れずに！
func New(oc *oanda.Client, adventureBook *os.File, instrument string, unitsPerTrade int) *Akindo {
	return &Akindo{
		Writer:        csv.NewWriter(adventureBook),
		oandaClient:   oc,
		instrument:    instrument,
		unitsPerTrade: unitsPerTrade,
		lastAction:    actionNothing,
	}
}

// GoToTrade : トレード開始
func (a *Akindo) GoToTrade(ctx context.Context) error {
	a.save(actionPreparation, nil)

	timer := time.NewTimer(0)

exitLoop:
	for {
		select {
		case <-ctx.Done():
			break exitLoop
		case <-timer.C:
		}

		switch result, candleStick := a.judge(ctx); result {
		case judgeResultBuy:
			a.buy(ctx)
			a.save(actionBuy, candleStick)
		case judgeResultSell:
			a.sell(ctx)
			a.save(actionSell, candleStick)
		default:
			fmt.Println("========================")
			fmt.Println("wait")
			fmt.Println("========================")
			//a.save(actionWait)
		}

		timer = time.NewTimer(2 * time.Minute)
	}

	//a.WriteTradeLog("Finish trade")
	return nil
}

// judge : 価格変動を確認して次のアクションを決定
// TODO: 判定ロジックをNew()時に注入できるようにする
func (a *Akindo) judge(ctx context.Context) (judgeResult, *candle.CandleStick) {
	// TODO: エラーハンドリング
	c, _ := a.oandaClient.GetLatestCandle(ctx, a.instrument)

	mid := (c.Highest + c.Lowest) / 2
	switch {
	case c.IsBullish() && c.Open == c.Lowest && c.Closing <= mid:
		return judgeResultSell, c
	case c.IsBearish() && c.Closing == c.Highest && c.Open >= mid:
		return judgeResultBuy, c
	}
	return judgeResultWait, c
}

// buy : 購入
// FIXME: レシーバをポインタ型にする
func (a *Akindo) buy(ctx context.Context) {
	a.oandaClient.Buy(ctx, a.instrument, a.unitsPerTrade)
	a.lastAction = actionBuy
}

// sell : 売却
func (a *Akindo) sell(ctx context.Context) {
	a.oandaClient.Sell(ctx, a.instrument, a.unitsPerTrade)
	a.lastAction = actionSell
}

// save : 取引記録に書き込む
func (a *Akindo) save(action action, candleStick *candle.CandleStick) {
	now := time.Now().Format(time.RFC3339)

	switch action {
	case actionPreparation:
		// 準備アクションではヘッダー情報を書き込む
		// TODO: エラーハンドリング
		_ = a.Write([]string{"time", "action", "instrument", "units", "candle stick"})
	default:
		// TODO: エラーハンドリング
		_ = a.Write([]string{now, action.String(), a.instrument, strconv.Itoa(a.unitsPerTrade), candleStick.String()})
	}
	a.Flush()
}
