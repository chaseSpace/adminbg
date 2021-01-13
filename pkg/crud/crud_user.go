package crud

import (
	"adminbg/cerror"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
)

func CheckUserPassport(userName string, plainPwd string) (*model.User, error) {
	// It use mysql SHA1 alg store encrypted password here.
	sql := `
		SELECT *
		FROM adminbg_user
		WHERE user_name = ?
			AND encrypted_pwd = SHA1(CONCAT(?, (
				SELECT salt
				FROM adminbg_user
				WHERE user_name = ?
				LIMIT 1
			)));
		`
	row := model.User{}
	err := g.Mysql.Raw(sql, userName, plainPwd, userName).Scan(&row).Error
	return &row, cerror.WrapMysqlErr(err)
}
