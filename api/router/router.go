package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	"mtimer/api/module"
	"log"
	"net/http"
	"strings"
	"fmt"
)

func GinRouter(router *gin.Engine) {

	greekRouterGroup := router.Group("/greek", CoreMiddleWare())

	router.Use(CoreMiddleWare())

	coreRouterGroup := router.Group("/core")

	GreekRouterGroup(greekRouterGroup)
	CoreRouterGroup(coreRouterGroup)

}

func GreekRouterGroup(gp *gin.RouterGroup) {
	gp.GET("/:acct", func(c *gin.Context){
		middleStr, exsist := c.Get("request")
		if exsist {
			log.Printf("get message from middle ware : %s", middleStr)
		}

		acct := c.Param("acct")
		name := c.Query("name")
		log.Print(" get param ACCT is : " + acct)
		if name == "/" {
			c.String(http.StatusOK, "Hello Bro !")
			return
		}
		name = strings.Trim(name, "/")
		c.String(http.StatusOK, "Hello " + name + " !")

	})
	gp.POST("/body", CoreMiddleWare(), func(c *gin.Context) {
		middleStr, exsist := c.Get("request")
		if exsist {
			log.Printf("get message from middle ware : %s", middleStr)
		}

		acct := c.Query("acct")
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
			"acct":    acct,
		})
	})
}

func CoreRouterGroup(gp *gin.RouterGroup) {
	gp.POST("/login", func(c *gin.Context){
		middleStr, exsist := c.Get("request")
		if exsist {
			log.Printf("get message from middle ware : %s", middleStr)
		}

		var user module.User

		contentType := c.Request.Header.Get("content-type")
		log.Printf("path:/login; content-type:%s", contentType)

		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}


		person := &module.Person{}
		err = person.GetPerson(user.Username, user.Password)

		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"password": "****",
			"first_name": person.FirstName,
			"last_name": person.LastName,
		})
	})
	gp.GET("/toBaidu", func(c *gin.Context) {
		middleStr, exsist := c.Get("request")
		if exsist {
			log.Printf("get message from middle ware : %s", middleStr)
		}

		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
}

