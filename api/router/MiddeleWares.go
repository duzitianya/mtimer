package router

import "gopkg.in/gin-gonic/gin.v1"
import "fmt"

func CoreMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("before middleware")
		c.Set("request", "from CoreMiddleWare")
		//c.Next()
		fmt.Println("after middleware")
	}
}

func LoginAuthentication() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}