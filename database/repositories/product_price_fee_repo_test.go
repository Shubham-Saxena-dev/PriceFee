package repositories

import (
	"chatapp/pkg/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("feeRepo", func() {
	var (
		db          *gorm.DB
		feeRepo     IFee
		expectedFee models.ProductPriceFees
	)

	BeforeSuite(func() {
		db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		db.AutoMigrate(&models.ProductPriceFees{})
		ppfRepo := NewProductPriceFeeRepo(db)
		feeRepo = NewFeeRepo(db, ppfRepo)
	})

	AfterSuite(func() {
		db.Migrator().DropTable(&models.ProductPriceFees{})
	})

	Describe("GetFee", func() {
		Context("When the fee exists", func() {
			BeforeEach(func() {
				expectedFee = models.ProductPriceFees{
					ProductId:      1,
					Fees:           models.Fee{FeeData: map[string]interface{}{"SF": "123"}},
					TotalPrice:     decimal.NewFromInt(200),
					PVP:            decimal.NewFromInt(10),
					PVN:            decimal.NewFromInt(20),
					SellerCurrency: "EUR",
					BuyerCurrency:  "USD",
				}
				db.Create(&expectedFee)
			})
			It("should return the fee", func() {
				fee, err := feeRepo.GetFee(1, "USD")
				Ω(err).NotTo(HaveOccurred())
				Ω(fee.FeeData).To(Equal(expectedFee.Fees.FeeData))
			})
		})

		Context("When the fee does not exist", func() {
			It("should return an empty fee", func() {
				fee, err := feeRepo.GetFee(1, "KGF")
				Ω(err).ToNot(HaveOccurred())
				Ω(fee).To(Equal(models.EmptyFee))
			})
		})
	})

	Describe("AddFee", func() {
		Context("When the fee does not exist", func() {
			It("should add the fee", func() {
				feeEntity := models.Fee{
					FeeData: map[string]interface{}{"CBF": "321"},
				}
				rowsAffected, err := feeRepo.AddFee(1, "USD", feeEntity)
				Ω(err).NotTo(HaveOccurred())
				Ω(rowsAffected).To(BeNumerically(">", 0))

				fee, err := feeRepo.GetFee(1, "USD")
				Ω(err).NotTo(HaveOccurred())
				Ω(fee).NotTo(BeNil())
				Expect(fee.FeeData).To(HaveKeyWithValue("CBF", Equal("321")))
			})
		})
	})

	Describe("RemoveFee", func() {
		Context("When the fee exists", func() {
			It("should remove the fee", func() {
				feeEntity := models.Fee{
					FeeData: map[string]interface{}{"SF": "123", "CBF": "321"},
				}
				rowsAffected, err := feeRepo.RemoveFee(1, "USD", feeEntity)
				Ω(err).NotTo(HaveOccurred())
				Ω(rowsAffected).To(BeNumerically(">", 0))

				fee, err := feeRepo.GetFee(1, "USD")
				Ω(fee.FeeData).To(BeEmpty())
			})
		})
	})
})
