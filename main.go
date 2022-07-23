package main

import (
	"fund_index/cron"
	"fund_index/web"
)

func main()  {
	cron.Run()
	web.Run()
}