package auth

import (
	"github.com/yangsf5/auto3mad/backend/controllers/base"
)

type LoginController struct {
	base.Controller
}

func (c *LoginController) Post() {
	userName := c.GetString("user_name")

	// TODO Check DB
	userID := 1
	nickName := ""

	_ = c.InitMyUserInfo(userID, userName, nickName)
}
