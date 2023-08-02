package product

import (
	"context"

	"github.com/RushikeshMarkad16/e-commerce/db"
	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
	FindByID(ctx context.Context, id int) (response FindByIdResponse, err error)
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

func (cs *productService) FindByID(ctx context.Context, id int) (response FindByIdResponse, err error) {
	product, err := cs.store.FindProductByID(ctx, id)
	if err == db.ErrProductNotExist {
		cs.logger.Error("No product present", "err", err.Error())
		return response, errNoProductId
	}
	if err != nil {
		cs.logger.Error("Error finding product", "err", err.Error(), "id", id)
		return
	}

	response.Product.ID = product.ID
	response.Product.Name = product.Name
	response.Product.Availability = product.Availability
	response.Product.Price = product.Price
	response.Product.Category = product.Category

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &productService{
		store:  s,
		logger: l,
	}
}
