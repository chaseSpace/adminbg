package cproto

// POST /web/v1/NewAPI
type NewAPIReq struct {
	Identity string `json:"identity" binding:"required"`
	Remark   string `json:"remark" binding:"required"`
}

type NewAPIRsp struct{}

// POST /web/v1/UpdateAPI
type UpdateAPIReq struct {
	Delete       bool    `json:"delete"`         // if true, the api execute delete action, or update action if false
	DeleteApiIds []int32 `json:"delete_api_ids"` // if delete is true, the api delete ids from `delete_api_ids`, instead of `api_id`
	ApiId        int32   `json:"api_id"`
	Identity     string  `json:"identity"`
	Remark       string  `json:"remark"`
}

type UpdateAPIRsp struct{}
