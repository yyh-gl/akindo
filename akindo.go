package main

// akindo : 商売人を表す構造体
type akindo struct {
	*oandaClient
	instrument string
}

// newAkindo : 商売人召喚
func newAkindo(oc *oandaClient, instrument string) *akindo {
	return &akindo{
		oandaClient: oc,
		instrument:  instrument,
	}
}

// check : 価格変動を確認
func (a akindo) check() {}

// buy : 購入
func (a akindo) buy() {}

// sell : 売却
func (a akindo) sell() {}
