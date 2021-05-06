package cerror

import (
	"errors"
)

// Basic error
var (
	ErrSys                    = errors.New("[system err]")
	ErrMysql                  = errors.New("[system err: mysql]")
	ErrRedis                  = errors.New("[system err: redis]")
	ErrSysConfig              = errors.New("[system err: config]")
	ErrJsonUnmarshal          = errors.New("[data err: json unmarshal err]")
	ErrExtractReqParams       = errors.New("[extract req params err]")
	ErrParams                 = errors.New("[params err]")
	ErrUnauthorized           = errors.New("[unauthorized request]")
	ErrNotAllowed             = errors.New("[operation is not allowed]")
	ErrResourceNotFound       = errors.New("[resource not found(:refresh page)]")
	ErrNothingUpdated         = errors.New("[nothing updated(:may be some invalid params)]")
	ErrNothingDeleted         = errors.New("[nothing deleted]")
	ErrInvalidSplitPageParams = errors.New("[invalid split page params]")
)

// Business error
var (
	ErrIncorrectInfoProvided = errors.New("[incorrect info provided]")
)
