package order

import (
	"context"

	"github.com/RushikeshMarkad16/e-commerce/db"
	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req Order) (err error)
}

type orderService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *orderService) create(ctx context.Context, c Order) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for order create", "msg", err.Error(), "order", c)
		return
	}
	var orders []*db.Order_item
	for _, v := range c.Item {
		// err = cs.store.CreateOrder(ctx, &db.Order_item{
		// 	Product_id: v.Product_Id,
		// 	Quantity:   v.Quantity,
		// })
		orders = append(orders, &db.Order_item{
			Product_id: v.Product_Id,
			Quantity:   v.Quantity,
		})
		// err = cs.store.CreateOrder(ctx, orders[])
		// if err != nil {
		// 	cs.logger.Error("Error creating order", "err", err.Error())
		// 	return
		// }
	}
	err = cs.store.CreateOrder(ctx, orders)
	if err != nil {
		cs.logger.Error("Error creating order", "err", err.Error())
		return
	}
	return
}

func (cr Order) Validate() (err error) {
	for _, det := range cr.Item {
		if det.Product_Id == 0 {
			return errEmptyProductID
		}
		if det.Quantity == 0 {
			return errEmptyQuantity
		}

	}
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &orderService{
		store:  s,
		logger: l,
	}
}
