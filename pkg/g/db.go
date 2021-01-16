package g

import (
	"adminbg/config"
	"adminbg/log"
	"adminbg/util"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conf *config.Conf

var (
	MySQL *gorm.DB
)

func initDB() {
	gormConf := &gorm.Config{
		AllowGlobalUpdate: false,
	}
	mysql := Conf.Mysql
	clogger := Conf.Logger

	var err error
	MySQL, err = gorm.Open(mysqldriver.Open(mysql.Source), gormConf)
	util.PanicIfErr(err, nil)

	MySQL.Logger = log.NewGormLogger(clogger)
}
