package controller

import (
	"books-store/model"
	"encoding/json"
	"fmt"
	"io"
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
			addCount := strconv.Itoa(count + 1)
			// price, _ := strconv.Atoi(cartItem.Book.Price)
			// amount := (count + 1) * price
			// model.UpdateCartItemByBookIdAndCartId(bookId, currentCart.CartId, count+1, amount)
			cartItem.Count = addCount
			model.UpdateCartItemByBookIdAndCartId(cartItem)

			cts := currentCart.CartItems
			for _, v := range cts {
				if v.Book.Id == bookId {
					v.Count = addCount
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

func Cart(w http.ResponseWriter, r *http.Request) {
	sess := SessionCheck(w, r)
	if sess == nil {
		return
	}
	stringUserId := strconv.Itoa(sess.User_id)
	cart, _ := model.GetCartByUserId(stringUserId)

	switch r.Method {
	case "GET":
		list := []*model.CartItem{}
		if cart == nil {
			respHandle(w, "尚未建立購物車", 200, list)
			return
		}
		cartItems, _ := model.GetCartItemsByCartId(cart.CartId)
		if cartItems == nil {
			respHandle(w, "尚未建立購物項", 200, list)
			return
		}
		respHandle(w, "成功取得購物車", 200, cartItems)
		// case "PATCH":
		// vars := mux.Vars(r)
		// cartItemId, err := vars["cartItemId"]
		// fmt.Println("cartItemId", cartItemId)
		// if err {
		// 	body, _ := io.ReadAll(r.Body)
		// 	var count string
		// 	json.Unmarshal(body, &count)
		// 	defer r.Body.Close()
		// 	fmt.Println("count", count)
		// 	for _, value := range cart.CartItems {
		// 		if value.CartItemId == cartItemId {
		// 			value.Count = count
		// 			err := model.UpdateCartItemByBookIdAndCartId(value)
		// 			if err != nil {
		// 				respHandle(w, "資料庫錯誤", 400, err)
		// 				return
		// 			}
		// 		}
		// 	}
		// 	err := model.UpdateCart(cart)
		// 	if err != nil {
		// 		respHandle(w, "資料庫錯誤", 400, err)
		// 		return
		// 	}
		// 	respHandle(w, "成功更新購物項數量", 200, cartItemId)
		// }
		// respHandle(w, "參數錯誤", 400, err)
	// case "PUT":

	case "DELETE":
		err := model.DeleteCartItemsByCartId(cart.CartId)

		// 清除購物項
		emptyCartItem := []*model.CartItem{}
		cart.CartItems = emptyCartItem
		model.UpdateCart(cart)

		fmt.Println(cart.CartId)
		if err != nil {
			respHandle(w, "資料庫錯誤", 400, err)
			return
		}
		respHandle(w, "成功清空購物車", 200, cart.CartId)
	default:
	}

}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	sess := SessionCheck(w, r)
	if sess == nil {
		return
	}

	vars := mux.Vars(r)
	cartItemId, err := vars["cartItemId"]

	stringUserId := strconv.Itoa(sess.User_id)
	cart, _ := model.GetCartByUserId(stringUserId)

	switch r.Method {
	case "PATCH":
		fmt.Println("cartItemId", cartItemId)
		if err {
			body, _ := io.ReadAll(r.Body)
			var data *struct {
				Count int
			}
			json.Unmarshal(body, &data)
			defer r.Body.Close()
			fmt.Println("count", data.Count)
			for _, value := range cart.CartItems {
				if value.CartItemId == cartItemId {
					value.Count = strconv.Itoa(data.Count)
					err := model.UpdateCartItemByBookIdAndCartId(value)
					if err != nil {
						respHandle(w, "資料庫錯誤", 400, err)
						return
					}
				}
			}
			err := model.UpdateCart(cart)
			if err != nil {
				respHandle(w, "資料庫錯誤", 400, err)
				return
			}
			respHandle(w, "成功更新購物項數量", 200, cartItemId)
			return
		}
		respHandle(w, "參數錯誤", 400, err)
	// case "PUT":
	case "DELETE":
		if err {
			err := model.DeleteCartItemByCartItemId(cartItemId)
			if err != nil {
				respHandle(w, "資料庫錯誤", 400, err)
				return
			}
			for k, v := range cart.CartItems {
				if v.CartItemId == cartItemId {
					cart.CartItems = append(cart.CartItems[:k], cart.CartItems[k+1:]...)
				}
			}
			model.UpdateCart(cart)
			respHandle(w, "成功刪除購物項", 200, cartItemId)
			return
		}
		respHandle(w, "參數錯誤", 400, err)
	default:
	}

}
