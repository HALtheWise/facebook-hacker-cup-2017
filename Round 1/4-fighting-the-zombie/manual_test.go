// main_test.go
package main

import "testing"

var manualtests = []struct {
	in  TestCase
	out TestResult
}{
//	{
//		TestCase{
//			numBoxes: 4,
//			weights:  []int{1, 2, 3, 4}},
//		0},
}

func TestManualCases(t *testing.T) {

	for i, test := range manualtests {
		sol := test.in.Solve()

		if sol != test.out {
			// Test case fails!
			t.Errorf("Manual test #%d failed. \n\tProblem was %#v. \n\tReturned %#v but expected %#v",
				i+1, test.in, sol, test.out)
		}
	}
}
