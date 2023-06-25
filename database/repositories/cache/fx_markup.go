package cache

type CurrencyInfo struct {
	Code          string
	DecimalDigits int
}

type CurrencyConversion struct {
	From CurrencyInfo
	To   CurrencyInfo
	Rate float64
}

type FxMarkup CurrencyConversion

type Rates CurrencyConversion
