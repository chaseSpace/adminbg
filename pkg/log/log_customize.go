package log

import (
	"adminbg/pkg/util"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

// 直接基于标准库log封装  TODO: add level, print with level
type Clogger struct {
	writers []*log.Logger
}

func NewClogger(fpath string, toStdout bool) *Clogger {
	cl := &Clogger{}

	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND, 0755)
	util.PanicIfErr(err, nil)

	defLogger := log.New(f, "", log.Ldate|log.Ltime)
	cl.AddWriter(defLogger)

	if toStdout {
		stdlogger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
		cl.AddWriter(stdlogger)
	}
	return cl
}

func (c *Clogger) Debugf(format string, v ...interface{}) {
	v = append([]interface{}{c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Printf(format, v...)
	})
}

func (c *Clogger) Debugln(v ...interface{}) {
	v = append([]interface{}{c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Println(v...)
	})
}

func (c *Clogger) Infof(format string, v ...interface{}) {
	v = append([]interface{}{c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Printf(format, v...)
	})
}

func (c *Clogger) Infoln(v ...interface{}) {
	v = append([]interface{}{c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Println(v...)
	})
}

func (c *Clogger) Panicf(format string, v ...interface{}) {
	format = fmt.Sprintf("%s %s", c.withCallerLoc(), format)
	c.LoopDo(func(l *log.Logger) {
		l.Panicf(format, v...)
	})
}

func (c *Clogger) Panicln(v ...interface{}) {
	v = append([]interface{}{c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Panicln(v...)
	})
}

func (c *Clogger) Errorf(format string, v ...interface{}) {
	format = fmt.Sprintf("%s %s", c.withCallerLoc(), format)
	c.LoopDo(func(l *log.Logger) {
		l.Printf(format, v...)
		l.Println(debug.Stack())
	})
}

func (c *Clogger) Errorln(v ...interface{}) {
	v = append([]interface{}{c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Println(v...)
		l.Println(debug.Stack())
	})
}

// =================== 分割线 ==========================

func (c *Clogger) withCallerLoc() string {
	return util.FileWithLineNum(4)
}

func (c *Clogger) AddWriter(writer *log.Logger) {
	c.writers = append(c.writers, writer)
}

func (c *Clogger) LoopDo(fc func(*log.Logger)) {
	for _, w := range c.writers {
		fc(w)
	}
}
