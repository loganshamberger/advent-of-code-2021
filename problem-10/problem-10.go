package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"sort"
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

	var inputs []string
	syntaxScore := 0
	for scanner.Scan() {
		readIn := scanner.Text()
		inputs = append(inputs,readIn)	
	}

	var autoCompleteScores []int

	for _, line := range inputs {
		legal, char := LegalChunk(line)
			if (!legal && char != "") {
				syntaxScore += int(score[char])
			} else {
				autoCompleteScores = append(autoCompleteScores, AutoCompletionScore(line))
			}
		}
	fmt.Printf("Syntax score is %d \n", syntaxScore)
	sort.Ints(autoCompleteScores)
	fmt.Println(autoCompleteScores[len(autoCompleteScores)/2])
	
		
}