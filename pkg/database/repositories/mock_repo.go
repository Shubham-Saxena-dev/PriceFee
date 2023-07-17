package repositories

import "chatapp/pkg/models"

type mockFeeRepo struct {
	GetFeeFunc    func(productId int, buyerCurrency string) (models.Fee, error)
	AddFeeFunc    func(productId int, buyerCurrency string, feeEntity models.Fee) (int64, error)
	RemoveFeeFunc func(productId int, buyerCurrency string, feeEntity models.Fee) (int64, error)
}

func (m *mockFeeRepo) GetFee(productId int, buyerCurrency string) (models.Fee, error) {
	if m.GetFeeFunc != nil {
		return m.GetFeeFunc(productId, buyerCurrency)
	}
	return models.EmptyFee, nil
}

func (m *mockFeeRepo) AddFee(productId int, buyerCurrency string, feeEntity models.Fee) (int64, error) {
	if m.AddFeeFunc != nil {
		return m.AddFeeFunc(productId, buyerCurrency, feeEntity)
	}
	return 0, nil
}

func (m *mockFeeRepo) RemoveFee(productId int, buyerCurrency string, feeEntity models.Fee) (int64, error) {
	if m.RemoveFeeFunc != nil {
		return m.RemoveFeeFunc(productId, buyerCurrency, feeEntity)
	}
	return 0, nil
}
