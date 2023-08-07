package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

const (
	ProductIdPresentQuery = `SELECT COUNT(*) FROM product where id = ?`
	ProductStatusQuery    = `SELECT COUNT(*) from product WHERE id = ? AND availability >= ?`
	CreateOrderQuery      = `INSERT INTO order1 VALUES (DEFAULT,0,0,0,NULL,"Placed")`
	GetOrderIdQuery       = `SELECT id FROM order1 ORDER BY id DESC LIMIT 1 ` //to get last id
	CreateOrderItemQuery  = `INSERT INTO order_item (id, order_id, product_id, quantity) VALUES(?, ?, ?, ?)`
	UpdQuantityQuery      = `UPDATE product SET availability = availability-? WHERE id = ? AND availability>0`
	// GetAmountQuery=`INSERT INTO`
	// GetCurrQuantityQuery  = `SELECT availability from product where id = ?`
)

type Order1 struct {
	ID           int
	Amount       int
	Disc_perc    int
	Final_amnt   int
	Disp_date    time.Time
	Order_status string
}

type Order_item struct {
	ID         string
	Order_id   int
	Product_id int
	Quantity   int
}

func (s *store) CreateOrder(ctx context.Context, orders []*Order_item) (err error) {

	for _, order := range orders {
		count := 0
		s.db.GetContext(ctx, &count, ProductIdPresentQuery, order.Product_id)
		if count < 1 {
			return ErrProductNotExist
		}

		curr_quantity := 0
		s.db.GetContext(ctx, &curr_quantity, ProductStatusQuery, order.Product_id, order.Quantity)
		if curr_quantity == 0 {
			return ErrZeroAvailable
		}
		if curr_quantity < 1 {
			return ErrLessAvailable
		}
	}
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {

		_, err = s.db.ExecContext(ctx,
			CreateOrderQuery,
		)

		if err != nil {
			fmt.Printf("error while crating order %s", err.Error())
			return err
		}

		row := s.db.QueryRow(GetOrderIdQuery)

		var o_id int
		err = row.Scan(&o_id)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No order id present")
				return err
			} else {
				fmt.Printf("Error while fetchng order id : %v", err)
			}
		}

		for _, order := range orders {
			uid := uuid.New().String()

			_, err = s.db.Exec(
				CreateOrderItemQuery,
				uid,
				o_id,
				order.Product_id,
				order.Quantity,
			)

			if err != nil {
				log.Printf("error :%s", err.Error())
				return err
			}

			_, err = s.db.Exec(
				UpdQuantityQuery,
				order.Quantity,
				order.Product_id,
			)
			if err != nil {
				return err
			}
		}

		return err
	})
}
