package url

import (
	"github.com/yangsf5/auto3mad/backend/controllers/base"

	"github.com/beego/beego/v2/server/web"
)

type APIController struct {
	base.Controller
}

type BackendAPI struct {
	RouterPattern string `json:"router_pattern"`
	Controller    string `json:"controller"`
}

func (c *APIController) Get() {
	rets := []BackendAPI{}

	ts := web.BeeApp.PrintTree()
	ms := ts["Data"].(web.M)
	gets := ms["GET"].(*[][]string)

	for _, get := range *gets {
		api := BackendAPI{
			RouterPattern: get[0],
			Controller:    get[2],
		}
		rets = append(rets, api)
	}

	c.JSONOK(rets)
}
