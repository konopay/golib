package idgen

import (
	"context"
	"hash/crc32"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

const (
	logIdPersistentKey        = "logid"
	machinePosition    uint64 = 0xffff000000000000
	pidCodePosition    uint64 = 0x0000ff0000000000
	secondPosition     uint64 = 0x000000ffffff0000
	incPosition        uint64 = 0x000000000000ffff
)

var (
	defaultLogID = NewLogID()
)

// GetID return a uint64 number
func GetID() uint64 {
	return defaultLogID.GetID()
}

// machineCode: 16bit
// pidCode: 8bit
// time: 24bit
// inc: 16bit
type LogID struct {
	machineCode uint64
	pidCode     uint64
	second      int64
	inc         uint64
}

func NewLogID() *LogID {
	ld := &LogID{
		second: time.Now().Unix(),
		inc:    0,
	}
	ld.machineCode, ld.pidCode = ld.machinePidCode()
	return ld
}

func (ld *LogID) machinePidCode() (uint64, uint64) {
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	m32Code := crc32.ChecksumIEEE([]byte(hostname + strconv.Itoa(pid)))
	return uint64(m32Code), uint64(pid)
}

func (ld *LogID) GetID() uint64 {
	now := time.Now().Unix()
	inc := atomic.AddUint64(&ld.inc, 1)
	ts := uint64(now)
	var ret uint64
	ret = (ld.machineCode << 48) & machinePosition
	ret = ret | ((ld.pidCode << 40) & pidCodePosition)
	ret = ret | ((ts << 16) & secondPosition)
	ret = ret | (inc & incPosition)
	return ret
}

func AddLogIdKiteXMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		logid, ok := metainfo.GetPersistentValue(ctx, logIdPersistentKey)
		if !ok || logid == "" {
			ctx = metainfo.WithPersistentValue(ctx, logIdPersistentKey, strconv.FormatUint(NewLogID().GetID(), 10))
		}

		return next(ctx, req, resp)
	}
}

// GetLogIdOrDefaultFromContext
//
// 从 ctx 中获取log-id 如果不存在 返回默认的id
func GetLogIdOrDefaultFromContext(ctx context.Context) string {
	logid, ok := metainfo.GetPersistentValue(ctx, logIdPersistentKey)
	if !ok {
		return "unknown-logid"
	}

	return logid
}
