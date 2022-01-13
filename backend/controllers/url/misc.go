package url

import (
	"backend/controllers/base"
	"backend/models/db"
)

type MiscController struct {
	base.BaseController
}

func (c *MiscController) Get() {
	k := c.GetString("kind")

	urlModel := db.ModelURL{}
	max, err := urlModel.GetMaxID(k)
	c.JSONErrorAbort(err)

	c.JSONOK(max)
}
