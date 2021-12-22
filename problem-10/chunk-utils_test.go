package main

import (
	"testing"
)

type chunkTest struct {
	input string
	expectedStatus bool
	illegalChar string
}

type autoCompleteTest struct {
	input string
	expectedCompletionScore int
}
var chunkTests = []chunkTest {
	chunkTest{"()",true, ""},
	chunkTest{"{()()()}", true, ""},
	chunkTest{"<([{}])>",true, ""},
	chunkTest{"(((((())))))", true, ""},
	chunkTest{"(((((>", false, ">"},
	chunkTest{"{([(<{}[<>[]}>{[]{[(<()>",false,"}"},
}

var autoCompleteTests = []autoCompleteTest{
	autoCompleteTest{"<",4},
	autoCompleteTest{"<[",14},
	autoCompleteTest{"<[]",4},
	autoCompleteTest{"[({(<(())[]>[[{[]{<()<>>",288957},
}

 
func TestLegal(t *testing.T) {
	
	for _, test := range chunkTests {
		if output, char := LegalChunk(test.input); output != test.expectedStatus || char != test.illegalChar {
			t.Errorf("Output %t not equal to expected %t", output, test.expectedStatus)
		}
	}
}

func TestAutoCompleteScore(t *testing.T) {

	for _, test := range autoCompleteTests {
		if output := AutoCompletionScore(test.input); output != test.expectedCompletionScore {
			t.Errorf("Output %d not equal to expected %d", output, test.expectedCompletionScore)
		}
	}
}