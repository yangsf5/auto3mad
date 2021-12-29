package base

import (
	beego "github.com/beego/beego/v2/server/web"

	"backend/util/enricherror"
)

type BaseController struct {
	beego.Controller
}

// Eg1: JSONError(error)
// Eg2: JSONError("something wrong")
// Eg3: JSONError("something wrong", bool)
func (c *BaseController) JSONError(err interface{}, errCondition ...bool) {
	if len(errCondition) == 1 && !errCondition[0] {
		return
	}
	content := enricherror.GetErrorContent(err)
	if content == "" {
		return
	}
	c.responseJSON(false, nil, enricherror.ErrorPosition()+content)
}

func (c *BaseController) JSONOK(v ...interface{}) {
	if len(v) == 0 {
		c.responseJSON(true, nil, nil)
	} else {
		c.responseJSON(true, v[0], nil)
	}
}

func (c *BaseController) responseJSON(success bool, data interface{}, err interface{}) {
	c.Data["json"] = map[string]interface{}{
		"success":      success,
		"data":         data,
		"errorMessage": err,
	}
	c.ServeJSON()
	c.StopRun()
}
