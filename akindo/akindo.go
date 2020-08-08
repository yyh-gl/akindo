package akindo

import (
	"github.com/yyh-gl/fx-auto-trader/oanda"
)

// Akindo : 商売人を表す構造体
type Akindo struct {
	oandaClient *oanda.Client
	instrument  string
}

// New : 商売人召喚
func New(oc *oanda.Client, instrument string) *Akindo {
	return &Akindo{
		oandaClient: oc,
		instrument:  instrument,
	}
}

//// GoToTrade : トレード開始
//func (a Akindo) GoToTrade(ctx context.Context) error {
//	for {
//		a.check()
//	}
//
//	return nil
//}

// judge : 価格変動を確認して次のアクションを決定
func (a Akindo) judge() judgeResult {
	return judgeResultWait
}

// buy : 購入
func (a Akindo) buy() {}

// sell : 売却
func (a Akindo) sell() {}
