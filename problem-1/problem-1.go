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
	var currentNum int
	var increases int
	currentNum = 0
	increases = 0
	for scanner.Scan() {

		readIn, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		
		if readIn > currentNum {
			increases++
		}
		currentNum=readIn
	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println(increases-1)
}

