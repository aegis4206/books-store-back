package main

import (
	"books-store/controller"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/regist", controller.Regist)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)

	http.ListenAndServe(":8001", handler)
}
