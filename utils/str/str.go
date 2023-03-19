package str

import "unicode"

type Conv string

// Uppercase 首字符大写
func (s Conv) Uppercase() string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
