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
	expectedCompletion string
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
	autoCompleteTest{"<",">"},
	autoCompleteTest{"{[", "]}"},
	autoCompleteTest{"[({(<(())[]>[[{[]{<()<>>", "}}]])})]"},
}

 
func TestLegal(t *testing.T) {
	
	for _, test := range chunkTests {
		if output, char := LegalChunk(test.input); output != test.expectedStatus || char != test.illegalChar {
			t.Errorf("Output %t not equal to expected %t", output, test.expectedStatus)
		}
	}
}

func TestAutoComplete(t *testing.T) {

	for _, test := range autoCompleteTests {
		if output := AutoComplete(test.input); output != test.expectedCompletion {
			t.Errorf("Output %s not equal to expected %s", output, test.expectedCompletion)
		}
	}
}