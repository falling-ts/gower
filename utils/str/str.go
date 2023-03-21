package str

import "unicode"

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
