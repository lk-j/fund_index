package web

import (
	"fund_index/model"
	"net/http"
)

func Run()  {
	http.HandleFunc("/getMaxBack", model.GetMaxBack)
	http.ListenAndServe("127.0.0.1:8080",nil)
}

