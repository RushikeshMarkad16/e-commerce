package product

import (
	"net/http"
	"strconv"

	"github.com/RushikeshMarkad16/e-commerce/api"
	"github.com/gorilla/mux"
)

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.List(req.Context())
		if err == errNoProducts {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		var temp []Product
		for _, j := range resp.Products {
			var temp1 Product
			temp1.ID = j.ID
			temp1.Name = j.Name
			temp1.Availability = j.Availability
			temp1.Price = j.Price
			temp1.Category = j.Category
			temp = append(temp, temp1)
		}

		api.Success(rw, http.StatusOK, temp)
	})
}

func FindByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		vars1, _ := strconv.Atoi(vars["id"])
		resp, err := service.FindByID(req.Context(), vars1)

		if err == errNoProductId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}
