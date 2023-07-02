package idgen

import (
	"github.com/rs/xid"
)

// 希望能把id的版本
func NextId() string {
	return xid.New().String() // 20char
}
