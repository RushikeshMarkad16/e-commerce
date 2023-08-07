package server

import (
	"github.com/RushikeshMarkad16/e-commerce/app"
	"github.com/RushikeshMarkad16/e-commerce/db"
	"github.com/RushikeshMarkad16/e-commerce/order"
	"github.com/RushikeshMarkad16/e-commerce/product"
)

type dependencies struct {
	ProductService product.Service
	OrderService   order.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	productService := product.NewService(dbStore, logger)
	orderService := order.NewService(dbStore, logger)

	return dependencies{
		ProductService: productService,
		OrderService:   orderService,
	}, nil
}
