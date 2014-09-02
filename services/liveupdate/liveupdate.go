package liveupdate

import (
	"github.com/robfig/cron"
	"services/scraping"
)

func CronLatestNews() {
	c := cron.New()
	c.AddFunc("@every 5m", func() {
			scraping.Scrape()
		})
	c.Start()
}
