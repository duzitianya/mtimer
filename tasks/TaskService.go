package tasks

import (
	"time"
	"fmt"
	"log"
	"github.com/robfig/cron"
)

var AllTasksMap map[int64]MtimerTask
var InsNum int
var CronService *cron.Cron

type CronCallback interface {
	Call(task MtimerTask)
}

type CronEntity struct {
	Task MtimerTask
	Callback CronCallback
}

type CronDefaultCallback struct {
}

func (callback CronDefaultCallback) Call(task MtimerTask) {
	isSuccess, err := task.TaskSuccess()
	if err != nil {
		log.Fatal(err)
	}
	if isSuccess {
		delete(AllTasksMap, AllTasksMap[task.Id])
	}

}

func StartCronService(sd <-chan int) {
	CronService = cron.New()
	CronService.Start()

	var err error
	AllTasksMap, err = getTaskList(InsNum)
	if err != nil {
		log.Fatal(err)
	}
	if len(AllTasksMap) > 0 {
		//获取所有需要加入到定时任务服务的MtimerTask
		allTasksMap := AllTasksMap
		for _, task := range allTasksMap {
			//转换为CronEntity添加进执行列表中
			newCronEntity := CronEntity{Task:task}
			newCronEntity.addNewCron()
		}

	}

	select{
	case <-sd : stopCronService()
	}

}

func stopCronService() {
	//停止服务时，获取未执行完成
	CronService.Stop()
	AllTasksMap = nil
}

func (cronEntity *CronEntity) addNewCron() bool{
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


func NewTaskReceived(groupId, bizId int, groupName, bizName, ip, param string, excutionTime time.Time) bool {
	task := MtimerTask{}
	task.GroupId = groupId
	task.GroupName = groupName
	task.BizId = bizId
	task.BizName = bizName
	task.Param = param
	task.Ip = ip

	if excutionTime.Before(time.Now()) {
		return false
	}

	cronFormat := " 0 %d %d %d %d ? "
	_, month, day := excutionTime.Date()
	hour := excutionTime.Hour()
	minute := excutionTime.Minute()
	cron := fmt.Sprintf(cronFormat, minute, hour, day, month)
	task.CronTime = cron
	task.ExcutionTime = excutionTime

	task.InsNum = InsNum

	task.CreateTime = time.Now().UTC()

	result, err := task.CreateNewOneTask()
	if err != nil {
		log.Fatal(err)
	}

	//如果数据库创建成功，且执行时间小于一天，则放入内存中
	if result > 0 && (excutionTime.Unix() - task.CreateTime.Unix())/3600 < 24 {
		AllTasksMap[result] = task
		cronBean := CronEntity{ Task:task }
		return cronBean.addNewCron()
	}


	return false
}

