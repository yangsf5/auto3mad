package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/yangsf5/auto3mad/backend/models/db/base"
	"github.com/yangsf5/auto3mad/backend/routers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	Init()

	e := fmt.Sprintf("display notification \"Restarted.\" with title \"%s\"", web.BeeApp.Cfg.AppName)

	cmd := exec.Command("osascript", "-e", e)
	var eout bytes.Buffer
	cmd.Stderr = &eout
	if err := cmd.Run(); err != nil {
		logs.Error("Exec stderr", err)
	}
	logs.Debug("Exec stdout: ", eout.String())

	web.Run()
}

func Init() {
	base.Init()
	routers.Init()
}
