package common

import (
	"adminbg/cerror"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
)

const (
	GinCtxKey_UID = "UID"
)

type HttpRsp struct {
	Data interface{} `json:"data"`
	Tips string      `json:"tips"` // common msg
	Succ bool        `json:"succ"`
}

func ExtractReqParams(c *gin.Context, req interface{}) (*HttpRsp, error) {
	rsp := &HttpRsp{Tips: "extract req params success"}

	if reflect.ValueOf(req).Elem().NumField() == 0 {
		return rsp, nil
	}
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	err := c.ShouldBind(req)
	if err != nil {
		err = errors.Wrap(cerror.ErrExtractReqParams, err.Error())
		rsp.Tips = err.Error()
		return rsp, err
	}
	return rsp, nil
}

// It should be called in the begin of API, then be processed by middleware when panics.
func MustExtractReqParams(c *gin.Context, req interface{}) {
	_, err := ExtractReqParams(c, req)
	if err != nil {
		panic(err)
	}
}

// The params `data` can be a pointer or not. If the data size is a little big, it's better to be a pointer.
func SetRsp(c *gin.Context, err error, data ...interface{}) {
	var r = &HttpRsp{}
	// Note: We must to use reflect to evaluate if data[0] is nil, because `data[0] == nil` would always be false.
	if len(data) > 0 && !reflect.ValueOf(data[0]).IsNil() {
		r.Data = data[0]
	} else {
		r.Data = &struct{}{}
	}
	// If code=200, frontend make a green tip popup, otherwise, make a rea warn tip popup;
	// the content of popup is rsp.Tips field.
	code := http.StatusOK
	if err != nil {
		if errors.Is(err, cerror.ErrParams) {
			code = http.StatusBadRequest
		} else if errors.Is(err, cerror.ErrUnauthorized) {
			code = http.StatusUnauthorized
		} else if errors.Is(err, cerror.ErrNotAllowed) {
			code = http.StatusForbidden
		} else if errors.Is(err, cerror.ErrNothingUpdated) ||
			errors.Is(err, cerror.ErrNothingDeleted) {
			// Keep 200 OK
			r.Succ = true
		} else { // It can be expanded here
			code = http.StatusBadRequest
		}
		r.Tips = err.Error()
	} else {
		r.Tips = "success"
		r.Succ = true
	}
	c.AbortWithStatusJSON(code, r)
}
