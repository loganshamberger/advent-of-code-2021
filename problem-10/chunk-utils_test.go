package main

import (
	"testing"
)

type chunkTest struct {
	input string
	expectedStatus bool
	illegalChar string
}

var chunkTests = []chunkTest {
	chunkTest{"()",true, ""},
	chunkTest{"{()()()}", true, ""},
	chunkTest{"<([{}])>",true, ""},
	chunkTest{"(((((())))))", true, ""},
	chunkTest{"(((((>", false, ">"},
	chunkTest{"{([(<{}[<>[]}>{[]{[(<()>",false,"}"},
}
func TestLegal(t *testing.T) {
	
	for _, test := range chunkTests {
		if output, char := LegalChunk(test.input); output != test.expectedStatus || char != test.illegalChar {
			t.Errorf("Output %t not equal to expected %t", output, test.expectedStatus)
		}
	}
}