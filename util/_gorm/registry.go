package _gorm

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
)

/*
MySQL registry with gorm, app can directly call them to initialize MySQL gorm DB client.
*/

var (
	lock      sync.Mutex
	DefGormDB *gorm.DB
)

// init default pool on registry
func MustInitDef(dialer gorm.Dialector, conf *gorm.Config) {
	lock.Lock()
	defer lock.Unlock()
	if DefGormDB != nil {
		panic("_gorm: DefGormDB already exists")
	}

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(dialer, conf)
	if err != nil {
		panic(fmt.Sprintf("_gorm: gorm.Open err:%v", err))
	}
	DefGormDB = db
}

func MustInit(dialer gorm.Dialector, conf *gorm.Config) *gorm.DB {
	db, err := gorm.Open(dialer, conf)
	if err != nil {
		panic(fmt.Sprintf("_gorm: gorm.Open err:%v", err))
	}
	return db
}

func Close() error {
	lock.Lock()
	defer lock.Unlock()
	db, _ := DefGormDB.DB()
	err := db.Close()
	DefGormDB = nil
	return err
}
