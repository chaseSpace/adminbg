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
	Mysql *gorm.DB
)

func initDB() {
	gormConf := &gorm.Config{}
	mysql := Conf.Mysql
	clogger := Conf.Logger

	var err error
	Mysql, err = gorm.Open(mysqldriver.Open(mysql.Source), gormConf)
	util.PanicIfErr(err, nil)

	Mysql.Logger = log.NewGormLogger(clogger)
}
