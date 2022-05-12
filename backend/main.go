package main

import (
	"bytes"
	"fmt"
	"os/exec"

	_ "backend/routers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

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
