package capitalise

import (
	"unicode"
)

// First capitalises the first character in a string and returns it
func First(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}
