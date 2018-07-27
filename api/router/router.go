package router

import (
	"github.com/go-martini/martini"
	"motimer/api/user"
)

func Router(r martini.Router) {
	r.Group("/user", UserRouter)
}

func UserRouter(r martini.Router) {
	u := user.UserController{}
	r.Get("/userInfo", u.GetUserInfo)
}
