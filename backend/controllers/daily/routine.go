package daily

import (
	"backend/controllers/base"
	"backend/models/db/daily"
)

type RoutineController struct {
	base.BaseController
	m daily.RoutineModel
}

func (c *RoutineController) Prepare() {
	c.m = *daily.NewRoutineModel()
	c.BaseController.Prepare()
}

func (c *RoutineController) Get() {
	var rs []daily.Routine
	err := c.m.GetAll(&rs)
	c.JSONErrorAbort(err)

	c.JSONOK(rs)
}
