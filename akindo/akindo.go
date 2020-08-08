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

// check : 価格変動を確認
func (a Akindo) check() {}

// buy : 購入
func (a Akindo) buy() {}

// sell : 売却
func (a Akindo) sell() {}
