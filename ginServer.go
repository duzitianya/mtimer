package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	mr "mtimer/api/router"
	"mtimer/db/mysql"
)


func main() {
	defer mysql.CloseDB()

	router := gin.Default()
	//router.Use(mr.CoreMiddleWare())

	mr.GinRouter(router)

	router.Run(":1238")


}
