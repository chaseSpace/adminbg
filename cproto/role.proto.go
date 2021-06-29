package cproto

// POST /web/v1/NewRole
type NewRoleReq struct {
	Role
}

type NewRoleRsp struct {
}

type Role struct {
	RoleId   int16  `json:"role_id"` // only be used at query/update action
	RoleName string `json:"role_name" binding:"required"`
}

// POST /web/v1/UpdateRole
type UpdateRoleReq struct {
	Delete bool `json:"delete"`
	Role
}

type UpdateRoleRsp struct {
}
