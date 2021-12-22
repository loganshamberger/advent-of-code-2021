package chunk_utils

import "strings"

var opening = "({<["
var mapping = map[string]string{
	")" : "(",
	"]" : "[",
	">" : "<",
	"}" : "{",
}
func LegalChunk(chunk string) (res bool) {
	chars := []rune(chunk)
	var stack []string
	var n int
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if (strings.Contains(opening, char)) {
			stack = append(stack, char)
		} else {
			n = len(stack)-1
			if (mapping[char] == stack[n]) {
				stack = stack[:n]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}else {
		return false
	}
}