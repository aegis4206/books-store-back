package controller

import (
	"books-store/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")

	var data model.User
	err = json.Unmarshal(body, &data)
	if err != nil {
		respHandle(w, "格式錯誤", 500, body)
		return
	}
	fmt.Println(data.Password)
	fmt.Println(data.Email)
	defer r.Body.Close()

	user, _ := model.CheckEmailAndPassword(data.Email, data.Password)
	if user.Id > 0 {
		respHandle(w, "登入成功", 200, data)

	} else {
		respHandle(w, "帳號密碼錯誤", 400, data)
	}

}
func Regist(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")

	var data model.User
	err = json.Unmarshal(body, &data)
	if err != nil {
		respHandle(w, "格式錯誤", 500, body)
		return
	}

	user, _ := model.CheckEmail(data.Email)
	if user.Id > 0 {
		respHandle(w, "註冊失敗信箱已存在", 400, data)
	} else {
		model.SaveUser(data.Email, data.Password)
		respHandle(w, "註冊成功，請登入", 200, data)
	}

}
