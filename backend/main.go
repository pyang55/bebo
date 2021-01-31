package main

import (
	"gopkg.in/robfig/cron.v2"
)

var c = cron.New()

func main() {
	cronWatch()
}
