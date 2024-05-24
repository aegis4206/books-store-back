package main

import (
	"books-store/controller"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")
	r.HandleFunc("/regist", controller.Regist).Methods("POST")
	r.HandleFunc("/books", controller.Books).Methods("GET", "POST", "DELETE", "PUT")
	r.HandleFunc("/bookimg/{Id}", controller.BookImg).Methods("GET")
	r.HandleFunc("/books/{Id}", controller.Books).Methods("DELETE")
	r.HandleFunc("/cart/{bookId}", controller.AddBookToCart).Methods("GET")
	r.HandleFunc("/cart", controller.Cart).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/cart/{cartItemId}", controller.UpdateCartItem).Methods("DELETE", "PATCH")
	r.HandleFunc("/checkout", controller.CheckOut).Methods("POST")
	r.HandleFunc("/orders", controller.GetOrders).Methods("GET")
	r.HandleFunc("/userorders", controller.GetOrdersByUserId).Methods("GET")
	r.HandleFunc("/order/{orderId}", controller.UpdateOrderState).Methods("PATCH")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://192.168.6.87:5173", "http://127.0.0.1:5173", "http://aegis4206.tplinkdns.com:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	handler := c.Handler(r)

	http.ListenAndServe(":8003", handler)
	// http.ListenAndServe(":8001", mux)
}
