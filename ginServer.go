package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	mr "mtimer/api/router"
)


func main() {
	router := gin.Default()
	//router.Use(mr.CoreMiddleWare())

	mr.GinRouter(router)

	router.Run(":1238")
}
