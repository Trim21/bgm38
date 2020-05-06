package model

type UserGroup int

const (
	UserGroup_Admin        UserGroup = 1
	UserGroup_BangumiAdmin UserGroup = 2
	UserGroup_DouJinAdmin  UserGroup = 3
	UserGroup_Banned       UserGroup = 4
	UserGroup_AccessDeny   UserGroup = 5
	UserGroup_RoleAdmin    UserGroup = 8
	UserGroup_WikiAdmin    UserGroup = 9
	UserGroup_User         UserGroup = 10
	UserGroup_Wiki         UserGroup = 11
)

func (u UserGroup) String() string {
	switch u {
	case UserGroup_Admin:
		return "管理员"
	case UserGroup_BangumiAdmin:
		return "Bangumi 管理猿"
	case UserGroup_DouJinAdmin:
		return "天窗管理猿"
	case UserGroup_Banned:
		return "禁言用户"
	case UserGroup_AccessDeny:
		return "禁止访问用户"
	case UserGroup_RoleAdmin:
		return "人物管理猿"
	case UserGroup_WikiAdmin:
		return "维基条目管理猿"
	case UserGroup_User:
		return "用户"
	case UserGroup_Wiki:
		return "维基人"
	}
	return "不明用户组"
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	UserGroup UserGroup `json:"usergroup"`
	Nickname  string    `json:"nickname"`
}
