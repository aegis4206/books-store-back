package controller

import (
	"books-store/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Books(w http.ResponseWriter, r *http.Request) {
	fmt.Println("請求方法", r.Method)
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		books, err := model.GetBooks()
		if err != nil {
			respHandle(w, "資料庫錯誤", 400, books)
		}
		remoteAddr := r.RemoteAddr
		fmt.Println("Request from:", remoteAddr)
		respHandle(w, "請求成功", 200, books)
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
		}
		var data *model.Book
		json.Unmarshal(body, &data)
		defer r.Body.Close()
		book, err := model.AddBook(data)
		if err != nil {
			respHandle(w, "格式錯誤", 500, data)
			return
		}
		respHandle(w, "請求成功", 200, book)
	case "PUT":
		vars := mux.Vars(r)
		Id := vars["Id"]
		s := fmt.Sprintf("更新編號%v!", Id)
		fmt.Println(Id, s)
		respHandle(w, "請求成功", 200, s)
	case "DELETE":
	default:
	}

	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(500)
	// }
	// w.Header().Set("Content-Type", "application/json")

	// var data model.User
	// err = json.Unmarshal(body, &data)
	// if err != nil {
	// 	respHandle(w, "格式錯誤", 500, body)
	// 	return
	// }
	// fmt.Println(data.Password)
	// fmt.Println(data.Email)
	// defer r.Body.Close()

	// user, _ := model.CheckEmailAndPassword(data.Email, data.Password)
	// if user.Id > 0 {
	// 	respHandle(w, "登入成功", 200, data)

	// } else {
	// 	respHandle(w, "帳號密碼錯誤", 400, data)
	// }

}
