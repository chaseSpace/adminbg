package cproto

// POST /web/v1/NewRole
type NewRoleReq struct {
	Role
}

type NewRoleRsp struct {
}

type Role struct {
	RoleId    int16  `json:"role_id"` // only be used at query/update action
	RoleName  string `json:"role_name" binding:"required"`
	CreatedAt string `json:"created_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
	UpdatedAt string `json:"updated_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
}

// POST /web/v1/UpdateRole
type UpdateRoleReq struct {
	Delete bool `json:"delete"`
	Role
}

type UpdateRoleRsp struct {
}

// GET /web/v1/QueryRole
type QueryRoleReq struct {
	RoleId int16 `form:"role_id" binding:"required"`
}

type QueryRoleRsp struct {
	*Role
}

// GET /web/v1/GetRoleList
type GetRoleListReq struct {
	/* No split page */
}

type GetRoleListRsp struct {
	List []*Role `json:"list"`
}
