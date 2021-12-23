package routers

import (
	"backend/controllers/day"
	"backend/controllers/url"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/v2/day/countdowns", &day.CountdownController{})
	web.Router("/v2/url/urls", &url.URLController{})
	web.Router("/v2/url/apis", &url.APIController{})
}
