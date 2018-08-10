package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	mr "mtimer/api/router"
	"mtimer/db/mysql"
	"os"
	"mtimer/tasks"
	"strconv"
)


func main() {
	args := os.Args
	if len(args) == 1 {
		tasks.InsNum = 1
	} else {
		inum, err := strconv.Atoi(args[1])
		if err != nil {
			tasks.InsNum = 1
		} else {
			tasks.InsNum = inum
		}
	}

	var shutdownCronService chan int
	go tasks.StartCronService(shutdownCronService)

	defer func() {
		shutdownCronService <- 1
	}()
	defer mysql.CloseDB()


	router := gin.Default()
	//router.Use(mr.CoreMiddleWare())

	mr.GinRouter(router)

	router.Run(":1238")

}
