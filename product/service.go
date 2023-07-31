package product

import (
	"context"

	"github.com/RushikeshMarkad16/e-commerce/db"
	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
}

type productService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *productService) List(ctx context.Context) (response ListResponse, err error) {
	dbProducts, err := cs.store.ListProducts(ctx)
	if err == db.ErrProductNotExist {
		cs.logger.Error("No product present", "err", err.Error())
		return response, errNoProducts
	}
	if err != nil {
		cs.logger.Error("Error listing products", "err", err.Error())
		return
	}

	for _, dbProduct := range dbProducts {
		var productData Product
		productData.ID = dbProduct.ID
		productData.Name = dbProduct.Name
		productData.Availability = dbProduct.Availability
		productData.Price = dbProduct.Price
		productData.Category = dbProduct.Category

		response.Products = append(response.Products, productData)
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &productService{
		store:  s,
		logger: l,
	}
}
