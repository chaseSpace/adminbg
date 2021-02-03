package crud

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"fmt"
	"gorm.io/gorm"
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
					  WHERE mf_id = ?
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
			JOIN %s b
		SET a.path = concat(b.path, a.mf_id, '/'), 
			a.parent_id = ?, 
			a.level = ?, 
			a.mf_name = ?, 
			a.menu_route = ?, 
			a.menu_display = ?, 
			a.sort_num = ?
		WHERE a.mf_id = ?
		AND b.mf_id = ?
		`, TN.MenuAndFunction, TN.MenuAndFunction,
	)
	update := g.MySQL.Exec(updateSQL, menu.ParentId, menu.Level, menu.MfName, menu.MenuRoute, menu.MenuDisplay,
		menu.SortNum, menu.MfId, menu.ParentId)
	if update.Error != nil {
		return update.Error
	}
	if update.RowsAffected == 0 {
		return cerror.ErrNothingUpdated
	}
	return nil
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

func DeleteMenus(ids []int32) error {
	updateSQL := fmt.Sprintf(`
		UPDATE %s a
			JOIN %s b
			ON b.mf_id in ?
		SET a.deleted_at = NOW()
		WHERE a.deleted_at IS NULL
			AND a.path LIKE concat(b.path, '%%');
	`, TN.MenuAndFunction, TN.MenuAndFunction)
	err := g.MySQL.Exec(updateSQL, ids).Error

	return err
}

func InsertNewFunc(fc *model.MenuAndFunction) error {
	fc.Type = cproto.Function
	if err := fc.Check(); err != nil {
		return err
	}
	// Functions must be created under the leaf-menu, so we must limit menu's level
	insertSQL := fmt.Sprintf(`
		INSERT INTO %s (mf_name, type, parent_id, sort_num, path)
		SELECT ?, ?, ?, ?, ''
		WHERE EXISTS (
			SELECT 1
			FROM %s
			WHERE deleted_at IS NULL
				AND mf_id = ?
				AND type = 'MENU'
				AND level = ?
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
		insert := tx.Exec(insertSQL, fc.MfName, fc.Type, fc.ParentId, fc.SortNum, fc.ParentId, model.MaxMenuLevel)
		if insert.Error != nil {
			return insert.Error
		}
		if insert.RowsAffected == 0 {
			// Check `insertSQL` above for how `RowsAffected` is 0
			return fmt.Errorf("invalid parent_id %d", fc.ParentId)
		}
		id, err := LastInsertId(tx)
		if err != nil {
			return err
		}
		err = tx.Exec(updateSQL, id, fc.ParentId).Error
		return err
	})
	return err
}

func UpdateFunction(fc *model.MenuAndFunction) error {
	fc.Type = cproto.Function
	if err := fc.Check(); err != nil {
		return err
	}
	updateSQL := fmt.Sprintf(`
		UPDATE %s a
			JOIN %s b
		SET a.path = concat(b.path, a.mf_id, '/'), 
			a.parent_id = ?, 
			a.mf_name = ?, 
			a.sort_num = ?
		WHERE a.mf_id = ?
		AND a.type = 'FUNCTION'
		AND b.mf_id = ?
		AND b.level = ?
		`, TN.MenuAndFunction, TN.MenuAndFunction,
	)
	// cannot be null
	update := g.MySQL.Exec(updateSQL, fc.ParentId, fc.MfName, fc.SortNum, fc.MfId, fc.ParentId, model.MaxMenuLevel)
	if update.Error != nil {
		return update.Error
	}
	if update.RowsAffected == 0 {
		return cerror.ErrNothingUpdated
	}
	return nil
}
