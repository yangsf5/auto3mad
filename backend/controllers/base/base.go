package base

import (
	beego "github.com/beego/beego/v2/server/web"

	"backend/util/enricherror"
)

type BaseController struct {
	beego.Controller
}

// Eg1: JSONDieIfError(error)
// Eg2: JSONDieIfError("something wrong")
// Eg3: JSONDieIfError("something wrong", bool)
func (c *BaseController) JSONDieIfError(err interface{}, errCondition ...bool) {
	if len(errCondition) == 1 && !errCondition[0] {
		return
	}
	content := enricherror.GetErrorContent(err)
	if content == "" {
		return
	}
	c.Data["json"] = map[string]string{"ret": enricherror.ErrorPosition() + content}
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) JSONOK(v ...interface{}) {
	if len(v) == 0 {
		c.Data["json"] = map[string]string{"ret": "ok"}
	} else {
		c.Data["json"] = v[0]
	}
	c.ServeJSON()
	c.StopRun()
}
