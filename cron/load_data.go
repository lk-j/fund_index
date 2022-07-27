package cron

import (
	"fmt"
	"fund_index/db"
	"fund_index/model"
	"sync"
)

type param struct {
	code string
	fundRatio *db.FundRatio
	fundRatioFormat   *model.FundRatioFormat
}
//载入基金净值数据到内存
func  loadFundRatio()  {
	//获取基金列表
	fundList := db.NewFundCodeService()
	fundRatio := db.NewFundRatio()
	fundRatioFormat := model.NewFundRatioFormat()
	//遍历基金列表，获取净值数据，写入内存
	list :=fundList.GetFundCodeList()
	var count int
	for _, item := range list {
		if count>100 {
			break
		}
		count++
		code := item.Code
		updateMap(code,fundRatio,fundRatioFormat)
	}
}
var  Lock sync.RWMutex
func updateMap(code string,fundRatio *db.FundRatio, fundRatioFormat *model.FundRatioFormat)  {
	Lock.Lock()
	defer Lock.Unlock()
	ratioList := fundRatio.GetFundRatioList(code)
	if code != "" && ratioList !=nil {
		fundRatioFormat.UpdateMap(code, ratioList)
	}
	fmt.Println(code)
}
