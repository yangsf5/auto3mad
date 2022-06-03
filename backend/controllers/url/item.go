package url

import (
	"encoding/json"

	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/db/url"
)

type ItemInfo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	URL     string `json:"url"`
	GroupID int    `json:"group_id"`
}

type ItemController struct {
	base.Controller
	m url.ItemModel
}

func (c *ItemController) Prepare() {
	c.Controller.Prepare()

	c.m = *url.NewItemModel(c.GetMyUserID())
}

func (c *ItemController) Get() {
	var rawItems []url.Item
	err := c.m.GetAllOrderBy("group_id", &rawItems)
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
		UserID:  c.GetMyUserID(),
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

	err = c.m.Delete(id)
	c.JSONErrorAbort(err)

	c.JSONOK()
}
