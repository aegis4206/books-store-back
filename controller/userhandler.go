package controller

import (
	"books-store/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"strconv"

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
			Name:  "SessionId",
			Value: uuid.String(),
			// HttpOnly: true,
			Expires:  expiration,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			// Domain:   "localhost",
		}
		http.SetCookie(w, &cookie)
		cookie2 := http.Cookie{
			Name:     "Email",
			Value:    user.Email,
			Expires:  expiration,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &cookie2)
		cookie3 := http.Cookie{
			Name:     "Id",
			Value:    strconv.Itoa(user.Id),
			Expires:  expiration,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &cookie3)
		user.SessionId = sess.Session_id
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
	var cookieDelete = CookieDelete(w, r)
	cookie, err := r.Cookie("SessionId")
	if err != nil {
		respHandle(w, "登出成功", 200, "無SessionId")
		return
	}
	cookieValue := cookie.Value
	model.DeleteSession(cookieValue)
	cookieDelete("SessionId")
	cookieDelete("Email")
	cookieDelete("Id")
	respHandle(w, "登出成功", 200, cookieValue)
}

// session check
func SessionCheck(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("SessionId")
	if err != nil {
		respHandle(w, "請重新登入", 400, nil)
		return false
	}
	sess, err := model.GetSessionByID(cookie.Value)
	if err != nil {
		var cookieDelete = CookieDelete(w, r)
		cookieDelete("SessionId")
		cookieDelete("Email")
		cookieDelete("Id")
		respHandle(w, "資料庫無seionss，請重新登入", 400, nil)
		return false
	}
	if sess.User_id > 0 {
		return true
	}
	return false
}

// cookie註銷
func CookieDelete(w http.ResponseWriter, r *http.Request) func(target string) {
	return func(target string) {
		cookie, err := r.Cookie(target)
		if err != nil {
			return
		}
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}
