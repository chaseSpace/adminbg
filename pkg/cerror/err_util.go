package cerror

import (
	"github.com/pkg/errors"
)

func WrapMysqlErr(err error) error {
	if _util.IsMysqlErr(err) {
		return errors.Wrap(ErrMysql, err.Error())
	}
	return nil
}

func WrapRedisErr(err error) error {
	if _util.IsRedisErr(err) {
		return errors.Wrap(ErrRedis, err.Error())
	}
	return nil
}
