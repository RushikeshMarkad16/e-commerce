package product

import "errors"

var (
	errNoProducts  = errors.New("No products present")
	errNoProductId = errors.New("Product id is not present")
)
