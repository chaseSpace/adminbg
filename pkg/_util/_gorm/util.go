package _gorm

import "gorm.io/gorm"

func IsGormErr(err error) bool {
	if err != nil && err != gorm.ErrRecordNotFound {
		return true
	}
	return false
}
