package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"strconv"
	
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []int
	currentNum := 0
	increases := 0
	for scanner.Scan() {

		readIn, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		
		if readIn > currentNum {
			increases++
		}
		currentNum=readIn
		result= append(result, readIn)
	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println(increases-1)
	fmt.Println(numberOfIncreases(slidingWindowSum(result)))
}

// Function to generate window sum
func slidingWindowSum(arr []int) []int {
	var slide []int

	for i := 0; i< len(arr)-2; i++ {
		element := arr[i]+arr[i+1]+arr[i+2]
		slide= append(slide, element)
	}
	return slide
}

func numberOfIncreases(arr []int) int {
	increase := 0
	for i := 0; i<len(arr)-1; i++ {
		if(arr[i]<arr[i+1]){
			increase++
		}
	}
	return increase
}

