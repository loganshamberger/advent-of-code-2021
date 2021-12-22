package chunk_utils

import (
	"strings"
	"fmt"
)

var opening = "({<["
var mapping = map[string]string{
	")" : "(",
	"]" : "[",
	">" : "<",
	"}" : "{",
}


func LegalChunk(chunk string) (res bool, illegal string) {
	chars := []rune(chunk)
	var stack []string
	var n int
	fmt.Printf("Analyzing %s \n", chunk)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if (strings.Contains(opening, char)) {
			stack = append(stack, char)
		} else {
			n = len(stack)-1
			if (mapping[char] == stack[n]) {
				stack = stack[:n]
			} else {
				fmt.Printf("Chunk %s Corrupted \n", chunk)
				return false, char
			}
		}
	}
	return true, ""
}

