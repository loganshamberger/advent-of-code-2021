package main

import "fmt"

func LanternfishSimulation(inputList []int, days int) (res []int) {
	var result []int
	result = inputList
	for i:=0; i<days; i++{
		newFish := 0;
		for j := 0; j<len(result); j++{
			if(result[j]>0){
				result[j]--
			}else{
				newFish++
				result[j]=6
			}
		}
		
		fmt.Printf("There are %d fish on day %d \n", len(result), i+1)
	}
	return result
}

