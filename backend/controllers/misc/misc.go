package misc

import (
	"backend/controllers/base"
	"backend/models/db"
)

type MiscController struct {
	base.BaseController
}

func (c *MiscController) Get() {
	k := c.GetString("kind")

	max, err := db.GetMaxID(k)
	c.JSONErrorAbort(err)

	c.JSONOK(max)
}
