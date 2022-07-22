package cron

import (
	"fmt"
	"fund_index/db"
)

//载入基金净值数据到内存
func loadFundRatio()  {
	//获取基金列表
	fundList := db.GetFundCodeList()
	//遍历基金列表，获取净值数据，写入内存
	for _, item := range fundList {
		code := item.Code
		fmt.Println(code)
	}

}