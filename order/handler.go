package order

import (
	"encoding/json"
	"net/http"

	"github.com/RushikeshMarkad16/e-commerce/api"
)

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var createOrderRequest Order
		err := json.NewDecoder(req.Body).Decode(&createOrderRequest)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.create(req.Context(), createOrderRequest)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created order Successfully"})
	})
}

func isBadRequest(err error) bool {
	return err == errEmptyProductID || err == errEmptyQuantity
}
