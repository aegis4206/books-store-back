package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int
	Msg  string
	Data interface{} // interface{} 等於any
}

func respHandle(w http.ResponseWriter, msg string, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	switch code {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 302:
		w.WriteHeader(302)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 404:
		w.WriteHeader(http.StatusNotFound)
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusOK)
	}
	resp := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	jsonResp, _ := json.Marshal(resp) // json.Marshal 一般型態或指針都可處理
	w.Write(jsonResp)
}

// func successHandle(){

// }
