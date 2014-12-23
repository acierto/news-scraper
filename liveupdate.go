package main

import (
	"github.com/robfig/cron"
)

func CronLatestNews() {
	c := cron.New()
	c.AddFunc("@every 5m", func() {
			Scrape()
		})
	c.Start()
}
