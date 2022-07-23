package db

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)


type FundCode struct {
	Name string
	Code  string
}

//获取基金列表
func (fund *FundCode) GetFundCodeList()  []FundCode {
	filter:= bson.D{}
	var fundCode []FundCode
	db := getDB()
	c := db.C("fundcode")
	if err := c.Find(filter).All(&fundCode); err !=nil {
		log.Println("GetFundCodeList error err:%s", err.Error())
	}
	return fundCode
}

func NewFundCode()  *FundCode {
	return &FundCode{}
}
