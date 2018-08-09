package crons

import (
	"github.com/robfig/cron"
)

type CronEntity struct {
	CronExp string
	Callback CronCallback
}

type CronCallback interface {
	Call()
}

type CronDefaultCallback struct {
}

func (callback *CronDefaultCallback) Call() {
	//TODO: implements logic

}

func (cronEntity *CronEntity) StartNewCron() bool{

	if cronEntity.Callback == nil {
		cronEntity.Callback = CronDefaultCallback{}
	}

	c := cron.New()
	c.AddFunc(cronEntity.CronExp, func(){
		cronEntity.Callback.Call()
	})
	c.Start()

	return true
}
