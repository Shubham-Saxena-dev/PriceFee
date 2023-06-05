package main

import (
	"chatapp/internal/registry"
	"chatapp/pkg/database"
	"chatapp/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitialiseDatabase()
	initialiseLayers()
	s := registry.GetContainer().Resolve(service.EmptyService).(service.Service)
	server := gin.Default()
	server.POST("/fee", s.CreateFee)
	server.GET("/fee/:id", s.GetFee)
	server.PUT("/fee/:id", s.UpdateFee)

	server.Run()
}

func initialiseLayers() {
	container := registry.NewSimpleIocContainer()
	registry.RegisterContainerMappings(container)
}
