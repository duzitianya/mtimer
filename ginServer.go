package main

import (
	//"gopkg.in/gin-gonic/gin.v1"
	//mr "mtimer/api/router"
	"mtimer/db/mysql"
	"mtimer/tasks"
	"time"
	"log"
)


func main() {
	defer mysql.CloseDB()

	task := tasks.MtimerTask{}
	task.GroupId = 1
	task.GroupName = "test"
	task.BizId = 1
	task.BizName = "test"
	task.CronTime = " * * * 1 * * "
	task.Ip = "192.168.100.123"
	task.Status = tasks.NewBee
	task.InsNum = 1
	task.Param = "";
	task.ExcutionTime = time.Now()
	task.CreateTime = time.Now()
	task.UpdateTime = time.Now()

	result, err := task.CreateNewOneTask()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" insert success %d \n", result)

	/*
	router := gin.Default()
	//router.Use(mr.CoreMiddleWare())

	mr.GinRouter(router)

	router.Run(":1238")*/


}
