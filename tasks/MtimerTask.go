package tasks

import (
	"time"
	"mtimer/db/mysql"
	"database/sql"
	"log"
	"fmt"
)

const (
	NewBee = iota
	Loaded
	Interupt
	Success
	Del
)

const (
	table_name = "mtimer_tasks"
	table_fields_insert = "group_id, group_name, biz_id, biz_name, cron_time, status, ip, param, ins_num, excution_time, create_time, update_time"
	table_fields_all = "id, group_id, group_name, biz_id, biz_name, cron_time, status, ip, param, ins_num, excution_time, create_time, update_time"
)


type MtimerTask struct {
	Id int64
	GroupId int
	GroupName string
	BizId int
	BizName string
	CronTime string
	Status int
	Ip string
	Param string
	InsNum int
	ExcutionTime time.Time
	CreateTime time.Time
	UpdateTime time.Time
}

func (task *MtimerTask) CreateNewOneTask() (int64, error) {
	db := mysql.GetDB()

	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO %s(%s) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", table_name, table_fields_insert))
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(task.GroupId, task.GroupName, task.BizId, task.BizName, task.CronTime, task.Status, task.Ip, task.Param, task.InsNum, task.ExcutionTime, task.CreateTime, task.UpdateTime)
	if err != nil {
		return 0, err
	}

	var rows int64
	rows, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if rows > 0 {
		return rows, nil
	}
	return 0, nil
}

func (task *MtimerTask) StopOneTask() (bool, error) {
	if task.Status == NewBee || task.Status == Loaded {
		return updateTaskStatus(Interupt, task.Id)
	}
	return false, nil
}

func (task *MtimerTask) DelOneTask() (bool, error) {
	return updateTaskStatus(Del, task.Id)
}

func (task *MtimerTask) TaskSuccess() (bool, error) {
	return updateTaskStatus(Success, task.Id)
}


func getTaskList(partition int) (map[int64]MtimerTask, error) {
	db := mysql.GetDB()

	rows, err := db.Query("SELECT " + table_fields_all + " FROM " + table_name + " WHERE ins_num=? AND status in (" + fmt.Sprintf("%d, %d", NewBee, Loaded) + ") order by excution_time desc limit 500", partition)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//lists := make([]MtimerTask, 500)
	var taskMap = make(map[int64]MtimerTask)

	for rows.Next() {
		task := MtimerTask{}
		err = rows.Scan(&task.Id, &task.GroupId, &task.GroupName, &task.BizId, &task.BizName, &task.CronTime, &task.Status, &task.Ip, &task.Param, &task.InsNum, &task.ExcutionTime, &task.CreateTime, &task.UpdateTime)
		if err != nil {
			log.Fatal(err)
			continue
		}
		taskMap[task.Id] = task
	}
	rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return taskMap, nil
}

func updateTaskStatus(status int, id int64) (bool, error) {
	db := mysql.GetDB()
	stmt, err := db.Prepare("UPDATE " + table_name + " SET status=? WHERE id=?")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var result sql.Result
	var rows int64
	result, err = stmt.Exec(status, id)
	rows, err = result.RowsAffected()
	if err !=nil {
		return false, err
	}
	if rows > 0 {
		return true, nil
	}

	return false, nil
}
