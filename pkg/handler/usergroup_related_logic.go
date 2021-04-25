package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"github.com/pkg/errors"
)

func NewUserGroupLogic(req *cproto.NewUserGroupReq) (*cproto.NewUserGroupRsp, error) {
	err := crud.InsertUserGroup(&req.Group)
	return new(cproto.NewUserGroupRsp), err
}

func UpdateUserGroupLogic(req *cproto.UpdateUserGroupReq) (*cproto.UpdateUserGroupRsp, error) {
	if req.GroupId <= 0 {
		return nil, errors.Wrap(cerror.ErrParams, "invalid group_id")
	}
	err := crud.UpdateUserGroup(&req.Group)
	return new(cproto.UpdateUserGroupRsp), err
}
