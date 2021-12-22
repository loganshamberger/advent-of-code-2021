package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var score = map[string]int64{
	")":3,
	"]":57,
	"}":1197,
	">": 25137,
}
func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	syntaxScore := 0
	for scanner.Scan() {
		readIn := scanner.Text()
		
		legal, char := LegalChunk(readIn)
		if (!legal && char != "") {
			syntaxScore += int(score[char])
		}
	
	}

	fmt.Printf("Syntax score is %d", syntaxScore)
	
}