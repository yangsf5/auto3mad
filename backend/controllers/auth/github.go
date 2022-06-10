package auth

import (
	"crypto/tls"
	"time"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/server/web"
	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/db/auth"
)

type GitHubController struct {
	base.Controller
}

type accessTokenGitHub struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type userInfoGitHub struct {
	ID        int    `json:"id"`
	LoginName string `json:"login"`
	NickName  string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (c *GitHubController) Get() {
	// GitHub Doc: https://docs.github.com/cn/developers/apps/building-oauth-apps/authorizing-oauth-apps
	// 1. 前端将用户重定向访问 https://github.com/login/oauth/authorize
	// 2. GitHub 得到用户授权后，再将用户重定向访问本接口，并带上参数 Code
	code := c.GetString("code")

	// 3. 本接口拿 Code，直接调用 GitHub API，获取 access_token
	req := httplib.Post("https://github.com/login/oauth/access_token")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // nolint
	req.SetTimeout(10*time.Second, 5*time.Second)                 // nolint

	clientID, err := web.AppConfig.String("GitHubClientID")
	c.JSONErrorAbort(err)
	clientSecret, err := web.AppConfig.String("GitHubClientSecret")
	c.JSONErrorAbort(err)

	req.Param("client_id", clientID)
	req.Param("client_secret", clientSecret)
	req.Param("code", code)
	req.Header("Accept", "application/json")

	var retToken accessTokenGitHub
	err = req.ToJSON(&retToken)
	c.JSONErrorAbort(err)

	// 4. 本接口，再拿 access_token 去调用 GitHub API，获取用户授权的资源（如用户 email 等授权的信息）
	req = httplib.Get("https://api.github.com/user")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // nolint
	req.SetTimeout(10*time.Second, 5*time.Second)                 // nolint
	req.SetBasicAuth("token", retToken.AccessToken)

	var retUser userInfoGitHub
	err = req.ToJSON(&retUser)
	c.JSONErrorAbort(err)

	model := auth.NewUserGitHubModel()

	ug := auth.UserGitHub{
		GitHubID:        retUser.ID,
		GitHubLoginName: retUser.LoginName,
		GitHubNickName:  retUser.NickName,
		GitHubEmail:     retUser.Email,
		GitHubCreatedAt: retUser.CreatedAt,
	}
	_, id, err := model.ReadOrCreate(&ug, "github_email")
	c.JSONErrorAbort(err)

	err = c.InitMyUserInfo(int(id), retUser.Email, retUser.NickName)
	c.JSONErrorAbort(err)

	c.Redirect("http://127.0.0.1:11270/#/", 302)
}
