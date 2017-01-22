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

const SAMPLE_OUTPUT = `Case #1: 107
Case #2: 20
Case #3: 10
Case #4: 18
Case #5: 79`

const SAMPLE_INPUT = `5
3 2
1 1
100 100
10000 10000
5 1
1
2
3
4
5
5 5
1 2 3 4 5
2 3 4 5 1
3 4 5 1 2
4 5 1 2 3
5 1 2 3 4
5 5
1 1 1 1 1
2 2 2 2 2
3 3 3 3 3
4 4 4 4 4
5 5 5 5 5
10 4
7 15 12 6
15 3 19 18
10 9 10 14
12 14 8 8
5 3 5 11
9 14 19 11
12 6 20 9
18 13 12 15
14 14 10 20
11 19 12 11`
