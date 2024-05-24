package controller

import (
	"books-store/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {
	sess := SessionCheck(w, r)
	if sess == nil {
		return
	}
	var list []string
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &list)
	defer r.Body.Close()

	stringUserId := strconv.Itoa(sess.User_id)
	cart, _ := model.GetCartByUserId(stringUserId)

	var checkoutItems []*model.CartItem
	var newCartItems []*model.CartItem
	for _, item := range cart.CartItems {
		for _, id := range list {
			// itemId, _ := strconv.Atoi(item.CartItemId)
			if item.CartItemId == id {
				model.DeleteCartItemByCartItemId(item.CartItemId)

				checkoutItems = append(checkoutItems, item)
			} else {
				newCartItems = append(newCartItems, item)
			}
		}
	}
	cart.CartItems = newCartItems
	model.UpdateCart(cart)

	orderUuid, _ := uuid.NewV4()
	tCount := 0
	tAmount := 0

	//計算購車總和
	for _, item := range checkoutItems {
		count, _ := strconv.Atoi(item.Count)
		amount, _ := strconv.Atoi(item.Amount)
		tCount = tCount + count
		tAmount = tAmount + amount
	}
	order := &model.Order{
		OrderId:     orderUuid.String(),
		CreatedTime: time.Now(),
		TotalCount:  tCount,
		TotalAmount: tAmount,
		State:       0, // 0 未發貨 1已發貨 2完成交易
		UserId:      sess.User_id,
	}
	err := model.AddOrder(order)
	if err != nil {
		respHandle(w, "資料庫錯誤", 400, err)
		return
	}

	for _, item := range checkoutItems {
		count, _ := strconv.Atoi(item.Count)
		amount, _ := strconv.Atoi(item.Amount)
		price, _ := strconv.Atoi(item.Book.Price)

		oderItem := &model.OrderItem{
			Count:   count,
			Amount:  amount,
			Title:   item.Book.Title,
			Author:  item.Book.Author,
			Price:   price,
			ImgPath: item.Book.ImgPath,
			OrderId: orderUuid.String(),
			BookId:  item.Book.Id,
		}
		err := model.AddOrderItem(oderItem)
		if err != nil {
			respHandle(w, "資料庫錯誤", 400, err)
			return
		}

		book := item.Book
		bookSales, _ := strconv.Atoi(book.Sales)
		bookStock, _ := strconv.Atoi(book.Stock)
		sales := strconv.Itoa(bookSales + count)
		stock := strconv.Itoa(bookStock - count)
		book.Sales = sales
		book.Stock = stock
		model.EditBook(book)
	}
	respHandle(w, "成功建立訂單", 200, checkoutItems)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := model.GetOrders()
	if err != nil {
		respHandle(w, "資料庫錯誤", 400, err)
		return
	}
	respHandle(w, "請求成功", 200, orders)
}

func GetOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	sess := SessionCheck(w, r)
	if sess == nil {
		return
	}
	orders, err := model.GetOrdersByUserId(sess.User_id)
	if err != nil {
		respHandle(w, "資料庫錯誤", 400, err)
		return
	}
	respHandle(w, "請求成功", 200, orders)
}

func UpdateOrderState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, err := vars["orderId"]
	if !err {
		respHandle(w, "訂單號碼錯誤", 400, err)
		return
	}
	fmt.Println(orderId)
	body, _ := io.ReadAll(r.Body)
	var data *struct {
		State string
	}
	json.Unmarshal(body, &data)
	defer r.Body.Close()
	fmt.Println(data.State)
	sqlErr := model.UpdateOrderState(orderId, data.State)
	if sqlErr != nil {
		respHandle(w, "資料庫錯誤", 400, err)
		return
	}
	respHandle(w, "訂單狀態已更新", 200, orderId)
}
