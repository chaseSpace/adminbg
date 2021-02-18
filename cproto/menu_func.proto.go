package cproto

import (
	"fmt"
)

type OneMenu struct {
	ParentId    int32  `json:"parent_id" binding:"required"`    // Set to 100 if is first class menu
	Level       int8   `json:"level" binding:"required"`        // Set to 1 or 2...
	Name        string `json:"name" binding:"required"`         // Name of menu or function
	Route       string `json:"route" binding:"required"`        // It was used at frontend eventually
	MenuDisplay string `json:"menu_display" binding:"required"` // Take Y|N, means display or not for menu. It was used at frontend eventually
	Id          int32  `json:"id"`                              // Menu id, that does **not need** to appear at /NewMenu request
	SortNum     uint16 `json:"sort_num"`
	/*
		NOTE: The following fields do not need to appear in the request.
	*/
	CreatedAt  int64      `json:"created_at"`
	Children   []*OneMenu `json:"children"`
	ChildFuncs []*OneFunc `json:"child_funcs"` // Only leaf-menu have child funcs
}

type OneFunc struct {
	Id      int32  `json:"id"`                         // Func id, that does not need to appear at /NewFunction request
	MenuId  int32  `json:"menu_id" binding:"required"` // Which menu to belong
	Name    string `json:"name" binding:"required"`
	SortNum uint16 `json:"sort_num"`
	/*
		NOTE: The following fields do not need to appear in the request.
	*/
	CreatedAt int64 `json:"created_at"`
}

// POST /web/v1/NewMenu
type NewMenuReq struct {
	OneMenu
}
type NewMenuRsp struct{}

// POST /web/v1/UpdateMenu
type UpdateMenuReq struct {
	OneMenu
}
type UpdateMenuRsp struct{}

// GET /web/v1/GetMenuList
// NOTE: This API only was used administrator.
type GetMenuListReq struct{}

type GetMenuListRsp struct {
	List []*OneMenu `json:"list"`
}

type MfType string

const (
	Menu     MfType = "MENU"
	Function MfType = "FUNCTION"
)

func (t MfType) Check() error {
	switch t {
	case Menu, Function:
		return nil
	}
	return fmt.Errorf("invalid MfType value:%s", t)
}

type MenuDisplay string

const (
	Display    MenuDisplay = "Y"
	NotDisplay MenuDisplay = "N"
)

func (t MenuDisplay) Check() error {
	switch t {
	case Display, NotDisplay:
		return nil
	}
	return fmt.Errorf("invalid MenuDisplay value:%s", t)
}

// DELETE /web/v1/DeleteMenus
type DeleteMenusReq struct {
	MenuIds []int32 `json:"menu_ids" binding:"required"`
}

type DeleteMenusRsp struct{}

// POST /web/v1/NewFunction
type NewFunctionReq struct {
	OneFunc
}

type NewFunctionRsp struct{}

// POST /web/v1/UpdateFunction
type UpdateFunctionReq struct {
	OneFunc
}

type UpdateFunctionRsp struct{}

// GET /web/v1//GetAPIList
type GetAPIListReq struct {
	// all below are optional params.
	BindFunctionId       int32  `form:"bind_function_id"`        // if not provided(default is zero), server will return all APIs' info
	FuzzySearchByAPIName string `form:"fuzzy_search_by_name"`    // Fuzzy search APIs by name
	SortByCreatedAtDesc  bool   `form:"sort_by_created_at_desc"` // Sort params, `Desc` means sort in descending order, default is sort by `created_at` field in ascending order
}

type GetAPIListRsp struct {
	List []*OneAPI `json:"list"`
}

type OneAPI struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
