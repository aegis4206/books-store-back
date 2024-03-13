package controller

import (
	"encoding/json"
	"net/http"
)

func respHandle(w http.ResponseWriter, msg string, code int, data interface{}) {
	switch code {
	case 200:
		w.WriteHeader(http.StatusOK)
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
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

// func successHandle(){

// }
