package models

import "errors"

var ErrInvalidCurrency = errors.New("currency is invalid")

type Currency string

const (
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	USD Currency = "USD"
	JPY Currency = "JPY"
	SGD Currency = "SGD"
	DKK Currency = "DKK"
	CHF Currency = "CHF"
	PLN Currency = "PLN"
	SEK Currency = "SEK"
	CAD Currency = "CAD"
	HKD Currency = "HKD"
	AUD Currency = "AUD"
	KRW Currency = "KRW"
)

func (c Currency) ValidateCurrency() bool {
	switch c {
	case EUR, GBP, USD, JPY, SGD, DKK, CHF, PLN, SEK, CAD, HKD, AUD, KRW:
		return true
	default:
		return false
	}
}
