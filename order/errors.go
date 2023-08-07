package order

import "errors"

var (
	errEmptyProductID   = errors.New("Order Id Cannot be empty")
	errEmptyQuantity    = errors.New("Quantity Cannot be empty or zero")
	errProdNotAvailable = errors.New("Sorry....Product you are looking for is currently unavailable")
)
