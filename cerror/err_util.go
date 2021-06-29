package cerror

import (
	"adminbg/util/_gorm"
	"adminbg/util/_redis"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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

func WrapDeleteDBOneRecordErr(exec *gorm.DB) error {
	if exec.Error != nil {
		return exec.Error
	}
	if exec.RowsAffected != 1 {
		return ErrNothingDeleted
	}
	return nil
}

func WrapUpdateDBOneRecordErr(exec *gorm.DB) error {
	if exec.Error != nil {
		return exec.Error
	}
	if exec.RowsAffected != 1 {
		return ErrNothingUpdated
	}
	return nil
}
