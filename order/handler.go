package order

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RushikeshMarkad16/e-commerce/api"
	"github.com/gorilla/mux"
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

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.List(req.Context())
		if err == errNoOrders {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		var temp []Order1
		for _, j := range resp.Orders {
			var temp1 Order1
			temp1.ID = j.ID
			temp1.Amount = j.Amount
			temp1.Disc_perc = j.Disc_perc
			temp1.Final_amnt = j.Final_amnt
			temp1.Disp_date = j.Disp_date
			temp1.Order_status = j.Order_status
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

		if err == errNoOrderId {
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
