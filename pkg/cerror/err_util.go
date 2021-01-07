package cerror

import (
	"adminbg/pkg/util/_gorm"
	"adminbg/pkg/util/_redis"
	"github.com/pkg/errors"
)

func WrapMysqlErr(err error) error {
	if _gorm.IsDBErr(err) {
		return errors.Wrap(ErrMysql, err.Error())
	}
	return nil
}

func WrapRedisErr(err error) error {
	if _redis.IsRedisErr(err) {
		return errors.Wrap(ErrRedis, err.Error())
	}
	return nil
}
