package server

import (
	"net/http"

	"github.com/RushikeshMarkad16/e-commerce/api"
	"github.com/RushikeshMarkad16/e-commerce/product"
	"github.com/gorilla/mux"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", product.List(dep.ProductService)).Methods(http.MethodGet)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
