package model

import (
	"fund_index/db"
	"strconv"
)

type FailBack struct {
	Code string
	Value float64
}
//计算最大回撤
func calMaxBack(codeList []db.FundCode, starttime string, endtime string) []FailBack {
	var format = NewFundRatioFormat()
	var failBackList  []FailBack
	for _, v := range codeList {
		ratioList:= format.GetFundRatio(v.Code, starttime, endtime)
		failBackList = append(failBackList, calValue(ratioList, v.Code))
	}
	return failBackList
}
//计算公式
func calValue(ratioList []db.Ratio, code string)  FailBack {
	var fallback = FailBack{
		Code: code,
		Value: 0,
	}
	if ratioList == nil {
		return fallback
	}
	 dp := make([]float64, len(ratioList))
	 dp[0] = 0
	max := ratioList[0].Value
	for i := 1; i < len(ratioList); i++ {
		if (ratioList[i].Value>max) {
			dp[i] = dp[i-1]
			max = ratioList[i].Value
		} else {
			chu,_ := strconv.ParseFloat(ratioList[i].Value, 8)
			beichu,_ := strconv.ParseFloat(max, 8)
			tmp := (chu-beichu)/beichu
			 if (tmp<dp[i-1]) {
				 dp[i] = tmp
			 }
		}
	}
	fallback.Value = dp[len(ratioList)-1]
	return fallback
}

