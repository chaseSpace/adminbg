package log

import (
	"adminbg/util"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

// 直接基于标准库log封装
type Clogger struct {
	writers   []*log.Logger
	level     Level
	shortPath bool
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
		level:     validLevel(level),
		shortPath: true,
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
	if c.level <= DEBUG {
		format = fmt.Sprintf("[DEBUG] %s %s", c.withCallerLoc(), format)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
		})
	}
}

func (c *Clogger) Debugln(v ...interface{}) {
	if c.level <= DEBUG {
		v = append([]interface{}{"[DEBUG]", c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
		})
	}
}

func (c *Clogger) Infof(format string, v ...interface{}) {
	if c.level <= INFO {
		format = fmt.Sprintf("[INFO] %s %s", c.withCallerLoc(), format)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
		})
	}
}

func (c *Clogger) Infoln(v ...interface{}) {
	if c.level <= INFO {
		v = append([]interface{}{"[INFO]", c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
		})
	}
}

func (c *Clogger) Warnf(format string, v ...interface{}) {
	if c.level <= WARN {
		format = fmt.Sprintf("[WARN] %s %s", c.withCallerLoc(), format)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
		})
	}
}

func (c *Clogger) Warnln(v ...interface{}) {
	if c.level <= WARN {
		v = append([]interface{}{"[WARN]", c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
		})
	}
}

func (c *Clogger) Errorf(format string, v ...interface{}) {
	if c.level <= ERROR {
		format = fmt.Sprintf("[ERROR] %s %s", c.withCallerLoc(), format)
		c.LoopDo(func(l *log.Logger) {
			l.Printf(format, v...)
			l.Println(string(debug.Stack()))
		})
	}
}

func (c *Clogger) Errorln(v ...interface{}) {
	if c.level <= ERROR {
		v = append([]interface{}{"[ERROR]", c.withCallerLoc()}, v...)
		c.LoopDo(func(l *log.Logger) {
			l.Println(v...)
			l.Println(string(debug.Stack()))
		})
	}
}

// Panicf is not controlled by level and is consistent with the standard library log.Panic behavior
func (c *Clogger) Panicf(format string, v ...interface{}) {
	format = fmt.Sprintf("[PANIC] %s %s", c.withCallerLoc(), format)
	c.LoopDo(func(l *log.Logger) {
		l.Panicf(format, v...)
	})
}

// Panicln is not controlled by level and is consistent with the standard library log.Panic behavior
func (c *Clogger) Panicln(v ...interface{}) {
	v = append([]interface{}{"[PANIC]", c.withCallerLoc()}, v...)
	c.LoopDo(func(l *log.Logger) {
		l.Panicln(v...)
	})
}

// =================== util ==========================

func (c *Clogger) withCallerLoc() string {
	fpath := util.FileWithLineNum(4)
	if !c.shortPath {
		return fpath
	}
	ss := strings.Split(fpath, "/")
	rootPath, _ := filepath.Abs(".")
	if os.PathSeparator == '\\' { // windows
		rootPath = strings.Replace(rootPath, "\\", "/", -1)
	}
	s := strings.Join(ss, "/")
	return strings.TrimPrefix(s, rootPath)
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
		println(fmt.Sprintf("/adminbg/log: invalid level str:%s; set to DEBUG", level))
		return DEBUG
	}
}
