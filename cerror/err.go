package cerror

import (
	"errors"
)

// Basic error
var (
	ErrSys                    = errors.New("[System err]")
	ErrMysql                  = errors.New("[System err: mysql]")
	ErrRedis                  = errors.New("[System err: redis]")
	ErrSysConfig              = errors.New("[System err: config]")
	ErrJsonUnmarshal          = errors.New("[Data err: json unmarshal err]")
	ErrExtractReqParams       = errors.New("[Extract req params err]")
	ErrParams                 = errors.New("[Params err]")
	ErrUnauthorized           = errors.New("[Unauthorized request]")
	ErrNotAllowed             = errors.New("[Operation is not allowed]")
	ErrResourceNotFound       = errors.New("[Resource not found(:refresh page)]")
	ErrNothingUpdated         = errors.New("[Nothing updated(:may be some expired params)]")
	ErrNothingDeleted         = errors.New("[Nothing deleted]")
	ErrInvalidSplitPageParams = errors.New("[Invalid split page params]")
	ErrReservedResource       = errors.New("[Warning: reserved resource]")
	ErrCantOptReservedData    = errors.New("[Err: can't operate reserved data]")
	ErrYouHaveDoneIt          = errors.New("[You've done it]")
)

// Business error
var (
	ErrIncorrectInfoProvided = errors.New("[incorrect info provided]")
)
