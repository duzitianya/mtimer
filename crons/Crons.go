package crons

import "github.com/robfig/cron"

type CronEntity struct {
	cronExp string
}

func (ce *CronEntity) start() {
	c := cron.New()
	c.AddFunc(" * * * * * * ", func(){})
	c.Start()
}
