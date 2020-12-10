package cerror

import (
	"errors"
)

/*
error定义，除了敏感错误类型，其他错误都应该使用 text 尽可能表达清楚错误类型（text会直接传到前端）；
-- 注意内部err的 text 定义格式应该是 [system err NUMBER]
*/

// Core error
var (
	ErrSys              = errors.New("[system err]")
	ErrMysql            = errors.New("[system err 001]")
	ErrRedis            = errors.New("[system err 002]")
	ErrSySConfig        = errors.New("[system config err]")
	ErrJsonUnmarshal    = errors.New("[json unmarshal err]")
	ErrExtractReqParams = errors.New("[extract req params err]")
	ErrParams           = errors.New("[params err]")
	ErrNotAllowed       = errors.New("[operation not be allowed]")
	ErrResourceNotFound = errors.New("[resource not found]")
)

// Business error
var ()
