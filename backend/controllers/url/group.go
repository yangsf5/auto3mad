package url

import (
	"backend/controllers/base"
	"backend/models/db"
)

type GroupInfo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type GroupController struct {
	base.BaseController
}

func (c *GroupController) Get() {
	urlModel := db.ModelURL{}

	rawGroups, err := urlModel.GetAllGroups()
	if err != nil {
		c.JSONDieIfError(err)
	}

	rets := []GroupInfo{}

	for _, rg := range rawGroups {
		gi := GroupInfo{
			ID:    rg.ID,
			Title: rg.Desc,
		}
		rets = append(rets, gi)
	}

	c.JSONOK(map[string]interface{}{
		"data": rets,
	})
}
