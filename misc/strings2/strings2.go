package strings2

import (
	"regexp"
	"strconv"
	"strings"
)

// \p{C} 匹配任何C类Unicode代码点，即所有控制字符和格式字符
var controlRe = regexp.MustCompile(`[\p{C}]+`)

// RemoveControlCharacters 移除字符串中的所有控制字符
func RemoveInvalidCharacters(s string) string {
	s = strings.TrimSpace(s)
	// 使用正则表达式替换所有控制字符为空字符串
	cleaned := controlRe.ReplaceAllString(s, "")
	return cleaned
}

func StringToInt(str string) int64 {
	n, _ := strconv.ParseInt(str, 10, 0)
	return n
}

func IntToString(orig []int8) string {
	ret := make([]byte, len(orig))
	size := -1
	for i, o := range orig {
		if o == 0 {
			size = i
			break
		}
		ret[i] = byte(o)
	}
	if size == -1 {
		size = len(orig)
	}
	return string(ret[0:size])
}

func UintToString(orig []uint8) string {
	ret := make([]byte, len(orig))
	size := -1
	for i, o := range orig {
		if o == 0 {
			size = i
			break
		}
		ret[i] = byte(o)
	}
	if size == -1 {
		size = len(orig)
	}

	return string(ret[0:size])
}

func ByteToString(orig []byte) string {
	n := -1
	l := -1
	for i, b := range orig {
		// skip left side null
		if l == -1 && b == 0 {
			continue
		}
		if l == -1 {
			l = i
		}

		if b == 0 {
			break
		}
		n = i + 1
	}
	if n == -1 {
		return string(orig)
	}
	return string(orig[l:n])
}
