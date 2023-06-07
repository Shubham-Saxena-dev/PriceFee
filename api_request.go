package models

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type ProductPriceFeesRequest struct {
	ProductId      int             `json:"productId"`
	Fee            json.RawMessage `json:"fees"`
	TotalPrice     decimal.Decimal `json:"totalPrice"`
	PVP            int             `json:"pvp"`
	PVN            int             `json:"pvn"`
	SellerCurrency string          `json:"sellerCurrency"`
	BuyerCurrency  string          `json:"buyerCurrency"`
}
