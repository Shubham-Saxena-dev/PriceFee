package repositories

import (
	"chatapp/pkg/models"
	"errors"
	"gorm.io/gorm"
)

type productPriceFeeRepo struct {
	db *gorm.DB
}

func NewProductPriceFeeRepo(gormDb *gorm.DB) IProductPriceFees {
	return &productPriceFeeRepo{
		db: gormDb,
	}
}

func (repo *productPriceFeeRepo) CreateProductPriceFee(productPriceFeeEntity models.ProductPriceFees) (int64, error) {

	result := repo.db.Create(&productPriceFeeEntity)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return 0, errors.New("product with id:`productId` already exists")
	} else if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (repo *productPriceFeeRepo) GetProductPriceFee(productId int, buyerCurrency models.Currency) (models.ProductPriceFees, error) {

	var product models.ProductPriceFees
	result := repo.db.Where(&models.ProductPriceFees{
		ProductId:     productId,
		BuyerCurrency: buyerCurrency,
	}).Find(&product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.EmptyProductPriceFee, errors.New("product with id:`productId` not found")
	} else if result.Error != nil {
		return models.EmptyProductPriceFee, result.Error
	}

	return product, nil
}

func (repo *productPriceFeeRepo) UpdateProductPriceFee(productPriceFeeEntity models.ProductPriceFees) (int64, error) {

	result := repo.db.Where(models.ProductPriceFees{
		ProductId:     productPriceFeeEntity.ProductId,
		BuyerCurrency: productPriceFeeEntity.BuyerCurrency,
	}).Updates(&productPriceFeeEntity)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (repo *productPriceFeeRepo) DeleteProductPriceFee(productId int, currency models.Currency) (int64, error) {

	result := repo.db.Unscoped().Where(models.ProductPriceFees{
		ProductId:     productId,
		BuyerCurrency: currency,
	}).Delete(&models.EmptyProductPriceFee)

	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

var EmptyProductPriceFeeRepo = productPriceFeeRepo{}
