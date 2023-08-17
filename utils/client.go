package utils

import (
	"log"

	"github.com/RushikeshMarkad16/e-commerce/utils/productpb"
	"google.golang.org/grpc"
)

func GrpcClient() *grpc.ClientConn {
	cc, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect : %v", err)
	}
	return cc
}

func GetClient(c *grpc.ClientConn) productpb.ProductServiceClient {
	if c != nil {
		client := productpb.NewProductServiceClient(c)
		return client
	}
	return nil
}
