package models

import (
	"chatapp/api/proto/stubs"
	"encoding/json"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

var EmptyProductPriceFee = ProductPriceFees{}

type ProductPriceFees struct {
	ID             uint            `gorm:"primarykey"`
	CreatedAt      time.Time       `gorm:"column:createdAt"`
	UpdatedAt      time.Time       `gorm:"column:updatedAt"`
	ProductId      int             `gorm:"uniqueIndex:ppf_index;not null" json:"productId"`
	Fees           Fee             `gorm:"embedded" json:"fees"`
	TotalPrice     decimal.Decimal `gorm:"type:decimal(20,4);" json:"totalPrice"`
	PVP            decimal.Decimal `gorm:"type:decimal(20,4);" json:"pvp"`
	PVN            decimal.Decimal `gorm:"type:decimal(20,4);" json:"pvn"`
	SellerCurrency Currency        `gorm:"type:varchar(10);not null" json:"sellerCurrency"`
	BuyerCurrency  Currency        `gorm:"uniqueIndex:ppf_index;type:varchar(10);not null" json:"buyerCurrency"`
}

func (*ProductPriceFees) TableName() string {
	return "product_price_fees"
}

func (ppf *ProductPriceFees) BeforeCreate(*gorm.DB) error {
	if !(ppf.BuyerCurrency.ValidateCurrency() && ppf.SellerCurrency.ValidateCurrency()) {
		return ErrInvalidCurrency
	}
	ppf.CreatedAt = time.Now()
	ppf.UpdatedAt = time.Now()
	return nil
}

func (ppf *ProductPriceFees) BeforeUpdate(*gorm.DB) error {
	ppf.UpdatedAt = time.Now()
	return nil
}

func BuildProductPriceFeeEntity(req *stubs.ProductPriceFees) (ProductPriceFees, error) {
	ppf := EmptyProductPriceFee

	if totalPrice := req.GetTotalPrice(); totalPrice != "" {
		if value, err := decimal.NewFromString(totalPrice); err == nil {
			ppf.TotalPrice = value
		} else {
			return EmptyProductPriceFee, err
		}
	}

	if pvn := req.GetPvn(); pvn != "" {
		if value, err := decimal.NewFromString(pvn); err == nil {
			ppf.PVN = value
		} else {
			return EmptyProductPriceFee, err
		}
	}

	if pvp := req.GetPvp(); pvp != "" {
		if value, err := decimal.NewFromString(pvp); err == nil {
			ppf.PVP = value
		} else {
			return EmptyProductPriceFee, err
		}
	}

	if sellerCurrency := req.GetSellerCurrency().String(); sellerCurrency != "" {
		ppf.SellerCurrency = Currency(sellerCurrency)
	}

	ppf.ProductId = int(req.GetProductId())
	ppf.BuyerCurrency = Currency(req.GetBuyerCurrency().String())

	if feeData := req.GetFees().GetFeeData(); feeData != nil {
		fees := Fee{
			FeeData: datatypes.JSONMap{},
		}
		if jsonBytes, err := json.Marshal(feeData); err == nil {
			if err = json.Unmarshal(jsonBytes, &fees.FeeData); err == nil {
				ppf.Fees = fees
			} else {
				return EmptyProductPriceFee, err
			}
		} else {
			return EmptyProductPriceFee, err
		}
	}

	return ppf, nil
}
