package url

import (
	"backend/controllers/base"
	"backend/models/db"
)

type MiscController struct {
	base.BaseController
}

func (c *MiscController) Get() {
	urlModel := db.ModelURL{}

	max, err := urlModel.GetMaxGroupID()
	if err != nil {
		c.JSONErrorAbort(err)
	}

	c.JSONOK(max)
}
