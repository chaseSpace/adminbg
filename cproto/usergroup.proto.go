package cproto

type NewUserGroupReq struct {
	Group
}

type NewUserGroupRsp struct {
}

type Group struct {
	GroupId   int16  `json:"group_id"` // ignored at create user-group action
	GroupName string `json:"group_name" binding:"required"`
	RoleId    int16  `json:"role_id"`
	CreatedAt string `json:"created_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
	UpdatedAt string `json:"updated_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
}

type UpdateUserGroupReq struct {
	Delete bool `json:"delete"` // if true, server will delete usergroup for this group_id, other params would be ignored.
	Group
}

type UpdateUserGroupRsp struct {
}

type QueryUserGroupReq struct {
	GroupId int32 `form:"group_id" binding:"required"`
}

type QueryUserGroupRsp struct {
	*Group // *XXX means this field is pointer type, it might be null.
}

type GetUserGroupListReq struct {
	PageNum  uint16 `form:"pn"`
	PageSize uint16 `form:"ps"`
}

type GetUserGroupListRsp struct {
	List  []*Group `json:"list"` // Order by CreatedAt by default.
	Total int64    `json:"total"`
}

type DeleteUserGroupReq struct {
}
