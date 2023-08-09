package server

import (
	"net/http"

	"github.com/RushikeshMarkad16/e-commerce/api"
	"github.com/RushikeshMarkad16/e-commerce/order"
	"github.com/RushikeshMarkad16/e-commerce/product"
	"github.com/gorilla/mux"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()

	//Product
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", product.List(dep.ProductService)).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", product.FindByID(dep.ProductService)).Methods(http.MethodGet)

	//Order
	router.HandleFunc("/order", order.Create(dep.OrderService)).Methods(http.MethodPost)
	router.HandleFunc("/orders", order.List(dep.OrderService)).Methods(http.MethodGet)
	router.HandleFunc("/order/{id}", order.FindByID(dep.OrderService)).Methods(http.MethodGet)
	router.HandleFunc("/order/status", order.UpdateStatus(dep.OrderService)).Methods(http.MethodPatch)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
