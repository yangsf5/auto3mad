package misc

import (
	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/db"
)

type MiscController struct {
	base.Controller
}

func (c *MiscController) Get() {
	k := c.GetString("kind")

	max, err := db.GetMaxID(k)
	c.JSONErrorAbort(err)

	c.JSONOK(max)
}
