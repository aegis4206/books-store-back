package controller

import (
	"books-store/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	book := model.GetBookById(bookId)
	fmt.Println(book)
	sess := SessionCheck(w, r)
	if sess == nil {
		return
	}
	stringUserId := strconv.Itoa(sess.User_id)
	currentCart, _ := model.GetCartByUserId(stringUserId)
	if currentCart != nil {
		cartItem, _ := model.GetCartItemByBookIdAndCartId(bookId, currentCart.CartId)
		if cartItem != nil {
			count, _ := strconv.Atoi(cartItem.Count)
			model.UpdateCartItemByBookIdAndCartId(bookId, currentCart.CartId, count+1)
			cts := currentCart.CartItems
			for _, v := range cts {
				if v.Book.Id == bookId {
					v.Count = strconv.Itoa(count + 1)
				}
			}
		} else {
			newCartItem := &model.CartItem{
				Book:   book,
				Count:  "1",
				CartId: currentCart.CartId,
			}
			model.AddCartItem(newCartItem)
			currentCart.CartItems = append(currentCart.CartItems, newCartItem)
		}

		model.UpdateCart(currentCart)

	} else {
		uuid, _ := uuid.NewV4()
		cart := &model.Cart{
			CartId: uuid.String(),
			UserId: sess.User_id,
		}
		// var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  "1",
			CartId: uuid.String(),
		}
		// cartItems = append(cartItems, cartItem)
		cart.CartItems = append(cart.CartItems, cartItem)
		model.AddCart(cart)
	}
	respHandle(w, "加入購物車成功", 200, book)

}
