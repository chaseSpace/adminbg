package handler

import (
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"adminbg/pkg/model"
)

func NewMenuLogic(req *cproto.NewMenuReq) (*cproto.NewMenuRsp, error) {
	row := model.MenuAndFunction{
		MfName: req.Name,
		//Path:        "",  insert by crud layer
		//Type:        "",  set by crud layer
		ParentId:    req.ParentId,
		Level:       req.Level,
		MenuRoute:   req.Route,
		MenuDisplay: cproto.MenuDisplay(req.MenuDisplay),
		SortNum:     req.SortNum,
	}
	err := crud.InsertNewMenu(&row)
	return new(cproto.NewMenuRsp), err
}
