package log

import (
	"adminbg/pkg/util"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

// 直接基于标准库log封装
type Clogger struct {
	writers []*log.Logger
	level   Level
}

type Level int8

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

func NewClogger(logPath string, level string, toStdout bool) *Clogger {
	cl := &Clogger{
		level: validLevel(level),
	}

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND, 0755)
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
	if c.level >= DEBUG {
		v = append([]interface{}{c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
		})
	}
}

func (c *Clogger) Debugln(v ...interface{}) {
	if c.level >= DEBUG {
		v = append([]interface{}{c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
		})
	}
}

func (c *Clogger) Infof(format string, v ...interface{}) {
	if c.level >= INFO {
		v = append([]interface{}{c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
		})
	}
}

func (c *Clogger) Infoln(v ...interface{}) {
	if c.level >= INFO {
		v = append([]interface{}{c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
		})
	}
}

func (c *Clogger) Warnf(format string, v ...interface{}) {
	if c.level >= WARN {
		v = append([]interface{}{c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
		})
	}
}

func (c *Clogger) Warnln(v ...interface{}) {
	if c.level >= WARN {
		v = append([]interface{}{c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
		})
	}
}

func (c *Clogger) Errorf(format string, v ...interface{}) {
	if c.level >= ERROR {
		format = fmt.Sprintf("%s %s", c.withCallerLoc(), format)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
			l.Println(debug.Stack())
		})
	}
}

func (c *Clogger) Errorln(v ...interface{}) {
	if c.level >= ERROR {
		v = append([]interface{}{c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
			l.Println(debug.Stack())
		})
	}
}

// Panic不受level管控，与标准库log.Panic行为一致
func (c *Clogger) Panicf(format string, v ...interface{}) {
	format = fmt.Sprintf("%s %s", c.withCallerLoc(), format)
	c.LoopDo(func(l *log.Logger) {
		l.Panicf(format, v...)
	})
}

// Panic不受level管控，与标准库log.Panic行为一致
func (c *Clogger) Panicln(v ...interface{}) {
	v = append([]interface{}{c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Panicln(v...)
	})
}

// =================== util ==========================

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

func validLevel(level string) Level {
	switch level {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	default:
		println(fmt.Sprintf("/pkg/log: invalid level str:%s; set to DEBUG", level))
		return DEBUG
	}
}
