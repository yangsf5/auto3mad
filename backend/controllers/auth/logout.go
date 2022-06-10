package auth

import "github.com/yangsf5/auto3mad/backend/controllers/base"

type LogoutController struct {
	base.Controller
}

func (c *LogoutController) Get() {
	_ = c.InitMyUserInfo(0, "", "")
	c.JSONOK()
}
