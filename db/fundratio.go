package db

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type FundRatio struct {
}


type Ratio struct {
	Code  string `bson:"code"`
	Date string  `bson:"date"`
	Value string `bson:"value"`
}
//获取基金净值列表 by code
func (fund FundRatio) GetFundRatioList(code string) []Ratio {

	var list  []Ratio
	db := getDB()

	c := db.C("fundratio")
	filter := bson.M{"code": code}
	if err := c.Find(filter).All(&list);err !=nil {
		log.Println("GetFundRatioList error err:%s", err.Error())
		return list
	}
	return list
}
func NewFundRatio() *FundRatio {
	return &FundRatio{}
}