package crud

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

func InsertNewMenu(menu *model.MenuAndFunction) error {
	menu.Type = cproto.Menu
	if err := menu.Check(); err != nil {
		return err
	}
	// The premise of insert data is either parent_id equal to 100 or an
	// parent_id provided exists
	insertSQL := fmt.Sprintf(`
		INSERT INTO %s ( mf_name, parent_id, level, type
						  , menu_route, menu_display, sort_num, path)
		SELECT ?, ?, ?, ?, ?, ?, ?, ''
		WHERE EXISTS(
					  SELECT 1
					  FROM %s
					  WHERE parent_id = ?
				  )
	`, TN.MenuAndFunction, TN.MenuAndFunction)
	updateSQL := fmt.Sprintf(`
		UPDATE %s a
			JOIN %s b
		SET a.path = concat(b.path, a.mf_id, '/')
		WHERE a.mf_id = ?
			AND b.mf_id = ?
	`, TN.MenuAndFunction, TN.MenuAndFunction)

	var err error
	err = g.MySQL.Transaction(func(tx *gorm.DB) error {
		insert := tx.Exec(insertSQL, menu.MfName, menu.ParentId, menu.Level, menu.Type, menu.MenuRoute,
			menu.MenuDisplay, menu.SortNum, menu.ParentId)
		if insert.Error != nil {
			return insert.Error
		}
		if insert.RowsAffected == 0 {
			return fmt.Errorf("parent_id %d does not exist", menu.ParentId)
		}
		id, err := LastInsertId(tx)
		if err != nil {
			return err
		}
		err = tx.Exec(updateSQL, id, menu.ParentId).Error
		return err
	})
	return err
}

func UpdateMenu(menu *model.MenuAndFunction) error {
	menu.Type = cproto.Menu
	if err := menu.Check(); err != nil {
		return err
	}
	updateSQL := fmt.Sprintf(`
		UPDATE %s a
			LEFT JOIN %s b ON b.mf_id = ?
		SET a.path = concat(b.path, a.mf_id, '/'), 
			a.parent_id = ?, 
			a.level = ?, 
			a.mf_name = ?, 
			a.menu_route = ?, 
			a.menu_display = ?, 
			a.sort_num = ?
		WHERE a.mf_id = ?
		`, TN.MenuAndFunction, TN.MenuAndFunction,
	)
	// cannot be null
	err := g.MySQL.Exec(updateSQL, menu.ParentId, menu.ParentId, menu.Level, menu.MfName, menu.MenuRoute, menu.MenuDisplay, menu.SortNum, menu.MfId).Error
	if err != nil && strings.Contains(err.Error(), "cannot be null") {
		// There might get error that is "column 'path' cannot be null", it means parent_id is not found.
		return errors.Wrap(cerror.ErrParams, fmt.Sprintf("parent_id %d not found", menu.ParentId))
	}
	return err
}

func GetMenuFuncList() ([]*model.MenuAndFunction, error) {
	var rows []*model.MenuAndFunction
	querySQL := fmt.Sprintf(`
		SELECT *
		FROM %s
		WHERE deleted_at IS NULL
		ORDER BY sort_num, created_at
		`, TN.MenuAndFunction,
	)
	err := g.MySQL.Raw(querySQL).Scan(&rows).Error
	return rows, err
}
