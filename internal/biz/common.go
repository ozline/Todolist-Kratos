package biz

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GenerateTokenSHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetTimestamp13() int64 {
	return time.Now().UnixNano() / 1e6
}
