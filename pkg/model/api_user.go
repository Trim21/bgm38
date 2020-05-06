package model

type UserGroup int

const (
	UserGroupAdmin        UserGroup = 1
	UserGroupBangumiAdmin UserGroup = 2
	UserGroupDouJinAdmin  UserGroup = 3
	UserGroupBanned       UserGroup = 4
	UserGroupAccessDeny   UserGroup = 5
	UserGroupRoleAdmin    UserGroup = 8
	UserGroupWikiAdmin    UserGroup = 9
	UserGroupUser         UserGroup = 10
	UserGroupWiki         UserGroup = 11
)

func (u UserGroup) String() string {
	switch u {
	case UserGroupAdmin:
		return "管理员"
	case UserGroupBangumiAdmin:
		return "Bangumi 管理猿"
	case UserGroupDouJinAdmin:
		return "天窗管理猿"
	case UserGroupBanned:
		return "禁言用户"
	case UserGroupAccessDeny:
		return "禁止访问用户"
	case UserGroupRoleAdmin:
		return "人物管理猿"
	case UserGroupWikiAdmin:
		return "维基条目管理猿"
	case UserGroupUser:
		return "用户"
	case UserGroupWiki:
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
