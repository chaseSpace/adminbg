package cproto

import "fmt"

type OneMenu struct {
	ParentId    int32  `json:"parent_id" binding:"required"`    // Set to 100 if is first class menu
	Level       int8   `json:"level" binding:"required"`        // Set to 1 or 2...
	Name        string `json:"name" binding:"required"`         // Name of menu or function
	Route       string `json:"route" binding:"required"`        // It was used at frontend eventually
	MenuDisplay string `json:"menu_display" binding:"required"` // Take Y|N, means display or not for menu. It was used at frontend eventually
	SortNum     uint16 `json:"sort_num"`
	/*
		NOTE: The following fields are unnecessary at request.
	*/
	CreatedAt string   `json:"created_at"` // YYYY-mm-dd HH:MM:SS
	Child     *OneMenu `json:"child"`
	ChildFunc *OneFunc `json:"child_func"`
	Id        int32    `json:"id"`
}

type OneFunc struct {
	MenuId int32  `json:"menu_id" binding:"required"` // Which menu to belong
	Name   string `json:"name" binding:"required"`    // Function's name
	Id     int32  `json:"id"`                         // NOTE: It is not necessary at request
}

type NewMenuReq struct {
	OneMenu
}
type NewMenuRsp struct{}

type UpdateMenuReq struct {
	OneMenu
}
type UpdateMenuRsp struct{}

type GetMenuListWithDetailReq struct{}

type GetMenuListWithDetailRsp struct {
	List []*OneMenu `json:"list"`
}

type MfType string

const (
	MENU     MfType = "MENU"
	FUNCTION MfType = "FUNCTION"
)

func (t MfType) Check() error {
	switch t {
	case MENU, FUNCTION:
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
