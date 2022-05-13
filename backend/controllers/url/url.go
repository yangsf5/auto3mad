package url

import (
	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/db/url"
)

type URLController struct {
	base.BaseController
}

type ItemURL struct {
	Icon  string `json:"icon"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type GroupURL struct {
	Title string    `json:"title"`
	URLs  []ItemURL `json:"urls"`
}

func (c *URLController) Get() {
	mi := url.NewItemModel()
	mg := url.NewGroupModel()

	// 先为 ret 构建一个 map 形式的 Groups，方便组装数据
	retMap := map[int]GroupURL{}

	var rawGroups []url.Group
	err := mg.GetAll(&rawGroups)
	c.JSONErrorAbort(err)

	// 先根据 raw 数据来初始化 ret group map
	for _, rg := range rawGroups {
		group := GroupURL{
			Title: rg.Desc,
			URLs:  []ItemURL{},
		}
		retMap[rg.ID] = group
	}

	var items []url.Item
	err = mi.GetAll(&items)
	c.JSONErrorAbort(err)

	// 往 ret group map 里组装 urls
	for _, si := range items {
		ti := ItemURL{
			Icon:  si.Icon,
			URL:   si.URL,
			Title: si.Title,
		}
		g := retMap[si.GroupID]
		g.URLs = append(g.URLs, ti)
		retMap[si.GroupID] = g
	}

	rets := []GroupURL{}

	// 以原始 Groups 来遍历，能有排序
	for _, rg := range rawGroups {
		rets = append(rets, retMap[rg.ID])
	}

	c.JSONOK(rets)
}
