package main

import (
	"testing"
	"reflect"
)

type simulationTest struct {
	inputList []int
	days int
	expectedFinalState []int
}

var simulationTests = []simulationTest {
	simulationTest{[]int{3,4,3,1,2},1,[]int{2,3,2,0,1}},
	simulationTest{[]int{3,4,3,1,2},2,[]int{1,2,1,6,0,8}},
}

func TestSimulation(t *testing.T) {
	
	for _, test := range simulationTests {
		if output := LanternfishSimulation(test.inputList, test.days); !reflect.DeepEqual(output, test.expectedFinalState) {
			t.Log( reflect.DeepEqual(output, test.expectedFinalState))
			t.Errorf("Output %v not equal to expected %v", output, test.expectedFinalState)
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