package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/RushikeshMarkad16/e-commerce/config"
	"github.com/RushikeshMarkad16/e-commerce/utils/productpb"
	"github.com/urfave/negroni"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartAPIServer(wg *sync.WaitGroup) {
	fmt.Println("Starting Rest Server....")
	defer wg.Done()
	port := config.AppPort()
	server := negroni.Classic()

	dependencies, err := initDependencies()
	if err != nil {
		panic(err)
	}

	router := initRouter(dependencies)
	server.UseHandler(router)
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(addr)
}

func StartgRPCServer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting grpc server....")
	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	dependencies, err := initDependencies()
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	productpb.RegisterProductServiceServer(s, NewGrpcServer(dependencies.ProductService))
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
