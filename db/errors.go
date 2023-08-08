package db

import "errors"

var (
	ErrProductNotExist = errors.New("Sorry...Product you are looking for does not exist")
	ErrZeroAvailable   = errors.New("Sorry...Product you are looking for is out of stock")
	ErrLessAvailable   = errors.New("Sorry we dont have the number of items you want....Please reduce the quantity and try again")
	ErrOrderNotExist   = errors.New("Sorry No order exist at the moment")
)
