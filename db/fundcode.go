package db

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"sync"
)


type FundCodeService struct {
	data []FundCode
}
type FundCode struct {
	Name string
	Code  string
}
var (
	once   sync.Once
	single *FundCodeService
)
//获取基金列表
func (fund *FundCodeService) GetFundCodeList()  []FundCode {
	if fund.data !=nil {
		return fund.data
	}
	filter:= bson.D{}
	var fundCode []FundCode
	db := getDB()
	c := db.C("fundcode")
	if err := c.Find(filter).All(&fundCode); err !=nil {
		log.Printf("GetFundCodeList error err:%s\n", err.Error())
	}
	fund.data = fundCode
	return fundCode
}

func NewFundCodeService()  *FundCodeService {
	once.Do(func() {
		single = &FundCodeService{}
	})
	return single
}
