package log

import (
	"adminbg/config"
	"adminbg/util"
	"adminbg/util/_file"
	"path/filepath"
)

/*
注意log模块不能调用g模块，会循环依赖
*/
var (
	DefLogger *Clogger
	//ReqLogger *Clogger // 建议DefLogger够用，不需要增加过多日志文件
)

func MustInit(c config.Logger) {

	err := _file.MkdirAllIfNotExist(c.Dir)
	util.PanicIfErr(err, nil)

	defLogPath := filepath.Join(c.Dir, c.CommonLogFilename)
	//reqLogPath := filepath.Join(c.Dir, c.RequestLogFilename)

	DefLogger = NewClogger(defLogPath, c.Level, c.ToStdout)
	//ReqLogger = NewClogger(reqLogPath, c.ToStdout)
}

/*
快捷方式，需要先调用MustInit
*/
func Debugf(format string, v ...interface{}) {
	DefLogger.Debugf(format, v...)
}

func Debugln(v ...interface{}) {
	DefLogger.Debugln(v...)
}

func Infof(format string, v ...interface{}) {
	DefLogger.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	DefLogger.Infoln(v...)
}

func Warnf(format string, v ...interface{}) {
	DefLogger.Warnf(format, v...)
}

func Warnln(v ...interface{}) {
	DefLogger.Warnln(v...)
}

func Errorf(format string, v ...interface{}) {
	DefLogger.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	DefLogger.Errorln(v...)
}

func Panicf(format string, v ...interface{}) {
	DefLogger.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	DefLogger.Panicln(v...)
}
