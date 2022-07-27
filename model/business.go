package model

import (
	"encoding/json"
	"fund_index/db"
	"net/http"
	"sort"
	"time"
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
	code := db.NewFundCodeService()
	Loc, _ := time.LoadLocation("Asia/Shanghai")

	// 格式 2006-01-02, 待转化日期 2021-11-02
	start, _ := time.ParseInLocation("2006-01-02", starttime[0], Loc)
	end, _ := time.ParseInLocation("2006-01-02", endtime[0], Loc)
	failBack := calMaxBack(code.GetFundCodeList(), int(start.Unix()), int(end.Unix()))
	sort.Slice(failBack, func(i, j int) bool {
		return failBack[i].Value<failBack[j].Value
	})
	res := &Resp{
		Status: "200",
		Msg: "",
		Data: failBack,
	}
	json.NewEncoder(w).Encode(res)
}
