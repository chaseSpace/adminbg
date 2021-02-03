package cerror

import (
	"errors"
)

// Core error
var (
	ErrSys              = errors.New("[system err]")
	ErrMysql            = errors.New("[system err 001]")
	ErrRedis            = errors.New("[system err 002]")
	ErrSySConfig        = errors.New("[system config err]")
	ErrJsonUnmarshal    = errors.New("[json unmarshal err]")
	ErrExtractReqParams = errors.New("[extract req params err]")
	ErrParams           = errors.New("[params err]")
	ErrUnauthorized     = errors.New("[unauthorized request]")
	ErrNotAllowed       = errors.New("[operation is not allowed]")
	ErrResourceNotFound = errors.New("[resource not found]")
	ErrNothingUpdated   = errors.New("[nothing updated(may be some invalid params)]")
)

// Business error
var (
	ErrIncorrectInfoProvided = errors.New("[incorrect info provided]")
)
