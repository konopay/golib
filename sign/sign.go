package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/konopay/golib/errs"
)

func SignRequest(payload, secretKey string) (string, errs.Error) {
	key := []byte(secretKey)
	mac := hmac.New(sha256.New, key)
	if _, err := mac.Write([]byte(payload)); err != nil {
		return "", errs.SystemError
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}
