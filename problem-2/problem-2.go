package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"strconv"
	
)

type Position struct {
	vertical  int
	horizontal int
}

func main(){
	pos := Position {
		vertical: 0,
		horizontal: 0,
	}

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
}
