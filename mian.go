package main

import (
	"github.com/lizongying/cron"
)

func main() {
	logger := cron.NewLoggerStdout()
	c := cron.New(cron.WithStdout())
	c.MustStart()
	id := c.MustAddJob(new(cron.Job).
		EverySecond(10).
		MustSince("10:15").
		Callback(func() {
			logger.Info("callback")
		}))
	logger.Info("id", id)
	select {}
}
