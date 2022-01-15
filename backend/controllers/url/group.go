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
	mg       url.GroupModel
}

func (c *GroupController) Prepare() {
	c.mg = *url.NewGroupModel()
	c.urlModel = url.ModelURL{}
	c.BaseController.Prepare()
}

func (c *GroupController) Get() {
	var rawGroups []url.Group
	err := c.mg.GetAll(&rawGroups)
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

	rg := &url.Group{
		ID:   info.ID,
		Desc: info.Title,
	}
	err = c.mg.Upsert(rg)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

func (c *GroupController) Delete() {
	id, err := c.GetInt("id")
	c.JSONErrorAbort(err)

	err = c.mg.DeleteByID(id)
	c.JSONErrorAbort(err)

	c.JSONOK()
}
