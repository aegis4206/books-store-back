package controller

import (
	"books-store/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Books(w http.ResponseWriter, r *http.Request) {
	isLogin := SessionCheck(r)
	if !isLogin {
		respHandle(w, "請重新登入", 500, nil)
		return
	}
	fmt.Println("請求方法", r.Method)
	// w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		books, err := model.GetBooks()
		if err != nil {
			respHandle(w, "資料庫錯誤", 400, books)
			break
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
			break
		}
		respHandle(w, "請求成功", 200, book)
	case "PUT":
		vars := mux.Vars(r)
		Id := vars["Id"]
		fmt.Printf("更新編號%v!", Id)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
		}
		var data *model.Book
		json.Unmarshal(body, &data)
		defer r.Body.Close()
		book, err := model.EditBook(Id, data)
		if err != nil {
			respHandle(w, "格式錯誤", 500, data)
			break
		}

		respHandle(w, "請求成功", 200, book)
	case "DELETE":
		var err error
		var list []string
		vars := mux.Vars(r)
		Id, error := vars["Id"]
		if !error {
			body, _ := io.ReadAll(r.Body)
			json.Unmarshal(body, &list)
			defer r.Body.Close()
			for index, value := range list {
				fmt.Printf("索引：%d，值：%v\n", index, value)
			}
			Ids := strings.Join(list, ",")
			err = model.DeleteBook(Ids)
		} else {
			list = append(list, Id)
			err = model.DeleteBook(Id)
		}
		if err != nil {
			respHandle(w, "SQL錯誤", 400, err)
			break
		}
		respHandle(w, "請求成功", 200, list)
	default:
	}

}
