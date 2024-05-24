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

func GetCartItemByBookIdAndCartId(bookId string, cartId string) (*CartItem, error) {
	sqlStr := "select id,cart_id,count,amount from cart_items where book_id = $1 and cart_id = $2"
	row := utils.Db.QueryRow(sqlStr, bookId, cartId)
	cartItem := &CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.CartId, &cartItem.Count, &cartItem.Amount)
	if err != nil {
		return nil, err
	}
	book := GetBookById(bookId)
	cartItem.Book = book
	return cartItem, nil
}

// bookId string, cartId string, bookCount int, bookAmount int
func UpdateCartItemByBookIdAndCartId(cartItem *CartItem) error {
	sqlStr := "update cart_items set count = $1,amount = $2 where book_id = $3 and cart_id = $4"
	// _, err := utils.Db.Exec(sqlStr, bookCount, bookAmount, bookId, cartId)
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemsByCartId(cartId string) ([]*CartItem, error) {
	sqlStr := "select id,cart_id,count,amount,book_id from cart_items where cart_id = $1 order by id asc"
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		return nil, err
	}
	var cartItems []*CartItem

	for rows.Next() {
		cartItem := &CartItem{}
		var bookId string
		//需將book_id關連到完整的book
		err = rows.Scan(&cartItem.CartItemId, &cartItem.CartId, &cartItem.Count, &cartItem.Amount, &bookId)
		if err != nil {
			return nil, err
		}
		book := GetBookById(bookId)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func GetCartByUserId(userId string) (*Cart, error) {
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = $1 order by id desc"
	row := utils.Db.QueryRow(sqlStr, userId)
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

func UpdateCart(cart *Cart) error {
	sqlStr := "update carts set total_count=$1,total_amount=$2 where id=$3"
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCartItemsByCartId(cartId string) error {
	sqlStr := "delete from cart_items where cart_id = $1"
	_, err := utils.Db.Exec(sqlStr, cartId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartItemByCartItemId(cartItemId string) error {
	sqlStr := "delete from cart_items where id = $1"
	_, err := utils.Db.Exec(sqlStr, cartItemId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCartItemByUserId(UserId string) []*CartItem {
	var list []*CartItem
	return list
}
