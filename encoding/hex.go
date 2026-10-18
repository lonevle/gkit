package encoding

import (
	"encoding/hex"
)

// Hex 将字节切片编码为 hex 字符串。
func Hex(b []byte) string {
	return hex.EncodeToString(b)
}

// UnHex 解码 hex 字符串为字节切片。
func UnHex(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
