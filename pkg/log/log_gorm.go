package log

import (
	"adminbg/config"
	"adminbg/pkg/util"
	"adminbg/pkg/util/_file"
	"context"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 继承gorm的logger，添加行为
type LoggerGorm struct {
	mu      sync.Mutex // 这个logger会被共享，所以可能会对同一文件并发写
	writers []logger.Interface
}

func NewGormLogger(c config.Logger) *LoggerGorm {

	ormLogger := &LoggerGorm{}
	_ = _file.MkdirAllIfNotExist(c.Dir)

	f, err := os.OpenFile(filepath.Join(c.Dir, c.DBLogFilename), os.O_CREATE|os.O_APPEND, 0755)
	util.PanicIfErr(err, nil)

	w := logger.New(log.New(f, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Info, // info表示记录所有db log
		Colorful:      true,
	})
	ormLogger.AddWriter(w)

	if c.ToStdout {
		ormLogger.AddWriter(logger.Default)
	}
	return ormLogger
}

func (l *LoggerGorm) LogMode(level logger.LogLevel) logger.Interface {
	l.mu.Lock()
	defer l.mu.Unlock()
	var ret logger.Interface
	l.LoopDo(func(l logger.Interface) {
		r := l.LogMode(level)
		if ret == nil {
			ret = r
		}
	})
	return ret
}

func (l *LoggerGorm) Info(ctx context.Context, s string, m ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.LoopDo(func(l logger.Interface) {
		l.Info(ctx, s, m...)
	})
}

func (l *LoggerGorm) Warn(ctx context.Context, s string, m ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.LoopDo(func(l logger.Interface) {
		l.Warn(ctx, s, m...)
	})
}

func (l *LoggerGorm) Error(ctx context.Context, s string, m ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.LoopDo(func(l logger.Interface) {
		l.Error(ctx, s, m...)
	})
}

func (l *LoggerGorm) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.LoopDo(func(l logger.Interface) {
		l.Trace(ctx, begin, fc, err)
	})
}

func (l *LoggerGorm) AddWriter(writer logger.Interface) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writers = append(l.writers, writer)
}

func (l *LoggerGorm) LoopDo(fc func(logger.Interface)) {
	for _, w := range l.writers {
		fc(w)
	}
}
