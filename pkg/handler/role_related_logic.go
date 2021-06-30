package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"adminbg/pkg/model"
)

func NewRoleLogic(req *cproto.NewRoleReq) (*cproto.NewRoleRsp, error) {
	err := crud.NewRole(&req.Role)
	return new(cproto.NewRoleRsp), err
}

func UpdateRoleLogic(req *cproto.UpdateRoleReq) (*cproto.UpdateRoleRsp, error) {
	if req.RoleId < 1 {
		return nil, cerror.ErrParams
	}
	if req.RoleId == model.DefaultRoleId {
		return nil, cerror.ErrCantOptReservedData
	}
	var err error
	if req.Delete {
		err = crud.DeleteRole(req.RoleId)
		return new(cproto.UpdateRoleRsp), err
	}
	// update
	err = crud.UpdateRole(&req.Role)
	return new(cproto.UpdateRoleRsp), err
}

func QueryRoleLogic(req *cproto.QueryRoleReq) (*cproto.QueryRoleRsp, error) {
	ent, err := crud.QueryRole(req.RoleId)
	if err != nil {
		return nil, err
	}
	rsp := &cproto.QueryRoleRsp{Role: ent.Proto()}
	return rsp, err
}

func GetRoleListLogic(_ *cproto.GetRoleListReq) (*cproto.GetRoleListRsp, error) {
	list, err := crud.GetRoleList(crud.CreatedAtAsc, crud.UpdatedAtAsc)
	if err != nil {
		return nil, err
	}
	rsp := &cproto.GetRoleListRsp{}
	for _, i := range list {
		rsp.List = append(rsp.List, i.Proto())
	}
	return rsp, err
}
