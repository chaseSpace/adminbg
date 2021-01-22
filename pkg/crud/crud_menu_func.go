package crud

import (
	"adminbg/cproto"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"fmt"
	"gorm.io/gorm"
)

func InsertNewMenu(menu *model.MenuAndFunction) error {
	menu.Type = cproto.MENU
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
		SET a.path = concat(b.path, ?, '/')
		WHERE a.mf_id = ?
			AND b.mf_id = ?
	`, TN.MenuAndFunction, TN.MenuAndFunction)

	var err error
	err = g.MySQL.Transaction(func(tx *gorm.DB) error {
		insert := tx.Exec(insertSQL, menu.MfName, menu.ParentId, menu.Level, menu.Type, menu.MenuRoute,
			menu.MenuDisplay, menu.SortNum, menu.ParentId, menu.ParentId)
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
		err = tx.Exec(updateSQL, id, id, menu.ParentId).Error
		return err
	})
	return err
}
