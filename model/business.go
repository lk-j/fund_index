package model

import (
	"encoding/json"
	"fund_index/db"
	"net/http"
)

type Resp struct {
	Status string
	Msg    string
	Data   interface{}
}

//获取最大回撤
func GetMaxBack(w http.ResponseWriter, r *http.Request)  {
	// 获取请求的参数
	query := r.URL.Query()
	// 参数校验
	starttime, ok := query["starttime"]
	if !ok || starttime == nil {
		res := &Resp{
			Status: "400",
			Msg: "starttime param error",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	endtime, ok := query["endtime"]
	if !ok || endtime == nil {
		res := &Resp{
			Status: "400",
			Msg: "endtime param error",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	code := db.NewFundCode()
	failBack := calMaxBack(code.GetFundCodeList(), starttime[0], endtime[0])
	res := &Resp{
		Status: "200",
		Msg: "",
		Data: failBack,
	}
	json.NewEncoder(w).Encode(res)
}
