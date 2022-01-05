package url

import (
	"backend/controllers/base"
	"backend/models/db"
	"encoding/json"
)

type ItemInfo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	URL     string `json:"url"`
	GroupID int    `json:"group_id"`
}

type ItemController struct {
	base.BaseController
	urlModel db.ModelURL
}

func (c *ItemController) Prepare() {
	c.urlModel = db.ModelURL{}
	c.BaseController.Prepare()
}

func (c *ItemController) Get() {
	rawItems, err := c.urlModel.GetAllItems()
	if err != nil {
		c.JSONErrorAbort(err)
	}

	rets := []ItemInfo{}

	for _, ri := range rawItems {
		i := ItemInfo{
			ID:      ri.ID,
			Title:   ri.Title,
			Icon:    ri.Icon,
			URL:     ri.URL,
			GroupID: ri.GroupID,
		}
		rets = append(rets, i)
	}

	c.JSONOK(rets)
}

func (c *ItemController) Post() {
	info := ItemInfo{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &info); err == nil {
		ri := db.URLItem{
			ID:      info.ID,
			Title:   info.Title,
			Icon:    info.Icon,
			URL:     info.URL,
			GroupID: info.GroupID,
		}
		err := c.urlModel.UpsertItem(ri)
		if err != nil {
			c.JSONErrorAbort(err)
		}
	} else {
		c.JSONErrorAbort(err)
	}

	c.JSONOK()
}

func (c *ItemController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.JSONErrorAbort(err)
	}

	err = c.urlModel.DeleteItem(id)
	if err != nil {
		c.JSONErrorAbort(err)
	}

	c.JSONOK()
}
