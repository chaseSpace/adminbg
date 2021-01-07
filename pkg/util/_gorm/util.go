package _gorm

import "gorm.io/gorm"

func IsDBErr(err error) bool {
	if err != nil && err != gorm.ErrRecordNotFound {
		return true
	}
	return false
}
