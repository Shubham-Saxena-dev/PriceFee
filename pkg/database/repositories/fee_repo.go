package repositories

import (
	"chatapp/pkg/models"
	"golang.org/x/exp/maps"
	"gorm.io/gorm"
)

type feeRepo struct {
	db      *gorm.DB
	ppfRepo IProductPriceFees
}

func NewFeeRepo(db *gorm.DB, ppfRepo IProductPriceFees) IFee {

	return &feeRepo{
		db:      db,
		ppfRepo: ppfRepo,
	}
}

func (repo *feeRepo) GetFee(productId int, buyerCurrency models.Currency) (models.Fee, error) {

	var fee models.Fee
	result := repo.db.Where(&models.ProductPriceFees{
		ProductId:     productId,
		BuyerCurrency: buyerCurrency,
	}).Find(&fee)

	if result.Error != nil {
		return models.EmptyFee, result.Error
	}
	return fee, nil
}

func (repo *feeRepo) AddFee(productId int, buyerCurrency models.Currency, feeEntity models.Fee) (int64, error) {

	fee, err := repo.GetFee(productId, buyerCurrency)
	if err != nil {
		return 0, err
	}
	maps.Copy(fee.FeeData, feeEntity.FeeData)
	rowsAffected, err := repo.ppfRepo.UpdateProductPriceFee(models.ProductPriceFees{
		Fees:          fee,
		ProductId:     productId,
		BuyerCurrency: buyerCurrency,
	})
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (repo *feeRepo) RemoveFee(productId int, buyerCurrency models.Currency, feeEntity models.Fee) (int64, error) {

	fee, err := repo.GetFee(productId, buyerCurrency)
	if err != nil {
		return 0, err
	}
	for key := range feeEntity.FeeData {
		delete(fee.FeeData, key)
	}

	rowsAffected, err := repo.ppfRepo.UpdateProductPriceFee(models.ProductPriceFees{
		Fees:          fee,
		ProductId:     productId,
		BuyerCurrency: buyerCurrency,
	})
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

var EmptyFeeRepo = feeRepo{}
