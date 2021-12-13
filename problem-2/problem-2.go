package main

import (
	"os"
	"log"
	"fmt"
	"bufio"	
	"strings"
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

	for scanner.Scan() {
		command := scanner.Text()
		pos = processCommand(pos, command)
	}

	fmt.Println(pos.horizontal*pos.vertical);
}

func processCommand(pos Position, input string) Position{
	command := strings.Fields(input)
	value, err := strconv.Atoi(command[1])

	if err != nil {
		log.Fatal(err)
	}
	direction := command[0]
	switch {
	case direction == "up":
		pos.vertical -= value
	case direction == "down":
		pos.vertical += value
	case direction == "forward":
		pos.horizontal += value
	}
	log.Println(pos)
	return pos
}