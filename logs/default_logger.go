package logs

// 从 cloudwego copy 过来进行了小幅度修改
// 帮业务自动将ctx中的log-id注入到 日志中 方便检索分析

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/konopay/golib/idgen"
)

var logger klog.FullLogger = &defaultLogger{
	level: klog.LevelInfo,
	// 打印完整调用路径 方便定位问题
	// 如果发现比较费钱 再改回short-file
	stdlog: log.New(os.Stdout, "", log.LstdFlags|log.Llongfile|log.Lmicroseconds), // stdout => SLS
}

// SetOutput sets the output of default logger. By default, it is stderr.
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

// SetLevel sets the level of logs below which logs will not be output.
// The default log level is LevelTrace.
// Note that this method is not concurrent-safe.
func SetLevel(lv klog.Level) {
	logger.SetLevel(lv)
}

// DefaultLogger return the default logger for kitex.
func DefaultLogger() klog.FullLogger {
	return logger
}

// SetLogger sets the default logger.
// Note that this method is not concurrent-safe and must not be called
// after the use of DefaultLogger and global functions in this package.
func SetLogger(v klog.FullLogger) {
	logger = v
}

// Fatal calls the default logger's Fatal method and then os.Exit(1).
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Error calls the default logger's Error method.
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Warn calls the default logger's Warn method.
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Notice calls the default logger's Notice method.
func Notice(v ...interface{}) {
	logger.Notice(v...)
}

// Info calls the default logger's Info method.
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Debug calls the default logger's Debug method.
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Trace calls the default logger's Trace method.
func Trace(v ...interface{}) {
	logger.Trace(v...)
}

// Fatalf calls the default logger's Fatalf method and then os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

// Errorf calls the default logger's Errorf method.
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Warnf calls the default logger's Warnf method.
func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Noticef calls the default logger's Noticef method.
func Noticef(format string, v ...interface{}) {
	logger.Noticef(format, v...)
}

// Infof calls the default logger's Infof method.
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Debugf calls the default logger's Debugf method.
func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Tracef calls the default logger's Tracef method.
func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

// CtxFatalf calls the default logger's CtxFatalf method and then os.Exit(1).
func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxFatalf(ctx, format, v...)
}

// CtxErrorf calls the default logger's CtxErrorf method.
func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxErrorf(ctx, format, v...)
}

// CtxWarnf calls the default logger's CtxWarnf method.
func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxWarnf(ctx, format, v...)
}

// CtxNoticef calls the default logger's CtxNoticef method.
func CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	logger.CtxNoticef(ctx, format, v...)
}

// CtxInfof calls the default logger's CtxInfof method.
func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	logger.CtxInfof(ctx, format, v...)
}

// CtxDebugf calls the default logger's CtxDebugf method.
func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxDebugf(ctx, format, v...)
}

// CtxTracef calls the default logger's CtxTracef method.
func CtxTracef(ctx context.Context, format string, v ...interface{}) {
	logger.CtxTracef(ctx, format, v...)
}

type defaultLogger struct {
	stdlog *log.Logger
	level  klog.Level
}

func (ll *defaultLogger) SetOutput(w io.Writer) {
	ll.stdlog.SetOutput(w)
}

func (ll *defaultLogger) SetLevel(lv klog.Level) {
	ll.level = lv
}

var strs = []string{
	"[Trace] ",
	"[Debug] ",
	"[Info] ",
	"[Notice] ",
	"[Warn] ",
	"[Error] ",
	"[Fatal] ",
}

func toString(lv klog.Level) string {
	if lv >= klog.LevelTrace && lv <= klog.LevelFatal {
		return strs[lv]
	}
	return fmt.Sprintf("[?%d] ", lv)
}
func (ll *defaultLogger) logf(ctx context.Context, lv klog.Level, format *string, v ...interface{}) {
	if ll.level > lv {
		return
	}
	// 追加log-id 到content 中
	logId := idgen.LogIdOrDefault(ctx)
	msg := logId + toString(lv)
	if format != nil {
		msg += fmt.Sprintf(*format, v...)
	} else {
		msg += fmt.Sprint(v...)
	}
	ll.stdlog.Output(4, msg)
	if lv == klog.LevelFatal {
		os.Exit(1)
	}
}

func (ll *defaultLogger) Fatal(v ...interface{}) {
	ll.logf(context.Background(), klog.LevelFatal, nil, v...)
}

func (ll *defaultLogger) Error(v ...interface{}) {
	ll.logf(context.Background(), klog.LevelError, nil, v...)
}

func (ll *defaultLogger) Warn(v ...interface{}) {
	ll.logf(context.Background(), klog.LevelWarn, nil, v...)
}

func (ll *defaultLogger) Notice(v ...interface{}) {
	ll.logf(context.Background(), klog.LevelNotice, nil, v...)
}

func (ll *defaultLogger) Info(v ...interface{}) {
	ll.logf(context.Background(), klog.LevelInfo, nil, v...)
}

func (ll *defaultLogger) Debug(v ...interface{}) {
	ll.logf(context.Background(), klog.LevelDebug, nil, v...)
}

func (ll *defaultLogger) Trace(v ...interface{}) {
	ll.logf(context.Background(), klog.LevelTrace, nil, v...)
}

func (ll *defaultLogger) Fatalf(format string, v ...interface{}) {
	ll.logf(context.Background(), klog.LevelFatal, &format, v...)
}

func (ll *defaultLogger) Errorf(format string, v ...interface{}) {
	ll.logf(context.Background(), klog.LevelError, &format, v...)
}

func (ll *defaultLogger) Warnf(format string, v ...interface{}) {
	ll.logf(context.Background(), klog.LevelWarn, &format, v...)
}

func (ll *defaultLogger) Noticef(format string, v ...interface{}) {
	ll.logf(context.Background(), klog.LevelNotice, &format, v...)
}

func (ll *defaultLogger) Infof(format string, v ...interface{}) {
	ll.logf(context.Background(), klog.LevelInfo, &format, v...)
}

func (ll *defaultLogger) Debugf(format string, v ...interface{}) {
	ll.logf(context.Background(), klog.LevelDebug, &format, v...)
}

func (ll *defaultLogger) Tracef(format string, v ...interface{}) {
	ll.logf(context.Background(), klog.LevelTrace, &format, v...)
}

func (ll *defaultLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(ctx, klog.LevelFatal, &format, v...)
}

func (ll *defaultLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(ctx, klog.LevelError, &format, v...)
}

func (ll *defaultLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(ctx, klog.LevelWarn, &format, v...)
}

func (ll *defaultLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(ctx, klog.LevelNotice, &format, v...)
}

func (ll *defaultLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	ll.logf(ctx, klog.LevelInfo, &format, v...)
}

func (ll *defaultLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(ctx, klog.LevelDebug, &format, v...)
}

func (ll *defaultLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(ctx, klog.LevelTrace, &format, v...)
}
