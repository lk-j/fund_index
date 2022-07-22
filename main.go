package main

import (
	"fund_index/cron"
)

func main()  {
	cron.Run()
	select {}
}