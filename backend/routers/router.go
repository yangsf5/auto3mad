package routers

import (
	"backend/controllers/day"
	"backend/controllers/url"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/v2/day/countdowns", &day.CountdownController{})
	beego.Router("/v2/url/urls", &url.URLController{})
}
