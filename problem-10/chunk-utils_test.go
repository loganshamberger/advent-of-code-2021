package chunk_utils

import (
	"testing"
)

type chunkTest struct {
	input string
	expected bool
}

var chunkTests = []chunkTest {
	chunkTest{"()",true},
	chunkTest{"{()()()}", true},
	chunkTest{"<([{}])>",true},
}
func TestLegal(t *testing.T) {
	
	for _, test := range chunkTests {
		if output := LegalChunk(test.input); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}