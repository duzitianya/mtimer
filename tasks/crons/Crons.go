package crons

import (
	"github.com/robfig/cron"
	"mtimer/tasks"
	"log"
)

type CronCallback interface {
	Call(task tasks.MtimerTask)
}

type CronEntity struct {
	Task tasks.MtimerTask
	Callback CronCallback
}

type CronDefaultCallback struct {
}

func (callback CronDefaultCallback) Call(task tasks.MtimerTask) {
	//TODO: implements logic
	isSuccess, err := task.TaskSuccess()
	if err != nil {
		log.Fatal(err)
	}
	if isSuccess {
		delete(tasks.AllTasksMap, tasks.AllTasksMap[task.Id])
	}

}

var cronChan chan int
var CronService *cron.Cron

func StartCronService() bool {
	CronService = cron.New()
	CronService.Start()

	//获取所有需要加入到定时任务服务的MtimerTask
	allTasksMap := tasks.AllTasksMap
	for _, task := range allTasksMap {
		//转换为CronEntity添加进执行列表中
		newCronEntity := CronEntity{Task:task}
		newCronEntity.AddNewCron()
	}

	return true
}

func StopCronService() bool {
	//停止服务时，获取未执行完成


	CronService.Stop()

	return true
}

func (cronEntity *CronEntity) AddNewCron() bool{

	if cronEntity.Callback == nil {
		cronEntity.Callback = CronDefaultCallback{}
	}

	err := CronService.AddFunc(cronEntity.Task.CronTime, func(){
		cronEntity.Callback.Call(cronEntity.Task)

	})
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}


