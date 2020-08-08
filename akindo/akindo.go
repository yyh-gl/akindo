package akindo

import (
	"context"
	"log"
	"os"

	"github.com/yyh-gl/fx-auto-trader/oanda"
)

// Akindo : 商売人を表す構造体
type Akindo struct {
	oandaClient *oanda.Client
	logger      *log.Logger
	instrument  string
}

// New : 商売人召喚
func New(oc *oanda.Client, instrument string) *Akindo {
	return &Akindo{
		oandaClient: oc,
		logger:      log.New(os.Stdout, "Akindo", log.LstdFlags),
		instrument:  instrument,
	}
}

// GoToTrade : トレード開始
func (a Akindo) GoToTrade(ctx context.Context) error {
	a.logger.Println("Start trade")

exitLoop:
	for {
		select {
		case <-ctx.Done():
			break exitLoop
		default:
		}

		switch result := a.judge(); result {
		case judgeResultBuy:
			a.buy()
			a.logger.Println("Buy")
		case judgeResultSell:
			a.sell()
			a.logger.Println("Sell")
		default:
			a.logger.Println("...")
		}
	}

	a.logger.Println("Finish trade")
	return nil
}

// judge : 価格変動を確認して次のアクションを決定
func (a Akindo) judge() judgeResult {
	return judgeResultWait
}

// buy : 購入
func (a Akindo) buy() {}

// sell : 売却
func (a Akindo) sell() {}
