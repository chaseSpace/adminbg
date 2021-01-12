package model

import (
	"adminbg/cerror"
	"adminbg/pkg/g"
)

func CheckUserPassport(userName string, plainPwd string) (*User, error) {
	sql := `
		select *
		from adminbg_user
		where user_name = ?
		  and encrypted_pwd = sha1(concat(
				?, (select salt from adminbg_user where user_name = ? limit 1))
			);
		`
	row := User{}
	err := g.Mysql.Raw(sql, userName, plainPwd, userName).Scan(&row).Error
	return &row, cerror.WrapMysqlErr(err)
}
