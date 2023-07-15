package logger

import (
	"context"
	"errors"
	"fmt"
	gormLogger "gorm.io/gorm/logger"
	gormUtils "gorm.io/gorm/utils"
	"time"
)

const (
	logTitle      = "[gorm] "
	messageFormat = logTitle + "%s, %s"
	slowThreshold = 200 * time.Millisecond
	traceStr      = logTitle + "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr   = "%s %s\n[%.3fms] [rows:%v] %s"
)

type Gorm interface {
	LogMode(level gormLogger.LogLevel) gormLogger.Interface
	Info(ctx context.Context, msg string, data ...interface{})
	Warn(ctx context.Context, msg string, data ...interface{})
	Error(ctx context.Context, msg string, data ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

// LogMode The log level of gorm logger is overwrited by the log level of Zap logger.
func (log *apiLogger) LogMode(_ gormLogger.LogLevel) gormLogger.Interface {
	return log
}

// Info prints a information log.
func (log *apiLogger) Info(_ context.Context, msg string, data ...interface{}) {
	log.Infof(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Warn prints a warning log.
func (log *apiLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	log.Warnf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Error prints a error log.
func (log *apiLogger) Error(_ context.Context, msg string, data ...interface{}) {
	log.Errorf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Trace prints a trace log such as sql, source file and error.
func (log *apiLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, gormLogger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			log.Errorf(traceErrStr, gormUtils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.Errorf(traceErrStr, gormUtils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > slowThreshold:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		if rows == -1 {
			log.Warnf(traceWarnStr, gormUtils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.Warnf(traceWarnStr, gormUtils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	default:
		sql, rows := fc()
		if rows == -1 {
			log.Debugf(traceStr, gormUtils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.Debugf(traceStr, gormUtils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
