package cron

import (
	"fund_index/db"
	"fund_index/model"
)

//载入基金净值数据到内存
func  loadFundRatio()  {
	//获取基金列表
	fundList := db.NewFundCode()
	fundRatio := db.NewFundRatio()
	fundRatioFormat := model.NewFundRatioFormat()
	//遍历基金列表，获取净值数据，写入内存
	for _, item := range fundList.GetFundCodeList() {
		code := item.Code
		ratioList := fundRatio.GetFundRatioList(code)
		if code != "" && ratioList !=nil {
			fundRatioFormat.UpdateMap(code, ratioList)
		}

	}
}
