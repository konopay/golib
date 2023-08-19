package logs

import (
	"context"
	"fmt"
	"os"

	KiteXLogs "github.com/cloudwego/kitex/pkg/klog"
	"github.com/konopay/golib/idgen"
)

// SetStdOutputAsLogDir
//
// 将stdout 设置为日志输出路径 便于采集
func SetStdOutputAsLogDir() {
	KiteXLogs.SetOutput(os.Stdout)
}

func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	KiteXLogs.CtxErrorf(ctx, AppendLogId(ctx, format), v...)
}
func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	KiteXLogs.CtxWarnf(ctx, AppendLogId(ctx, format), v...)
}
func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	KiteXLogs.CtxInfof(ctx, AppendLogId(ctx, format), v...)
}
func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	KiteXLogs.CtxDebugf(ctx, AppendLogId(ctx, format), v...)
}

// CtxDebugf
//
// 追加log-id到日志内容中 方便搜索上下文
func AppendLogId(ctx context.Context, format string) string {
	logId := idgen.LogIdOrDefault(ctx)
	return fmt.Sprintf("%s %s", logId, format)
}
