package akindo

type action int

const (
	actionNothing = iota + 1
	actionPreparation
	actionBuy
	actionSell
	actionTidyingUp
)

func (a *action) String() string {
	switch *a {
	case actionPreparation:
		return "preparation"
	case actionBuy:
		return "buy"
	case actionSell:
		return "sell"
	case actionTidyingUp:
		return "tidying up"
	}
	return ""
}
