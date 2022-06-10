package auth

import (
	"github.com/yangsf5/auto3mad/backend/models/db/base"
)

type UserGitHub struct {
	ID              int    `orm:"pk;column(id)"`
	UserID          int    `orm:"column(user_id)"`
	GitHubID        int    `orm:"column(github_id)"`
	GitHubLoginName string `orm:"column(github_login_name)"`
	GitHubNickName  string `orm:"column(github_nick_name)"`
	GitHubEmail     string `orm:"column(github_email)"`
	GitHubCreatedAt string `orm:"column(github_created_at)"`
	LastLogin       int
}

func (o *UserGitHub) TableName() string {
	return "user_github"
}

func (o *UserGitHub) GetID() int {
	return o.ID
}

func (o *UserGitHub) NewObject() interface{} {
	return new(UserGitHub)
}

type UserGitHubModel struct {
	base.Model
}

func NewUserGitHubModel() *UserGitHubModel {
	m := new(UserGitHubModel)
	m.Model = *base.NewModelSTD(&UserGitHub{})

	return m
}
