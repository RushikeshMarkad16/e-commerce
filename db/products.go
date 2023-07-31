package db

import (
	"context"
	"database/sql"
)

const (
	listProductsQuery = `SELECT id, name, availability, price, category FROM product ORDER BY name`
)

type Product struct {
	ID           int32
	Name         string
	Availability int32
	Price        int32
	Category     string
}

func (s *store) ListProducts(ctx context.Context) (products []Product, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &products, listProductsQuery)
	})
	if err == sql.ErrNoRows {
		return products, ErrProductNotExist
	}
	return
}
