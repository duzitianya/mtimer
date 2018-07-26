package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"os"
	"git.mobike.io/motimer/api/router"
)

func main() {
	m := martini.Classic();
	m.Map(log.New(os.Stdout, "[MOTIMER] INFO ", log.Ldate|log.Ltime))
	m.Use(render.Renderer())
	m.Get("/", func() string {
		return "Hello"
	})
	m.Group("/api", router.Router)
	m.RunOnAddr("1278")

}
