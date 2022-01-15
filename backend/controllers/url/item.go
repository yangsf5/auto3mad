package url

import (
	"backend/controllers/base"
	"backend/models/db/url"
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
	m url.ItemModel
}

func (c *ItemController) Prepare() {
	c.m = *url.NewItemModel()
	c.BaseController.Prepare()
}

func (c *ItemController) Get() {
	var rawItems []url.Item
	err := c.m.GetAll(&rawItems)
	c.JSONErrorAbort(err)

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
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	c.JSONErrorAbort(err)

	ri := &url.Item{
		ID:      info.ID,
		Title:   info.Title,
		Icon:    info.Icon,
		URL:     info.URL,
		GroupID: info.GroupID,
	}
	err = c.m.Upsert(ri)
	c.JSONErrorAbort(err)

	c.JSONOK()
}

func (c *ItemController) Delete() {
	id, err := c.GetInt("id")
	c.JSONErrorAbort(err)

	err = c.m.DeleteByID(id)
	c.JSONErrorAbort(err)

	c.JSONOK()
}
