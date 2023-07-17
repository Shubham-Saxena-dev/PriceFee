package service

import (
	"chatapp/api/proto/stubs"
	"chatapp/pkg/database/repositories"
	"chatapp/pkg/models"
	"chatapp/registry"
	"context"
	"encoding/json"
)

var ContainerKeyFeeService = feeService{}

type feeService struct {
	stubs.UnimplementedFeeServiceServer
}

func (service *feeService) GetFee(_ context.Context, request *stubs.GetFeeRequest) (*stubs.GetFeeResponse, error) {
	result, err := service.getFeeRepository().GetFee(int(request.GetProductId()), models.Currency(request.GetBuyerCurrency().String()))
	if err != nil {
		return &stubs.GetFeeResponse{}, nil
	}
	return &stubs.GetFeeResponse{
		Fee:   ParseIntoFeeStub(result),
		Error: "",
	}, nil
}

func (service *feeService) AddFee(_ context.Context, request *stubs.FeeRequest) (*stubs.FeeResponse, error) {
	result, err := service.getFeeRepository().AddFee(int(request.GetProductId()),
		models.Currency(request.GetBuyerCurrency().String()),
		ParseFromFeeStub(request.GetFee()))
	if err != nil {
		return &stubs.FeeResponse{}, err
	}

	return &stubs.FeeResponse{
		Result: result,
		Error:  "",
	}, nil
}

func (service *feeService) RemoveFee(_ context.Context, request *stubs.FeeRequest) (*stubs.FeeResponse, error) {
	result, err := service.getFeeRepository().RemoveFee(int(request.GetProductId()),
		models.Currency(request.GetBuyerCurrency().String()),
		ParseFromFeeStub(request.GetFee()))
	if err != nil {
		return &stubs.FeeResponse{}, err
	}

	return &stubs.FeeResponse{
		Result: result,
		Error:  "",
	}, nil
}

func (service *feeService) getFeeRepository() repositories.IFee {
	return registry.GetContainer().Resolve(repositories.EmptyFeeRepo).(repositories.IFee)
}

func ParseIntoFeeStub(fee models.Fee) *stubs.Fee {
	parsedFee := &stubs.Fee{}
	if fee.FeeData != nil {
		if err := marshalUnmarshalJSON(fee.FeeData, &parsedFee.FeeData); err != nil {
			return nil
		}
	}
	return parsedFee
}

func ParseFromFeeStub(fee *stubs.Fee) models.Fee {
	parsedFee := &models.Fee{}
	if fee.FeeData != nil {
		if err := marshalUnmarshalJSON(fee.FeeData, &parsedFee.FeeData); err != nil {
			return models.EmptyFee
		}
	}
	return *parsedFee
}

func marshalUnmarshalJSON(data interface{}, target interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return err
	}
	return nil
}
