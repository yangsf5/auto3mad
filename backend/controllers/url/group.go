package url

import (
	"backend/controllers/base"
	"backend/models/db"
	"encoding/json"
	"fmt"
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

func (c *GroupController) Post() {
	info := GroupInfo{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &info); err == nil {
		fmt.Printf("%v", info)
	} else {
		panic(err)
	}
}
