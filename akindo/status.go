package akindo

type judgeResult int

const (
	judgeResultWait judgeResult = iota + 1
	judgeResultBuy
	judgeResultSell
)
