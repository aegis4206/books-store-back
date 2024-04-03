package controller

import (
	"books-store/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	// w.Header().Set("Content-Type", "application/json")

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
		uuid, _ := uuid.NewV4()
		sess := &model.Session{
			Session_id: uuid.String(),
			User_name:  user.Email,
			User_id:    user.Id,
		}
		fmt.Println(sess)
		model.AddSession(sess)
		expiration := time.Now().Add(24 * time.Hour)
		cookie := http.Cookie{
			Name:  "user",
			Value: uuid.String(),
			// HttpOnly: true,
			Expires:  expiration,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			// Domain:   "localhost",
		}
		http.SetCookie(w, &cookie)
		cookie2 := http.Cookie{
			Name:     "user_name",
			Value:    user.Email,
			Expires:  expiration,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &cookie2)
		respHandle(w, "登入成功", 200, user)

	} else {
		respHandle(w, "帳號密碼錯誤", 400, data)
	}

}
func Regist(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	// w.Header().Set("Content-Type", "application/json")

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

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		model.DeleteSession(cookieValue)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		respHandle(w, "登出成功", 200, cookieValue)
	}
}

// session check
func SessionCheck(r *http.Request) bool {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		sess, _ := model.GetSessionByID(cookie.Value)
		if sess.User_id > 0 {
			return true
		}
	}
	return false
}
