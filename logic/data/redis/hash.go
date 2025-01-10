package redis

import (
	"crypto/md5"
	"encoding/hex"
)

func hashKey(sessionID string) string {
	hasher := md5.New()
	hasher.Write([]byte(sessionID))
	return hex.EncodeToString(hasher.Sum(nil)) // 返回32字符的十六进制字符串
}
