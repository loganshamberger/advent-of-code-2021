package main

import (
	"fmt"
	"strconv"
	"strings"
)

func LanternfishSimulation(inputList []int, days int) (res int) {
	var result []int
	result = inputList
	population := len(result)
	for i := 0; i < days; i++ {
		newFish := 0
		for j := 0; j < len(result); j++ {
			if result[j] == 0 {
				newFish++
				population++
				result[j] = 6
			} else {
				result[j]--
			}
		}
		for k := 0; k < newFish; k++ {
			result = append(result, 8)
		}
		fmt.Printf("There are %d fish on day %d \n", len(result), i+1)
	}
	return population
}

func LanternfishSimulationOptimized(inputList string, days int) (res int) {
	fishData := [9]int{}
	population := 0
	//Perform initialSort
	for i := 0; i < len(fishData); i++ {
		fishData[i] = strings.Count(inputList, strconv.Itoa(i))
		population += fishData[i]
	}

	fmt.Println(population)
	var newFish int

	for j := 0; j < days; j++ {
		newFish = fishData[0]
		population += newFish
		fishData[0] = fishData[1]
		fishData[1] = fishData[2]
		fishData[2] = fishData[3]
		fishData[3] = fishData[4]
		fishData[4] = fishData[5]
		fishData[5] = fishData[6]
		fishData[6] = newFish + fishData[7]
		fishData[7] = fishData[8]
		fishData[8] = newFish
		fmt.Printf("Day %d Fish Data %v Population %d \n", j+1, fishData, population)
	}

	return 0

}
