package model

import (
	"books-store/utils"
	"time"
)

type Order struct {
	OrderId     string
	CreatedTime time.Time
	TotalCount  int
	TotalAmount int
	State       int // 0 未發貨 1已發貨 2完成交易
	UserId      int
	// OrderItems  []*OrderItem
}

type OrderItem struct {
	OrderItemId int
	Count       int
	Amount      int
	Title       string
	Author      string
	Price       int
	ImgPath     string
	OrderId     string
}

func AddOrder(order *Order) error {
	sqlstr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values($1,$2,$3,$4,$5,$6)"
	_, err := utils.Db.Exec(sqlstr, order.OrderId, order.CreatedTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		return err
	}
	return nil
}

func AddOrderItem(orderItem *OrderItem) error {
	sqlstr := "insert into order_items(count,amount,title,author,price,imgpath,order_id) values($1,$2,$3,$4,$5,$6,$7)"
	_, err := utils.Db.Exec(sqlstr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderId)
	if err != nil {
		return err
	}
	return nil
}

func GetOrders() ([]*Order, error) {
	sqlstr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err := utils.Db.Query(sqlstr)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	for rows.Next() {
		order := &Order{}
		rows.Scan(&order.OrderId, &order.CreatedTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
}

func GetOrdersByUserId(user_id int) ([]*Order, error) {
	sqlstr := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = $1"
	rows, err := utils.Db.Query(sqlstr, user_id)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	for rows.Next() {
		order := &Order{}
		rows.Scan(&order.OrderId, &order.CreatedTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
}
