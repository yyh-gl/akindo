package akindo

import (
	"context"

	"github.com/yyh-gl/fx-auto-trader/oanda"
)

// Akindo : 商売人を表す構造体
type Akindo struct {
	adventureBookWriter

	oandaClient *oanda.Client
	instrument  string
}

// New : 商売人召喚
func New(oc *oanda.Client, instrument string) *Akindo {
	return &Akindo{
		oandaClient:         oc,
		adventureBookWriter: newAdventureBookWriter(),
		instrument:          instrument,
	}
}

// GoToTrade : トレード開始
func (a Akindo) GoToTrade(ctx context.Context) error {
	a.WriteTradeLog("Start trade")

exitLoop:
	for {
		select {
		case <-ctx.Done():
			break exitLoop
		default:
		}

		switch result := a.judge(ctx); result {
		case judgeResultBuy:
			a.buy(ctx)
			a.WriteTradeLog("Buy")
		case judgeResultSell:
			a.sell(ctx)
			a.WriteTradeLog("Sell")
		default:
			a.WriteTradeLog("Wait")
		}
	}

	a.WriteTradeLog("Finish trade")
	return nil
}

// judge : 価格変動を確認して次のアクションを決定
func (a Akindo) judge(ctx context.Context) judgeResult {
	return judgeResultWait
}

// buy : 購入
func (a Akindo) buy(ctx context.Context) {}

// sell : 売却
func (a Akindo) sell(ctx context.Context) {}
