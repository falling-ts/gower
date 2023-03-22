package slice

import "strings"

type Strings []string

// Has 查找字符串是否存在
func (s Strings) Has(str string) bool {
	for _, item := range s {
		if item == str {
			return true
		}
	}

	return false
}

// HasPrefix 判断字符串集中是否有 str 的开头子串
func (s Strings) HasPrefix(str string) bool {
	for _, item := range s {
		if strings.HasPrefix(str, item) {
			return true
		}
	}

	return false
}

// Map 迭代处理
func (s Strings) Map(fn func(string) string) Strings {
	newS := make(Strings, len(s))
	for i, item := range s {
		newS[i] = fn(item)
	}

	return newS
}
