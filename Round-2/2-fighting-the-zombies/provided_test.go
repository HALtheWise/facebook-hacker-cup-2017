// main_test.go
package main

import (
	"strings"
	"testing"
)

func TestSampleCases(t *testing.T) {
	input := strings.NewReader(SAMPLE_INPUT)

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
				i+1, cases[i], sol, strings.SplitN(remainingOutput, "\n", 2)[0])

			// Attempt to recover to the next test case (assumes that one case == one line
			if strings.Contains(remainingOutput, "\n") {
				remainingOutput = strings.SplitN(remainingOutput, "\n", 2)[1]
			}
		}
	}
}

const SAMPLE_OUTPUT = `Case #1: 6
Case #2: 2
Case #3: 4
Case #4: 6
Case #5: 7`

const SAMPLE_INPUT = `5
7 3
1 5
2 3
2 1
5 1
6 3
4 4
4 5
4 1
0 0
0 2
10 0
10 2
4 2
0 0
0 2
10 0
10 2
7 3
8 5
3 6
9 2
4 5
3 2
1 8
2 8
7 6
8 5
3 6
9 2
4 5
3 2
1 8
2 8
`
