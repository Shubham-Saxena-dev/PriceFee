package service

import (
	"chatapp/api/proto/stubs"
	"chatapp/pkg/database/repositories"
	"chatapp/pkg/models"
	"chatapp/registry"
	"context"
	"encoding/json"
)

var ContainerKeyGrpcService = grpcService{}

type grpcService struct {
	stubs.UnimplementedProductPriceFeesServiceServer
}

func (g *grpcService) CreateProductPriceFee(ctx context.Context, request *stubs.ProductPriceFeeRequest) (*stubs.ProductPriceFeeResponse, error) {
	repo := registry.GetContainer().Resolve(repositories.EmptyProductPriceFeeRepo).(repositories.IProductPriceFees)
	ppfEntity, err := models.BuildProductPriceFeeEntity(request.GetProductPriceFee())
	if err != nil {
		return nil, err
	}
	rowsAffected, err := repo.CreateProductPriceFee(ppfEntity)
	if err != nil {
		return nil, err
	}
	return &stubs.ProductPriceFeeResponse{
		Result: rowsAffected,
		Error:  "",
	}, nil
}

func (g *grpcService) GetProductPriceFee(ctx context.Context, request *stubs.GetProductPriceFeeRequest) (*stubs.GetProductPriceFeeResponse, error) {
	repo := registry.GetContainer().Resolve(repositories.EmptyProductPriceFeeRepo).(repositories.IProductPriceFees)
	ppfResponse, err := repo.GetProductPriceFee(int(request.GetProductId()), models.Currency(request.GetCurrency().String()))
	if err != nil {
		return nil, err
	}
	return g.parseProductPriceFeeResponse(&ppfResponse), nil
}

func (g *grpcService) UpdateProductPriceFee(ctx context.Context, request *stubs.ProductPriceFeeRequest) (*stubs.ProductPriceFeeResponse, error) {
	repo := registry.GetContainer().Resolve(repositories.EmptyProductPriceFeeRepo).(repositories.IProductPriceFees)
	ppfEntity, err := models.BuildProductPriceFeeEntity(request.GetProductPriceFee())
	if err != nil {
		return nil, err
	}
	rowsAffected, err := repo.UpdateProductPriceFee(ppfEntity)
	if err != nil {
		return nil, err
	}
	return &stubs.ProductPriceFeeResponse{
		Result: rowsAffected,
		Error:  "",
	}, nil
}

func (g *grpcService) DeleteProductPriceFee(ctx context.Context, request *stubs.DeleteProductPriceFeeRequest) (*stubs.ProductPriceFeeResponse, error) {
	repo := registry.GetContainer().Resolve(repositories.EmptyProductPriceFeeRepo).(repositories.IProductPriceFees)
	rowsAffected, err := repo.DeleteProductPriceFee(int(request.GetProductId()), models.Currency(request.GetCurrency().String()))
	if err != nil {
		return nil, err
	}
	return &stubs.ProductPriceFeeResponse{
		Result: rowsAffected,
		Error:  "",
	}, nil
}

func (g *grpcService) parseProductPriceFeeResponse(ppfResponse *models.ProductPriceFees) *stubs.GetProductPriceFeeResponse {
	fee := &stubs.Fee{}
	if ppfResponse.Fees.FeeData != nil {
		jsonBytes, err := json.Marshal(ppfResponse.Fees.FeeData)
		if err != nil {
			return nil
		}
		if err := json.Unmarshal(jsonBytes, &fee.FeeData); err != nil {
			return nil
		}
	}

	return &stubs.GetProductPriceFeeResponse{
		ProductPriceFee: &stubs.ProductPriceFees{
			ProductId:      int32(ppfResponse.ProductId),
			TotalPrice:     ppfResponse.TotalPrice.String(),
			Pvp:            ppfResponse.PVP.String(),
			Pvn:            ppfResponse.PVN.String(),
			SellerCurrency: stubs.Currency(stubs.Currency_value[string(ppfResponse.SellerCurrency)]),
			BuyerCurrency:  stubs.Currency(stubs.Currency_value[string(ppfResponse.BuyerCurrency)]),
			Fees:           fee,
		},
	}
}
