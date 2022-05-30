package main

import (
	_ "github.com/beego/beego/v2/server/web/session/mysql" // for Beego Session MySQL

	"github.com/yangsf5/auto3mad/backend/models/db/base"
	"github.com/yangsf5/auto3mad/backend/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	Init()

	web.Run()
}

func Init() {
	base.Init()
	routers.Init()

	web.BConfig.EnableGzip = true
	web.SetStaticPath("/", "../frontend/dist")
}
