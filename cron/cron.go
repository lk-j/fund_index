package cron

import (
	"github.com/robfig/cron"
)

func Run()  {
	loadFundRatio()
	c := cron.New()
	c.AddFunc("* * * * *", loadFundRatio)
	c.Start()
}
