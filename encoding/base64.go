package encoding

import "encoding/base64"

// Base64 将字节切片编码为标准 base64 字符串。
func Base64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// UnBase64 解码标准 base64 字符串为字节切片。
func UnBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Base64URL 将字节切片编码为 URL 安全的 base64 字符串。
func Base64URL(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

// UnBase64URL 解码 URL 安全的 base64 字符串为字节切片。
func UnBase64URL(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

// Base64RawURL 将字节切片编码为 URL 安全且无 padding 的 base64 字符串。
func Base64RawURL(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

// UnBase64RawURL 解码 URL 安全且无 padding 的 base64 字符串为字节切片。
func UnBase64RawURL(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
