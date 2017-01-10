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

const SAMPLE_OUTPUT = `Case #1: 2
Case #2: 1
Case #3: 2
Case #4: 3
Case #5: 8`

const SAMPLE_INPUT = `5
4
30
30
1
1
3
20
20
20
11
1
2
3
4
5
6
7
8
9
10
11
6
9
19
29
39
49
59
10
32
56
76
8
44
60
47
85
71
91`
