// main_test.go
package main

import "testing"

var manyZombies []Zombie

func init() {
	manyZombies = make([]Zombie, 50)
	for i := range manyZombies {
		manyZombies[i].x = i
		manyZombies[i].y = 0
	}
	manualtests[0].in.zombies = manyZombies
	manualtests[1].in.zombies = manyZombies
	manualtests[2].in.zombies = manyZombies
}

var manualtests = []struct {
	in  TestCase
	out TestResult
}{
	{
		TestCase{
			zombies: manyZombies,
			r:       1,
		},
		4},
	{
		TestCase{
			zombies: manyZombies,
			r:       3,
		},
		8},
	{
		TestCase{
			zombies: manyZombies,
			r:       1000,
		},
		50},
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
