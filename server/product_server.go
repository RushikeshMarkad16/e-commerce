package server

import (
	"context"
	"fmt"

	"github.com/RushikeshMarkad16/e-commerce/product"
	"github.com/RushikeshMarkad16/e-commerce/utils/productpb"
)

type server struct {
	productpb.ProductServiceServer
	service product.Service
}

func NewGrpcServer(service product.Service) productpb.ProductServiceServer {
	return &server{
		service: service,
	}
}

func (s *server) GetProduct(ctx context.Context, req *productpb.GetProductByIDRequest) (resp *productpb.ProductResponse, err error) {
	i_id := int(req.GetId())
	fmt.Println("Above FindByID")
	product, err := s.service.FindByID(ctx, i_id)
	fmt.Println("Below FindByID")
	if err != nil {
		return nil, err
	}
	return &productpb.ProductResponse{
		Id:           product.Product.ID,
		Name:         product.Product.Name,
		Availability: product.Product.Availability,
		Price:        product.Product.Price,
		Category:     product.Product.Category,
	}, nil
}

// func (s *server) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (resp *productpb.ProductResponse, err error) {
// err:=s.service.
// write update service in product service
// }
