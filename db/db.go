package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ctxKey int

const (
	defaultTimeout        = 1 * time.Second
	dbKey          ctxKey = 0
)

type Storer interface {
	//Product
	ListProducts(ctx context.Context) (products []Product, err error)
	FindProductByID(ctx context.Context, id int) (product Product, err error)

	//Order
	CreateOrder(ctx context.Context, order []*Order_item) (err error)
	ListOrders(ctx context.Context) (orders []Order1, err error)
	FindOrderByID(ctx context.Context, id int) (order Order1, err error)
	UpdateOrderStatus(ctx context.Context, order *Order1) (err error)
}

type store struct {
	db *sqlx.DB
}

func newContext(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, dbKey, tx)
}

func Transact(ctx context.Context, dbx *sqlx.DB, opts *sql.TxOptions, txFunc func(context.Context) error) (err error) {
	tx, err := dbx.BeginTxx(ctx, opts)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = errors.WithStack(p)
			default:
				err = errors.Errorf("%s", p)
			}
		}
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				err = errors.WithStack(e)
			}
			return
		}
		err = errors.WithStack(tx.Commit())
	}()
	ctxWithTx := newContext(ctx, tx)
	err = WithDefaultTimeout(ctxWithTx, txFunc)
	return err
}

func WithTimeout(ctx context.Context, timeout time.Duration, op func(ctx context.Context) error) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return op(ctxWithTimeout)
}

func WithDefaultTimeout(ctx context.Context, op func(ctx context.Context) error) (err error) {
	return WithTimeout(ctx, defaultTimeout, op)
}

func NewStorer(d *sqlx.DB) Storer {
	return &store{
		db: d,
	}
}
