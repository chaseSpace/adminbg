package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"github.com/pkg/errors"
)

func NewUserGroupLogic(req *cproto.NewUserGroupReq) (*cproto.NewUserGroupRsp, error) {
	if req.GroupName == "" {
		return nil, errors.Wrap(cerror.ErrParams, "GroupName is required")
	}
	err := crud.InsertUserGroup(req.GroupName, req.RoleId)
	return new(cproto.NewUserGroupRsp), err
}
