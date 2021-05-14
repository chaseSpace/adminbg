package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"adminbg/util"
	"github.com/pkg/errors"
)

func NewUserGroupLogic(req *cproto.NewUserGroupReq) (*cproto.NewUserGroupRsp, error) {
	err := crud.InsertUserGroup(&req.Group)
	return new(cproto.NewUserGroupRsp), err
}

func UpdateUserGroupLogic(req *cproto.UpdateUserGroupReq) (*cproto.UpdateUserGroupRsp, error) {
	if req.GroupId < 1 {
		return nil, errors.Wrap(cerror.ErrParams, "invalid group_id")
	}
	if req.Delete {
		return new(cproto.UpdateUserGroupRsp), crud.DelUserGroup(req.GroupId)
	}
	err := crud.UpdateUserGroup(&req.Group)
	return new(cproto.UpdateUserGroupRsp), err
}

func QueryUserGroupLogic(req *cproto.QueryUserGroupReq) (*cproto.QueryUserGroupRsp, error) {
	entity, err := crud.QueryUserGroup(req.GroupId)
	rsp := &cproto.QueryUserGroupRsp{Group: entity}
	return rsp, err
}

func GetUserGroupListLogic(req *cproto.GetUserGroupListReq) (*cproto.GetUserGroupListRsp, error) {
	if err := util.CheckSplitPageParams(req.PageNum, req.PageSize); err != nil {
		return nil, err
	}
	list, total, err := crud.GetUserGroupList(req.PageNum, req.PageSize, crud.CreatedAtAsc, crud.UpdatedAtAsc)
	if err != nil {
		return nil, err
	}
	rsp := &cproto.GetUserGroupListRsp{Total: total}
	for _, item := range list {
		rsp.List = append(rsp.List, item.Proto())
	}
	return rsp, nil
}
