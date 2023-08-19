package logs

import (
	"context"
	"testing"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/konopay/golib/idgen"
)

func TestPrintLogId(t *testing.T) {
	ctx := context.Background()
	ctx = metainfo.WithPersistentValue(ctx, idgen.LogIdPersistentKey, "test-log-id")

	CtxErrorf(ctx, "test-log-content")

}
