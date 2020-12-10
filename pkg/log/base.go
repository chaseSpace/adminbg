package log

import (
	"adminbg/config"
	"adminbg/pkg/_util/_file"
	"path/filepath"
)

/*
注意log模块不能调用g模块，会循环依赖
*/
var (
	DefLogger *Clogger
	ReqLogger *Clogger
)

func MustInit(c config.Logger) {

	_ = _file.MkdirIfNotExist(c.Dir)

	defLogPath := filepath.Join(c.Dir, c.DefaultLogFilename)
	reqLogPath := filepath.Join(c.Dir, c.RequestLogFilename)

	DefLogger = NewClogger(defLogPath, c.ToStdout)
	ReqLogger = NewClogger(reqLogPath, c.ToStdout)
}

// 快捷方式，需要先调用MustInit
func Printf(format string, v ...interface{}) {
	DefLogger.Printf(format, v...)
}

func Println(v ...interface{}) {
	DefLogger.Println(v...)
}

func Panicf(format string, v ...interface{}) {
	DefLogger.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	DefLogger.Panicln(v...)
}

func Errorf(format string, v ...interface{}) {
	DefLogger.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	DefLogger.Errorln(v...)
}
