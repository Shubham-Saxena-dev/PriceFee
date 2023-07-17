package repositories

import "chatapp/pkg/models"

type IFee interface {
	GetFee(productId int, buyerCurrency models.Currency) (models.Fee, error)
	AddFee(productId int, buyerCurrency models.Currency, feeEntity models.Fee) (int64, error)
	RemoveFee(productId int, buyerCurrency models.Currency, feeEntity models.Fee) (int64, error)
}
