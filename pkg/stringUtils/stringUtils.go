package stringUtils

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

// CompressStr 利用正则表达式压缩字符串，去除空格或制表符
func CompressStr(str string) string {
	if str == "" {
		return ""
	}
	// 匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func Md5(origin string) string {
	h := md5.New()
	h.Write([]byte(origin))
	return hex.EncodeToString(h.Sum(nil))
}
