package kv

type (
	MarketType string
	Target     int
	Customer   string
	Period     string
	OrderType  string
)

const (
	Imitation Target = iota + 1
	Real

	Person   Customer = "P"
	Business Customer = "B"

	Stock MarketType = "J"
	ETF   MarketType = "ETF"
	ETN   MarketType = "ETN"

	Day   Period = "D"
	Week  Period = "W"
	Month Period = "M"

	Limit         OrderType = "00"
	MarketPrice   OrderType = "01"
	Conditional   OrderType = "02"
	BestAdvantage OrderType = "03"
	BestPriority  OrderType = "04"
	PreMarket     OrderType = "05"
	PostMarket    OrderType = "06"
)
