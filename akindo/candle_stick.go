package akindo

type (
	// CandleStick : ローソク足情報を表す構造体
	CandleStick struct {
		Complete bool
		Open     float64
		Highest  float64
		Lowest   float64
		Closing  float64
	}

	// CandleSticks : ローソク足情報の集合を表す構造体
	CandleSticks []*CandleStick
)

// IsBullish : 陽線であるかどうか判定
func (c *CandleStick) IsBullish() bool {
	if c.Open < c.Closing {
		return true
	}
	return false
}

// IsBearish : 陰線であるかどうか判定
func (c *CandleStick) IsBearish() bool {
	if c.Closing < c.Open {
		return true
	}
	return false
}

// IsNeutral : 拮抗状態（1本線に見えるやつ）かどうか判定
func (c *CandleStick) IsNeutral() bool {
	if c.Open == c.Closing {
		return true
	}
	return false
}
