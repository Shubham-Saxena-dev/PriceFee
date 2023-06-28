package registry

import (
	"chatapp/pkg/database/repositories"
	"chatapp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"log"
	"net/http"
	"strconv"
)

type Service interface {
	CreateFee(context *gin.Context)
	GetFee(context *gin.Context)
	UpdateFee(context *gin.Context)
}

type service struct {
	repo repositories.IProductPriceFees
}

func NewService(repo repositories.IProductPriceFees) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateFee(context *gin.Context) {

	item := models.ProductPriceFees{
		BuyerCurrency:  "GBP",
		SellerCurrency: "USD",
		TotalPrice:     decimal.NewFromInt(200),
		ProductId:      3,
		PVN:            decimal.NewFromInt(10),
		PVP:            decimal.NewFromInt(20),
		Fees: models.Fee{
			FeeData: datatypes.JSONMap{"SF": 18, "CBK": 20},
		},
	}

	_, err := s.repo.CreateProductPriceFee(item)
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}
	context.JSON(http.StatusOK, "saved")
}

func (s *service) GetFee(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	response, err := s.repo.GetProductPriceFee(id, "EUR")
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}
	context.JSON(http.StatusOK, response)
}

func (s *service) UpdateFee(context *gin.Context) {

	item := models.ProductPriceFees{
		BuyerCurrency:  "GBP",
		SellerCurrency: "USD",
		TotalPrice:     decimal.NewFromInt(200),
		ProductId:      3,
		PVN:            decimal.NewFromInt(10),
		PVP:            decimal.NewFromInt(20),
		Fees: models.Fee{
			FeeData: datatypes.JSONMap{"TTK": 18, "CBK": 20},
		},
	}

	response, err := s.repo.UpdateProductPriceFee(item)
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}
	context.JSON(http.StatusOK, response)
}

var EmptyService = service{}
