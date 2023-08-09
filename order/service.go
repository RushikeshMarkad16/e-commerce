package order

import (
	"context"
	"fmt"
	"time"

	"github.com/RushikeshMarkad16/e-commerce/db"
	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req Order) (err error)
	List(ctx context.Context) (response ListResponse, err error)
	FindByID(ctx context.Context, id int) (response FindByIdResponse, err error)
	Update(ctx context.Context, req OrderStatus) (err error)
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
		orders = append(orders, &db.Order_item{
			Product_id: v.Product_Id,
			Quantity:   v.Quantity,
		})
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
		if det.Quantity > 10 {
			return errGreaterthanTen
		}

	}
	return
}

func (cs *orderService) List(ctx context.Context) (response ListResponse, err error) {
	dbOrders, err := cs.store.ListOrders(ctx)
	if err == db.ErrOrderNotExist {
		cs.logger.Error("No order present", "err", err.Error())
		return response, errNoOrders
	}
	if err != nil {
		cs.logger.Error("Error listing orders", "err", err.Error())
		return
	}

	for _, dbOrder := range dbOrders {
		var orderData Order1

		orderData.ID = dbOrder.ID
		orderData.Amount = dbOrder.Amount
		orderData.Disc_perc = dbOrder.Disc_perc
		orderData.Final_amnt = dbOrder.Final_amnt
		if dbOrder.Disp_date.Valid {
			date, _ := time.Parse("2006-01-02", dbOrder.Disp_date.String)
			fmt.Print(date, dbOrder.Disp_date)
			orderData.Disp_date = date.String()
		}
		orderData.Order_status = dbOrder.Order_status

		response.Orders = append(response.Orders, orderData)
	}

	return
}

func (cs *orderService) FindByID(ctx context.Context, id int) (response FindByIdResponse, err error) {
	order, err := cs.store.FindOrderByID(ctx, id)
	if err == db.ErrOrderNotExist {
		cs.logger.Error("No order present", "err", err.Error())
		return response, errNoOrderId
	}
	if err != nil {
		cs.logger.Error("Error finding order", "err", err.Error(), "id", id)
		return
	}
	response.Order.ID = order.ID
	response.Order.Amount = order.Amount
	response.Order.Disc_perc = order.Disc_perc
	response.Order.Final_amnt = order.Final_amnt
	if order.Disp_date.Valid {
		date, _ := time.Parse("2006-01-02", order.Disp_date.String)
		fmt.Print(date, order.Disp_date)
		response.Order.Disp_date = date.String()
	}
	response.Order.Order_status = order.Order_status

	return
}

func (cs *orderService) Update(ctx context.Context, c OrderStatus) (err error) {
	err = c.ValidateUpdate()
	if err != nil {
		cs.logger.Error("Invalid Request for order status update", "err", err.Error(), "orderstatus", c)
		return
	}

	err = cs.store.UpdateOrderStatus(ctx, &db.Order1{
		ID:           c.Order_id,
		Order_status: c.Status,
	})
	if err != nil {
		cs.logger.Error("Error updating order", "err", err.Error(), "orderstatus", c)
		return
	}
	return
}

func (up OrderStatus) ValidateUpdate() (err error) {
	if up.Order_id == 0 {
		return errEmptyOrderId
	}

	if up.Status == "" {
		return errEmptyStatus
	}
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &orderService{
		store:  s,
		logger: l,
	}
}
