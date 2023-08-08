package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

const (
	ProductIdPresentQuery = `SELECT COUNT(*) FROM product where id = ?`
	ProductStatusQuery    = `SELECT COUNT(*) from product WHERE id = ? AND availability >= ?`
	CreateOrderQuery      = `INSERT INTO order1 VALUES (DEFAULT,0,0,0,NULL,"Placed")`
	GetOrderIdQuery       = `SELECT id FROM order1 ORDER BY id DESC LIMIT 1 ` //to get last id
	CreateOrderItemQuery  = `INSERT INTO order_item (id, order_id, product_id, quantity) VALUES(?, ?, ?, ?)`
	UpdQuantityQuery      = `UPDATE product SET availability = availability-? WHERE id = ? AND availability>0`
	GetPriceQuery         = `SELECT price from product WHERE id = ?`
	GetCurrAmount         = `SELECT amount from order1 where id = ?`
	UpdAmntQuery          = `UPDATE order1 SET amount = amount+? WHERE id = ?`
	GetCategoryQuery      = `SELECT category FROM product WHERE id = ?`
	GetAmountQuery        = `SELECT amount FROM order1 WHERE id = ?`
	UpdFinalAmount        = `UPDATE order1 SET discount_perc = ? , final_amount = ? WHERE id = ?`
	listOrdersQuery       = `SELECT id, amount, discount_perc, final_amount, dispatch_date, order_status FROM order1 ORDER BY id`
)

type Order1 struct {
	ID           int            `db:"id"`
	Amount       int            `db:"amount"`
	Disc_perc    int            `db:"discount_perc"`
	Final_amnt   int            `db:"final_amount"`
	Disp_date    sql.NullString `db:"dispatch_date"`
	Order_status string         `db:"order_status"`
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

		_, err = s.db.ExecContext(
			ctx,
			CreateOrderQuery,
		)

		if err != nil {
			fmt.Printf("error while creating order %s", err.Error())
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
				fmt.Printf("Error while fetching order id : %v", err)
			}
		}

		premium_count := 0

		for _, order := range orders {
			uid := uuid.New().String()

			//To create order1 entry
			_, err = s.db.Exec(
				CreateOrderItemQuery,
				uid,
				o_id,
				order.Product_id,
				order.Quantity,
			)
			if err != nil {
				return err
			}

			//Retrive the price of product
			price_row := s.db.QueryRow(
				GetPriceQuery,
				order.Product_id,
			)
			var o_price int
			err = price_row.Scan(&o_price)

			//Retrive the category of product
			catg_row := s.db.QueryRow(
				GetCategoryQuery,
				order.Product_id,
			)
			var o_catg string
			err = catg_row.Scan(&o_catg)
			if o_catg == "Premium" {
				premium_count = premium_count + 1
			}

			o_amnt := order.Quantity * o_price

			_, err := s.db.Exec(
				UpdAmntQuery,
				o_amnt,
				o_id,
			)
			if err != nil {
				return err
			}

			//Update the product quantity
			_, err = s.db.Exec(
				UpdQuantityQuery,
				order.Quantity,
				order.Product_id,
			)

			if err != nil {
				return err
			}
		}

		amount_row := s.db.QueryRow(
			GetAmountQuery,
			o_id,
		)
		var o_amount int
		err = amount_row.Scan(&o_amount)
		final_amount := o_amount
		disc := 0
		if premium_count >= 3 {
			final_amount = o_amount - (o_amount * 10 / 100)
			disc = 10
		}
		_, err = s.db.Exec(
			UpdFinalAmount,
			disc,
			final_amount,
			o_id,
		)
		if err != nil {
			return err
		}

		return err
	})
}

func (s *store) ListOrders(ctx context.Context) (orders []Order1, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &orders, listOrdersQuery)
	})
	if err != nil {
		fmt.Printf("error while fetching Orders %s", err.Error())
	}
	if err == sql.ErrNoRows {
		return orders, ErrOrderNotExist
	}
	return
}
