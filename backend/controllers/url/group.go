package url

import (
	"backend/controllers/base"
	"backend/models/db/url"
	"encoding/json"
)

type GroupInfo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type GroupController struct {
	base.BaseController
	urlModel url.ModelURL
}

func (c *GroupController) Prepare() {
	c.urlModel = url.ModelURL{}
	c.BaseController.Prepare()
}

func (c *GroupController) Get() {
	rawGroups, err := c.urlModel.GetAllGroups()
	c.JSONErrorAbort(err)

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

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	c.JSONErrorAbort(err)

	rg := url.URLGroup{
		ID:   info.ID,
		Desc: info.Title,
	}
	err = c.urlModel.UpsertGroup(rg)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

func (c *GroupController) Delete() {
	id, err := c.GetInt("id")
	c.JSONErrorAbort(err)

	err = c.urlModel.DeleteGroup(id)
	c.JSONErrorAbort(err)

	c.JSONOK()
}
