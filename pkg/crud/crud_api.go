package crud

import (
	"adminbg/cerror"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"errors"
	"fmt"
)

func GetAPIList(bindFuncId int32, searchName string, orderByCreateAtDesc bool) ([]*model.Api, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at IS NULL %%s ORDER BY created_at", TN.Api)
	if bindFuncId != 0 {
		sql = fmt.Sprintf(`
			SELECT *
			FROM %s
			WHERE deleted_at IS NULL
				AND api_id IN (
					SELECT api_id
					FROM %s
					WHERE mf_id = %d
				)
				%%s
			ORDER BY created_at`, TN.Api, TN.MfApiRef, bindFuncId)
	}
	WHERE := ""
	if searchName != "" {
		// We do a full fuzzy matching search, because the number of APIs will not be too large.
		WHERE += fmt.Sprintf("AND name like '%%%s%%'", searchName)
	}
	sql = fmt.Sprintf(sql, WHERE)
	if orderByCreateAtDesc {
		sql += fmt.Sprintf(" DESC")
	}
	var ret []*model.Api
	err := g.MySQL.Raw(sql).Scan(&ret).Error
	return ret, cerror.WrapMysqlErr(err)
}

func NewAPI(identity, remark string) error {
	sql := fmt.Sprintf(`
			INSERT INTO %s (identity, remark)
			VALUES (?, ?)
			ON DUPLICATE KEY UPDATE identity = identity`, TN.Api)
	exec := g.MySQL.Exec(sql, identity, remark)
	if exec.Error != nil {
		return exec.Error
	}
	if exec.RowsAffected == 0 {
		return errors.New("api exists")
	}
	return nil
}

func UpdateAPI(id int32, identity, remark string) error {
	entity := &model.Api{
		Identity: identity,
		Remark:   remark,
	}
	exec := g.MySQL.Where("api_id=?", id).Select([]string{"identity", "remark"}).Updates(entity)
	if exec.Error != nil {
		return exec.Error
	}
	if exec.RowsAffected == 0 {
		return cerror.ErrNothingUpdated
	}
	return nil
}

func DeleteAPIs(id ...int32) error {
	exec := g.MySQL.Delete(new(model.Api), "api_id in (?)", id)
	if exec.Error != nil {
		return exec.Error
	}
	if exec.RowsAffected == 0 {
		return cerror.ErrNothingDeleted
	}
	return nil
}
