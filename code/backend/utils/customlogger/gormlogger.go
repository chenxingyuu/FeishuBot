package customlogger

import (
	"context"
	"gorm.io/gorm/logger"
	"log"

	"time"
)

type GormCustomLogger struct {
	GormLogger logger.Interface
}

func (l *GormCustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.GormLogger = l.GormLogger.LogMode(level)
	return l
}

func (l *GormCustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.GormLogger.Info(ctx, msg, data...)
}

func (l *GormCustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.GormLogger.Warn(ctx, msg, data...)
}

func (l *GormCustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.GormLogger.Error(ctx, msg, data...)
}

func (l *GormCustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if err != nil {
		sql, rows := fc()
		log.Printf("[GORM] [%12v] [%19s] [rows:%3d] %s ERROR: %v | %v", time.Since(begin), "", rows, sql, err, ctx.Value("a"))
	} else {
		sql, rows := fc()
		log.Printf("[GORM] [%12v] [%19s] [rows:%3d] %s | %v", time.Since(begin), "", rows, sql, ctx.Value("a"))
	}
}
