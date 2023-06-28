package registry

func RegisterContainerMappings(container Container) {

	registerFeeRepository(container)
	registerProductPriceFeeRepository(container)
	//registerService(container)
}
