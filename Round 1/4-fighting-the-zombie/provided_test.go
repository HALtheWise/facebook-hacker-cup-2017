// main_test.go
package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSampleCases(t *testing.T) {
	input := bufio.NewReader(strings.NewReader(SAMPLE_INPUT))

	cases, err := ParseCases(input)
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
				i+1, cases[i], str, strings.SplitN(remainingOutput, "\n", 2)[0])

			// Attempt to recover to the next test case (assumes that one case == one line
			if strings.Contains(remainingOutput, "\n") {
				remainingOutput = strings.SplitN(remainingOutput, "\n", 2)[1]
			}
		}
	}
}

const SAMPLE_OUTPUT = `Case #1: 1.000000
Case #2: 0.998520
Case #3: 0.250000
Case #4: 0.002500
Case #5: 0.400000`

const SAMPLE_INPUT = `5
2 2
2d4 1d8
10 2
10d6-10 1d6+1
8 3
1d4+4 2d4 3d4-4
40 3
10d4 5d8 2d20
10 4
1d10 1d10+1 1d10+2 1d10+3`
