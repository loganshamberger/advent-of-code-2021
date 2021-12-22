package main

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

var autoCompleteMapping = map[string]int{
	"(" : 1,
	"[" : 2,
	"<" : 4,
	"{" : 3,
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


func AutoCompletionScore(chunk string) (res int) {
	input := strings.Split(chunk, "")
	var stack []string
	
	for i := 0; i<len(input); i++{
		if(strings.Contains(opening, input[i])){
			stack = append(stack, input[i])
		} else {
			fmt.Println(stack[len(stack)-1])
			if(stack[len(stack)-1] == mapping[input[i]]) {
				stack = stack[:len(stack)-1]
			}
		}
	}
	
	score := 0
	fmt.Println(stack)
	for i:=len(stack)-1;i>=0; i-- {
		score = 5*score
		score += autoCompleteMapping[stack[i]]
		fmt.Println(score)
	}

	return score;
}

