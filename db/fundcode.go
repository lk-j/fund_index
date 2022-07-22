package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionFundCode   *mongo.Collection

func init() {
	collectionFundCode = Connect().Database("stock").Collection("fundcode")
}
type FundCode struct {
	Name string
	Code  string
}

//获取基金列表
func GetFundCodeList()  []FundCode {
	filter:= bson.D{}
	cursor, _ := collectionFundCode.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	var fundCode []FundCode
	cursor.All(context.TODO(), &fundCode)
	return fundCode
}

