package registry

import (
	"chatapp/pkg/database"
	"chatapp/pkg/database/repositories"
)

func registerProductPriceFeeRepository(container Container) {
	container.RegisterFunc(repositories.EmptyProductPriceFeeRepo, resolveProductPriceFeeRepository)
}

func resolveProductPriceFeeRepository(container Container) (interface{}, error) {
	return repositories.NewProductPriceFeeRepo(database.GetInstance()), nil
}

func registerFeeRepository(container Container) {
	container.RegisterFunc(repositories.EmptyFeeRepo, resolveFeeRepository)
}

func resolveFeeRepository(container Container) (interface{}, error) {
	ppfRepo := container.Resolve(repositories.EmptyProductPriceFeeRepo).(repositories.IProductPriceFees)
	return repositories.NewFeeRepo(database.GetInstance(), ppfRepo), nil
}

/*func registerService(container Container) {
	container.RegisterFunc(EmptyService, resolveService)
}

func resolveService(container Container) (interface{}, error) {
	ppfRepo := container.Resolve(repositories.EmptyProductPriceFeeRepo).(repositories.IProductPriceFees)
	return NewService(ppfRepo), nil
}*/
