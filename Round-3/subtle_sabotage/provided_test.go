// main_test.go
package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSampleCases(t *testing.T) {
	input := strings.NewReader(SAMPLE_INPUT)

	cases, err := ParseCases(bufio.NewReader(input))
	if err != nil {
		t.Errorf("Parsing failed with error (%v)", err.Error())
	}

	remainingOutput := SAMPLE_OUTPUT + "\n"

	for i := 0; i < len(cases); i++ {
		sol := cases[i].Solve()

		str := OutputResult(i, sol)

		if strings.HasPrefix(remainingOutput, str) {
			// Test succeeds
			remainingOutput = remainingOutput[len(str):]
		} else {
			// Test fails
			t.Errorf("Sample case #%d failed. \n\tProblem was %#v. \n\tReturned %#v but expected %#v",
				i+1, cases[i], sol, strings.SplitN(remainingOutput, "\n", 2)[0])

			// Attempt to recover to the next test case (assumes that one case == one line
			if strings.Contains(remainingOutput, "\n") {
				remainingOutput = strings.SplitN(remainingOutput, "\n", 2)[1]
			}
		}
	}
}

const SAMPLE_OUTPUT = `Case #1: -1
Case #2: 4
Case #3: 5
Case #4: -1
Case #5: 3`

const SAMPLE_INPUT = `5
2 2 1
4 5 1
6 6 1
6 6 2
773632 635271 223841`
