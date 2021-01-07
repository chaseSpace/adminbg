package g

import (
	"adminbg/config"
	"adminbg/pkg/log"
	"adminbg/pkg/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Conf *config.Conf

var (
	Mysql *gorm.DB
)

func mustInitDB() {
	gormConf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	cdb := Conf.Database
	clogger := Conf.Logger

	var err error
	Mysql, err = gorm.Open(mysql.Open(cdb.Mysql.Source), gormConf)
	util.PanicIfErr(err, nil)

	l := log.NewGormLogger(clogger)

	Mysql.Logger = l
}
