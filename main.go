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
	r.HandleFunc("/regist", controller.Regist).Methods("POST")
	r.HandleFunc("/books", controller.Books).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/books/{Id}", controller.Books).Methods("PUT", "DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(r)

	http.ListenAndServe(":8001", handler)
	// http.ListenAndServe(":8001", mux)
}
