package capitalise

import (
	"unicode"
)

// Capitalises the first character of a string
func First(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}
