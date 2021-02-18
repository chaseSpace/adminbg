package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"adminbg/pkg/model"
	"adminbg/util"
	"github.com/pkg/errors"
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

func UpdateMenuLogic(req *cproto.UpdateMenuReq) (*cproto.UpdateMenuRsp, error) {
	if req.Id == 0 {
		return nil, errors.Wrap(cerror.ErrParams, "need an valid id")
	}
	row := model.MenuAndFunction{
		MfId:        req.Id,
		MfName:      req.Name,
		ParentId:    req.ParentId,
		Level:       req.Level,
		MenuRoute:   req.Route,
		MenuDisplay: cproto.MenuDisplay(req.MenuDisplay),
		SortNum:     req.SortNum,
	}
	err := crud.UpdateMenu(&row)
	return new(cproto.UpdateMenuRsp), err
}

func GetMenuListLogic(_ *cproto.GetMenuListReq) (*cproto.GetMenuListRsp, error) {
	// Get all menus and functions, then restructure them on memory.
	menuFuncList, err := crud.GetMenuFuncList()
	if err != nil {
		return nil, err
	}
	var firstLevelMenus []*cproto.OneMenu
	menuMap := make(map[int32][]*cproto.OneMenu) // key=>parent menu-id of menu
	funcMap := make(map[int32][]*cproto.OneFunc) // key=>parent menu-id of function

	for _, mf := range menuFuncList {
		if model.IsReservedMenuId(mf.MfId) {
			continue
		}
		switch mf.Type {
		case cproto.Menu:
			menu := &cproto.OneMenu{
				ParentId:    mf.ParentId,
				Level:       mf.Level,
				Name:        mf.MfName,
				Route:       mf.MenuRoute,
				MenuDisplay: string(mf.MenuDisplay),
				Id:          mf.MfId,
				SortNum:     mf.SortNum,
				CreatedAt:   mf.CreatedAt.Unix(),
				//Child:       nil,
				//ChildFunc:   nil,
			}
			if mf.Level == 1 {
				firstLevelMenus = append(firstLevelMenus, menu)
			}
			if model.IsReservedMenuId(mf.ParentId) {
				continue
			}
			if menuMap[mf.ParentId] != nil {
				menuMap[mf.ParentId] = append(menuMap[mf.ParentId], menu)
			} else {
				menuMap[mf.ParentId] = []*cproto.OneMenu{menu}
			}
		case cproto.Function:
			Func := &cproto.OneFunc{
				Id:        mf.MfId,
				MenuId:    mf.ParentId,
				Name:      mf.MfName,
				SortNum:   mf.SortNum,
				CreatedAt: mf.CreatedAt.Unix(),
			}
			if funcMap[mf.ParentId] != nil {
				funcMap[mf.ParentId] = append(funcMap[mf.ParentId], Func)
			} else {
				funcMap[mf.ParentId] = []*cproto.OneFunc{Func}
			}
		}
	}

	for _, menu := range firstLevelMenus {
		setChildWithRecursive(menu, menuMap, funcMap)
	}
	rsp := cproto.GetMenuListRsp{
		List: firstLevelMenus,
	}
	return &rsp, nil
}

func setChildWithRecursive(menu *cproto.OneMenu, menuMap map[int32][]*cproto.OneMenu, funcMap map[int32][]*cproto.OneFunc) {
	menu.Children = menuMap[menu.Id]
	for _, menu := range menu.Children {
		setChildWithRecursive(menu, menuMap, funcMap)
	}
	menu.ChildFuncs = funcMap[menu.Id]
}

func DeleteMenusLogic(req *cproto.DeleteMenusReq) (*cproto.DeleteMenusRsp, error) {
	err := crud.DeleteMenus(req.MenuIds)
	return new(cproto.DeleteMenusRsp), err
}

func NewFunctionLogic(req *cproto.NewFunctionReq) (*cproto.NewFunctionRsp, error) {
	row := model.MenuAndFunction{
		MfName: req.Name,
		//Path:        "",  insert by crud layer
		//Type:        "",  set by crud layer
		//Level:    0,
		//MenuRoute:   "",
		MenuDisplay: cproto.NotDisplay, // For check pass
		ParentId:    req.MenuId,
		SortNum:     req.SortNum,
	}
	err := crud.InsertNewFunc(&row)
	return new(cproto.NewFunctionRsp), err
}

func UpdateFunctionLogic(req *cproto.UpdateFunctionReq) (*cproto.UpdateFunctionRsp, error) {
	if req.Id == 0 {
		return nil, errors.Wrap(cerror.ErrParams, "need an valid id")
	}
	row := model.MenuAndFunction{
		MfId:        req.Id,
		MfName:      req.Name,
		ParentId:    req.MenuId,
		SortNum:     req.SortNum,
		MenuDisplay: cproto.Display,
	}
	err := crud.UpdateFunction(&row)
	return new(cproto.UpdateFunctionRsp), err
}

func GetAPIListLogic(req *cproto.GetAPIListReq) (*cproto.GetAPIListRsp, error) {
	apis, err := crud.GetAPIList(req.BindFunctionId, req.FuzzySearchByAPIName, req.SortByCreatedAtDesc)
	if err != nil {
		return nil, err
	}
	rsp := new(cproto.GetAPIListRsp)
	for _, api := range apis {
		rsp.List = append(rsp.List, &cproto.OneAPI{
			Id:        api.ApiId,
			Name:      api.Name,
			CreatedAt: api.CreatedAt.Format(util.TimeLayout),
		})
	}
	return rsp, nil
}
