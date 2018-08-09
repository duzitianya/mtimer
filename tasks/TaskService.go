package tasks

import (
	"time"
	"fmt"
	"log"
	"mtimer/tasks/crons"
)

var AllTasks []MtimerTask
var insNum int

func init() {
	//TODO: 获取该实例编号
	insNum = 1

	var err error
	AllTasks, err = getTaskList(insNum)
	if err != nil {
		log.Fatal(err)
	}
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

	task.InsNum = insNum

	task.CreateTime = time.Now().UTC()

	result, err := task.CreateNewOneTask()
	if err != nil {
		log.Fatal(err)
	}

	//如果数据库创建成功，且执行时间小于一天，则放入内存中
	if result > 0 && (excutionTime.Unix() - task.CreateTime.Unix())/3600 < 24 {
		AllTasks = append(AllTasks, task)
		//TODO: 添加到cron执行
		cronBean := crons.CronEntity{ CronExp:task.CronTime }
		cronBean.StartNewCronWithoutCallBack()

		return true
	}


	return false
}

