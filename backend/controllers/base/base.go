package base

import (
	"github.com/yangsf5/auto3mad/backend/util/enricherror"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type Controller struct {
	web.Controller
}

func (c *Controller) Prepare() {
	c.Controller.Prepare()
	c.auth()
}

func (c *Controller) auth() {
	if mode, _ := web.AppConfig.String("runmode"); mode == "dev" {
		_ = c.InitMyUserInfo(1, "3mad", "三疯")

		return
	}

	curPath := c.Ctx.Request.URL.Path

	notNeedAuthPaths := map[string]bool{
		"/v2/auth/github": true, // Login URL
	}

	if _, ok := notNeedAuthPaths[curPath]; ok {
		return
	}

	c.JSONErrorAbort("Please Login First!", c.IsLogined())
}

// response: Client AntDesignPro Request Struct
type response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"` // nolint
}

// Eg1: JSONErrorAbort(error)
// Eg2: JSONErrorAbort("something wrong")
// Eg3: JSONErrorAbort("something wrong", bool)
func (c *Controller) JSONErrorAbort(err interface{}, errCondition ...bool) {
	if err == nil {
		return
	}

	if len(errCondition) == 1 && !errCondition[0] {
		return
	}

	content := enricherror.GetErrorContent(err)
	if content == "" {
		return
	}

	emsg := enricherror.ErrorPosition() + content
	logs.Error(emsg)

	c.Data["json"] = response{
		Success:      false,
		ErrorMessage: emsg,
	}
	_ = c.ServeJSON()
	c.StopRun()
}

func (c *Controller) JSONOKAbort(v ...interface{}) {
	c.JSONOK(v...)
	c.StopRun()
}

func (c *Controller) JSONOK(v ...interface{}) {
	res := response{
		Success: true,
	}
	if len(v) > 0 {
		res.Data = v[0]
	}

	c.Data["json"] = res
	_ = c.ServeJSON()
}
