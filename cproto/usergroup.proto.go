package cproto

type NewUserGroupReq struct {
	Group
}

type NewUserGroupRsp struct {
}

type Group struct {
	GroupId   int16  `json:"group_id"` // only be used at query action
	GroupName string `json:"group_name" binding:"required"`
	RoleId    int16  `json:"role_id"`
	CreatedAt string `json:"created_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
	UpdatedAt string `json:"updated_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
}
