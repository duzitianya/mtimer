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


	defer mysql.CloseDB()


	router := gin.Default()
	//router.Use(mr.CoreMiddleWare())

	mr.GinRouter(router)

	router.Run(":1238")

	/*c := cron.New()
	c.AddFunc("0 28 18 9 8 ?", func() {
		fmt.Println("=================")
	})
	c.AddFunc("0 30 18 9 8 ?", func() {
		fmt.Println("=================")
	})
	c.Start()

	c.AddFunc("0 32 18 9 8 ?", func() {
		fmt.Println("=================")
	})

	select {

	}*/


}
