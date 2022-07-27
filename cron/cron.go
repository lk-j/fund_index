package cron

import (
	"fmt"
	"github.com/robfig/cron"
)

func Run()  {
	loadFundRatio()
	c := cron.New()
	c.AddFunc("0 0 */1 * *", loadFundRatio)
	c.Start()
	fmt.Println("cron run exit")
}
