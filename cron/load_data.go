package cron

import (
	"fmt"
	"fund_index/db"
	"fund_index/model"
	"github.com/panjf2000/ants/v2"
)
var pool *ants.PoolWithFunc
func init()  {
	pool, _ = ants.NewPoolWithFunc(5, updateMap)
}

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
	for _, item := range list {
		code := item.Code
		 params := param{
			code:code,fundRatio:fundRatio,fundRatioFormat:fundRatioFormat,
		}
		pool.Invoke(params)
	}
}
func updateMap(data interface{})  {
	params := data.(param)
	ratioList := params.fundRatio.GetFundRatioList(params.code)
	if params.code != "" && ratioList !=nil {
		params.fundRatioFormat.UpdateMap(params.code, ratioList)
	}
	fmt.Println(params.code)
}
