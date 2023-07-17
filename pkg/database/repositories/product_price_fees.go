package repositories

import "chatapp/pkg/models"

type IProductPriceFees interface {
	CreateProductPriceFee(models.ProductPriceFees) (int64, error)
	GetProductPriceFee(int, models.Currency) (models.ProductPriceFees, error)
	UpdateProductPriceFee(models.ProductPriceFees) (int64, error)
	DeleteProductPriceFee(int, models.Currency) (int64, error)
}
