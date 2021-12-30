package url

import (
	"backend/controllers/base"
	"backend/models/db"
	"encoding/json"
)

type GroupInfo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type GroupController struct {
	base.BaseController
	urlModel db.ModelURL
}

func (c *GroupController) Prepare() {
	c.urlModel = db.ModelURL{}
	c.BaseController.Prepare()
}

func (c *GroupController) Get() {
	rawGroups, err := c.urlModel.GetAllGroups()
	if err != nil {
		c.JSONErrorAbort(err)
	}

	rets := []GroupInfo{}

	for _, rg := range rawGroups {
		gi := GroupInfo{
			ID:    rg.ID,
			Title: rg.Desc,
		}
		rets = append(rets, gi)
	}

	c.JSONOK(rets)
}

func (c *GroupController) Post() {
	info := GroupInfo{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &info); err == nil {
		rg := db.URLGroup{
			ID:   info.ID,
			Desc: info.Title,
		}
		err := c.urlModel.UpsertGroup(rg)
		if err != nil {
			c.JSONErrorAbort(err)
		}
	} else {
		c.JSONErrorAbort(err)
	}

	c.JSONOK()
}

func (c *GroupController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.JSONErrorAbort(err)
	}

	err = c.urlModel.DeleteGroup(id)
	if err != nil {
		c.JSONErrorAbort(err)
	}

	c.JSONOK()
}
