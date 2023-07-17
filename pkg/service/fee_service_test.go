package service

import (
	"chatapp/api/proto/stubs"
	"chatapp/pkg/database/repositories"
	"chatapp/pkg/models"
	"chatapp/registry"
	"context"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
)

var _ = Describe("FeeService", func() {
	var (
		server     *grpc.Server
		client     stubs.FeeServiceClient
		conn       *grpc.ClientConn
		testData   *stubs.GetFeeRequest
		listener   net.Listener
		serverPort = 50052
	)

	BeforeSuite(func() {
		var err error

		Db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		err = Db.AutoMigrate(&models.ProductPriceFees{})
		if err != nil {
			return
		}
		ppfRepo := repositories.NewProductPriceFeeRepo(Db)
		productPriceFee := models.ProductPriceFees{
			ProductId:      1234,
			Fees:           models.Fee{FeeData: map[string]interface{}{"CBF": "321"}},
			TotalPrice:     decimal.NewFromInt(500),
			PVP:            decimal.NewFromInt(40),
			PVN:            decimal.NewFromInt(50),
			SellerCurrency: "GBP",
			BuyerCurrency:  "CAD",
		}
		ppfRepo.CreateProductPriceFee(productPriceFee)

		feeRepos := repositories.NewFeeRepo(Db, ppfRepo)
		registry.NewSimpleIocContainer()
		registry.GetContainer().RegisterInstance(repositories.EmptyFeeRepo, feeRepos)

		listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", serverPort))
		Expect(err).NotTo(HaveOccurred())

		server = grpc.NewServer()
		stubs.RegisterFeeServiceServer(server, &ContainerKeyFeeService)

		go func() {
			server.Serve(listener)
		}()

		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", serverPort), grpc.WithInsecure())
		Expect(err).NotTo(HaveOccurred())

		client = stubs.NewFeeServiceClient(conn)
	})

	BeforeEach(func() {

		testData = &stubs.GetFeeRequest{
			ProductId:     1234,
			BuyerCurrency: 9,
		}
	})

	AfterSuite(func() {
		err := conn.Close()
		if err != nil {
			return
		}
		server.Stop()
	})

	Describe("GetFee", func() {
		Context("when a valid GetFee request is made", func() {
			It("should return a valid GetFeeResponse", func() {
				response, err := client.GetFee(context.Background(), testData)
				Expect(err).NotTo(HaveOccurred())

				Expect(response).NotTo(BeNil())
				Expect(response.Fee).NotTo(BeNil())
				Expect(response.Error).To(BeEmpty())
			})
		})
	})

	Describe("AddFee", func() {
		Context("when a valid AddFee request is made", func() {
			It("should return a valid FeeResponse", func() {
				addFeeRequest := &stubs.FeeRequest{
					ProductId:     testData.ProductId,
					BuyerCurrency: testData.BuyerCurrency,
					Fee:           &stubs.Fee{FeeData: map[string]int32{"SF": 123}},
				}

				response, err := client.AddFee(context.Background(), addFeeRequest)
				Expect(err).NotTo(HaveOccurred())

				Expect(response).NotTo(BeNil())
				Expect(response.Result).NotTo(BeNil())
				Expect(response.Error).To(BeEmpty())
			})
		})
	})

	Describe("RemoveFee", func() {
		Context("when a valid RemoveFee request is made", func() {
			It("should return a valid FeeResponse", func() {
				removeFeeRequest := &stubs.FeeRequest{
					ProductId:     testData.ProductId,
					BuyerCurrency: testData.BuyerCurrency,
					Fee:           &stubs.Fee{FeeData: map[string]int32{"SF": 123}},
				}

				response, err := client.RemoveFee(context.Background(), removeFeeRequest)
				Expect(err).NotTo(HaveOccurred())

				Expect(response).NotTo(BeNil())
				Expect(response.Result).NotTo(BeNil())
				Expect(response.Error).To(BeEmpty())
			})
		})
	})
})
