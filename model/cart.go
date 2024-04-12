package model

import (
	"books-store/utils"
	"strconv"
)

type CartItem struct {
	CartId     string
	CartItemId string
	Book       *Book
	Count      string
	Amount     string
}

type Cart struct {
	CartId      string
	CartItems   []*CartItem
	TotalCount  int
	TotalAmount int
	UserId      int
}

func (cartItem *CartItem) GetAmount() int {
	count, _ := strconv.Atoi(cartItem.Count)
	price, _ := strconv.Atoi(cartItem.Book.Price)
	return count * price
}

func (cart *Cart) GetTotalCount() int {
	var totalCount int
	for _, v := range cart.CartItems {
		count, _ := strconv.Atoi(v.Count)
		totalCount = totalCount + count
	}
	return totalCount
}
func (cart *Cart) GetTotalAmount() int {
	var totalAmount int
	for _, v := range cart.CartItems {
		// amount, _ := strconv.Atoi(v.Amount)
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}

func (cart *Cart) GetTotal(property string) int {
	var total int
	for _, v := range cart.CartItems {
		switch property {
		case "Count":
			count, _ := strconv.Atoi(v.Count)
			total += count
		case "Amount":
			amount, _ := strconv.Atoi(v.Amount)
			total += amount
		}
	}
	return total
}

func AddCartItem(cartItem *CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values($1,$2,$3,$4)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

func AddCart(cart *Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values($1,$2,$3,$4)"
	_, err := utils.Db.Exec(sqlStr, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		return err
	}
	for _, cartItem := range cart.CartItems {
		AddCartItem(cartItem)
	}

	return nil
}

func GetCartItemByBookId(bookId string) (*CartItem, error) {
	sqlStr := "select id,cart_id,count,amount,book_id from cart_items where book_id = $1"
	row := utils.Db.QueryRow(sqlStr, bookId)
	cartItem := &CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.CartId, &cartItem.Count, &cartItem.Amount, &cartItem.Book.Id)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func GetCartItemsByCartId(cartId string) ([]*CartItem, error) {
	sqlStr := "select id,cart_id,count,amount,book_id from cart_items where cart_id = $1"
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		return nil, err
	}
	var cartItems []*CartItem

	for rows.Next() {
		cartItem := &CartItem{}
		err = rows.Scan(&cartItem.CartItemId, &cartItem.CartId, &cartItem.Count, &cartItem.Amount, &cartItem.Book.Id)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func GetCartByUserId(cartId string) (*Cart, error) {
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = $1"
	row := utils.Db.QueryRow(sqlStr, cartId)
	cart := &Cart{}
	err := row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemsByCartId(cart.CartId)
	// for _,item := range cartItems{
	// 	cart.CartItems = append(cart.CartItems, *item)
	// }
	cart.CartItems = cartItems

	return cart, nil
}
