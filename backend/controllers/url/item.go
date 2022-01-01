package url

import (
	"backend/controllers/base"
	"backend/models/db"
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
