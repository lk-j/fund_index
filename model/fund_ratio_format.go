package model

import (
	"fund_index/db"
	"sync"
)
type FundRatioFormat struct {
	data map[string][]db.Ratio
}

var (
	once   sync.Once
	single *FundRatioFormat
)
//更新内存数据
func (f FundRatioFormat) UpdateMap(code string,ratioList []db.Ratio)  {
	f.data[code] = ratioList
}
//获取单个股票的数据
func (f FundRatioFormat)  GetFundRatio(code string,starttime string, endtime string)  []db.Ratio {
	var ratioList []db.Ratio
	for _, v := range f.data[code] {
		if v.Date>=starttime && v.Date<endtime {
			ratioList = append(ratioList, v)
		}
	}
	return ratioList
}
func NewFundRatioFormat() *FundRatioFormat {
	once.Do(func() {
		single = &FundRatioFormat{data: make(map[string][]db.Ratio)}
	})
	return single
}