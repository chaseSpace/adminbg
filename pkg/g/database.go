package g

import (
	"adminbg/config"
	"adminbg/pkg/_util"
	"adminbg/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Conf *config.Conf

var (
	Mysql *gorm.DB
	CkDB  *gorm.DB
)

func MustInit() {
	gormConf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	cdb := Conf.Database
	clogger := Conf.Logger

	var err error
	Mysql, err = gorm.Open(mysql.Open(cdb.Mysql.Source), gormConf)
	_util.PanicIfErr(err, nil)

	l := log.NewGormLogger(clogger)

	Mysql.Logger = l

	// 连不上?
	//CkDB, err = gorm.Open(mysql.Open(cdb.Clickhouse.Source), gormConf)
	//_util.PanicIfErr(err, nil)
	//
	//fmt.Println("CK OK!")
	//CkDB.Logger = l
}

func Stop() {
	if Mysql != nil {
		db, _ := Mysql.DB()
		_ = db.Close()
	}
	if CkDB != nil {
		db, _ := CkDB.DB()
		_ = db.Close()
	}
}
