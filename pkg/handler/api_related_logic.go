package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"github.com/pkg/errors"
)

func NewAPILogic(req *cproto.NewAPIReq) (*cproto.NewAPIRsp, error) {
	err := crud.NewAPI(req.Identity, req.Remark)
	return new(cproto.NewAPIRsp), err
}

func UpdateAPILogic(req *cproto.UpdateAPIReq) (*cproto.UpdateAPIRsp, error) {
	rsp := new(cproto.UpdateAPIRsp)
	if req.Delete {
		err := crud.DeleteAPIs(req.DeleteApiIds...)
		return rsp, err
	}
	// Update
	if req.Identity == "" {
		return nil, errors.Wrap(cerror.ErrParams, "invalid params `identity`")
	}
	err := crud.UpdateAPI(req.ApiId, req.Identity, req.Remark)
	return rsp, err
}
