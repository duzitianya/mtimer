package module

type User struct {
	//其中json标签用来绑定content-type=application/json时的参数名称
	//form标签用来绑定x-www-form-urlencoded时的参数名称
	Username string `json:"uname" form:"un" binding:"required"`
	Password string `json:"pwds" form:"pwd" binding:"required"`
	Age int `json:"age" form:"age"`
}
