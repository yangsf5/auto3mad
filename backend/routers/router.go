package routers

import (
	"github.com/yangsf5/auto3mad/backend/controllers/daily"
	"github.com/yangsf5/auto3mad/backend/controllers/day"
	"github.com/yangsf5/auto3mad/backend/controllers/url"

	"github.com/beego/beego/v2/server/web"
)

func Init() {
	web.Router("/v2/url/apis", &url.APIController{})

	web.Router("/v2/day/memorials", &day.MemorialController{})
	web.Router("/v2/day/timestamp", &day.TimestampController{})

	web.Router("/v2/url/urls", &url.URLController{})
	web.Router("/v2/url/groups", &url.GroupController{})
	web.Router("/v2/url/items", &url.ItemController{})

	web.Router("/v2/daily/routines", &daily.RoutineController{})
	web.Router("/v2/daily/events", &daily.EventController{})
	web.Router("/v2/daily/stats", &daily.StatController{})
}
