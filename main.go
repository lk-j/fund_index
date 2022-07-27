package main

import (
	"fund_index/cron"
	"fund_index/web"
)

func main() {
	go cron.Run()
	web.Run()
}
