package main

import (
	"log"
	"net"

	ds "github.com/ErwinSalas/inventory-service/cmd"
	inventorypb "github.com/ErwinSalas/inventory-service/proto"
	"github.com/ErwinSalas/inventory-service/seed"
	services "github.com/ErwinSalas/inventory-service/services"
	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	datastore := ds.InitDatastore()
	seed.Load(datastore)

	inventoryService := services.NewInventoryService(datastore)
	inventorypb.RegisterInventoryServer(server, inventoryService)
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
