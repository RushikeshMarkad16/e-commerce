package server

import (
	"net/http"

	"github.com/RushikeshMarkad16/e-commerce/order"
	"github.com/RushikeshMarkad16/e-commerce/product"
	"github.com/gorilla/mux"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()

	//Product
	router.HandleFunc("/products", product.List(dep.ProductService)).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", product.FindByID(dep.ProductService)).Methods(http.MethodGet)

	//Order
	router.HandleFunc("/order", order.Create(dep.OrderService)).Methods(http.MethodPost)
	router.HandleFunc("/orders", order.List(dep.OrderService)).Methods(http.MethodGet)
	router.HandleFunc("/order/{id}", order.FindByID(dep.OrderService)).Methods(http.MethodGet)
	router.HandleFunc("/order/status", order.UpdateStatus(dep.OrderService)).Methods(http.MethodPatch)

	return
}
