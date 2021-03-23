package test_octal

import (
	"regexp"
	"strconv"
)

// 转义
func EscapeOctalString(src string) string {
	reg := regexp.MustCompile("\\\\[0-7]{3}")
	// dest := reg.ReplaceAllFunc([]byte(src), func(repl []byte) []byte {
	// 	i, _ := strconv.ParseInt(string(repl[1:]), 8, 64)
	// 	return []byte{byte(i)}
	// })
	// return string(dest)
	dest := reg.ReplaceAllStringFunc(src, func(repl string) string {
		i, _ := strconv.ParseInt(string(repl[1:]), 8, 64)
		return string([]byte{byte(i)})
	})

	return dest
}
