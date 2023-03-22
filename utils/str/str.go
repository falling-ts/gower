package str

import (
	"bytes"
	"unicode"
)

type Conv string

// Uppercase 首字符大写
func (s Conv) Uppercase() string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// Lowercase 首字母小写
func (s Conv) Lowercase() string {
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// Snake 获取蛇形字符
func (s Conv) Snake() string {
	var buf bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteByte('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}

	return Conv(buf.String()).Lowercase()
}

// Camel 获得小驼峰字符
func (s Conv) Camel() string {
	var buf bytes.Buffer
	toUpper := false

	for _, r := range s {
		if r == '_' {
			toUpper = true
		} else if toUpper {
			buf.WriteRune(unicode.ToUpper(r))
			toUpper = false
		} else {
			buf.WriteRune(r)
		}
	}

	return Conv(buf.String()).Lowercase()
}

// UpCamel 获得大驼峰字符
func (s Conv) UpCamel() string {
	return Conv(s.Camel()).Uppercase()
}
