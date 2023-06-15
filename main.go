package main

import (
	"chatapp/api/proto/stubs"
	"chatapp/pkg/database"
	"chatapp/pkg/service"
	"chatapp/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	database.InitialiseDatabase()
	initialiseLayers()
	initializeServer()
	/*	s := registry.GetContainer().Resolve(registry.EmptyService).(registry.Service)
		server := gin.Default()
		server.POST("/fee", s.CreateFee)
		server.GET("/fee/:id", s.GetFee)
		server.PUT("/fee/:id", s.UpdateFee)

		server.Run()*/
}

func initializeServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Unable to open server at port {}", 8081)
		panic(err)
	}

	server := grpc.NewServer()
	stubs.RegisterProductPriceFeesServiceServer(server, &service.ContainerKeyGrpcService)
	stubs.RegisterFeeServiceServer(server, &service.ContainerKeyFeeService)
	reflection.Register(server)

	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}

func initialiseLayers() {
	container := registry.NewSimpleIocContainer()
	registry.RegisterContainerMappings(container)
}
