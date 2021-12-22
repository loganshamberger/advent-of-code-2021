package main

import (
	"testing"
	"reflect"
)

type simulationTest struct {
	inputList []int
	days int
	expectedPopulation int
}

var simulationTests = []simulationTest {
	simulationTest{[]int{3,4,3,1,2},1, 5},
	simulationTest{[]int{3,4,3,1,2},2, 6},
	simulationTest{[]int{3,4,3,1,2},18,26},
}

func TestSimulation(t *testing.T) {
	
	for _, test := range simulationTests {
		if output := LanternfishSimulation(test.inputList, test.days); !reflect.DeepEqual(output, test.expectedPopulation) {
			t.Errorf("Output %v not equal to expected %v", output, test.expectedPopulation)
		}
	}
}

func Equal(a, b []int) bool {
	if (len(a) != len(b)){
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}