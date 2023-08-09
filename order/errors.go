package order

import "errors"

var (
	errEmptyProductID   = errors.New("Order Id Cannot be empty")
	errEmptyQuantity    = errors.New("Quantity Cannot be empty or zero")
	errProdNotAvailable = errors.New("Sorry....Product you are looking for is currently unavailable")
	errGreaterthanTen   = errors.New("Sorry....cannot order product quantity more than 10")
	errNoOrders         = errors.New("No orders found")
	errNoOrderId        = errors.New("Order ID is not present")
	errEmptyOrderId     = errors.New("Order ID cannot be empty")
	errEmptyStatus      = errors.New("Order Status cannot be empty")
)
