package service

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FeeService Integration Tests")
}

/*
var (
	feeClient stubs.FeeServiceClient
	conn      *grpc.ClientConn
)

var _ = Describe("FeeService Integration Tests", func() {

	BeforeSuite(func() {
		var err error
		conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
		Expect(err).NotTo(HaveOccurred())
		feeClient = stubs.NewFeeServiceClient(conn)
		ioc.GetContainer()
	})

	AfterSuite(func() {
		err := conn.Close()
		if err != nil {
			return
		}
	})
})*/
