package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

const PROBLEM_NAME = "fighting_the_zombies"

// Represents all of the information contained in the input file regarding a single
// test case.
type TestCase struct {
	numBoxes int
	weights  []int
}

// Represents all of the information regarding a single test case
// that needs to be provided in the output file.
type TestResult int

// Parse a single test case from an io.Reader, including consuming the newline
// following the last line.
func ParseCase(filein io.Reader) (t TestCase, err error) {
	_, err = fmt.Fscanf(filein, "%v\n", &t.numBoxes)
	if err != nil {
		return
	}

	t.weights = make([]int, t.numBoxes)

	for i := 0; i < t.numBoxes; i++ {
		_, err = fmt.Fscanf(filein, "%v\n", &t.weights[i])
		if err != nil {
			return
		}
	}
	return
}

// Solve a test case, returning a TestResult
func (t TestCase) Solve() (trips TestResult) {
	boxStackHeights := make([]int, t.numBoxes)

	for i := 0; i < t.numBoxes; i++ {
		boxStackHeights[i] = int(math.Ceil(50.0 / float64(t.weights[i])))
	}

	sort.Ints(boxStackHeights)

	fmt.Println(boxStackHeights)

	usedBoxes := 0

	for i := 0; i < t.numBoxes; i++ {
		thisbox := boxStackHeights[i]
		if usedBoxes+thisbox <= t.numBoxes {
			trips++
			usedBoxes += thisbox
		} else {
			return
		}
	}

	return
}

func OutputResult(testID int, result TestResult) string {
	return fmt.Sprintf("Case #%d: %v\n", testID+1, result)
}

func ParseCases(filein io.Reader) (ts []TestCase, err error) {
	var numCases int
	_, err = fmt.Fscanf(filein, "%v\n", &numCases)
	if err != nil {
		return
	}

	ts = make([]TestCase, numCases)

	for i := 0; i < numCases; i++ {
		ts[i], err = ParseCase(filein)
		if err != nil {
			return
		}
	}
	return
}

func main() {
	filein, err := os.Open(PROBLEM_NAME + ".txt")
	defer filein.Close()
	if err != nil {
		fmt.Printf("Did you mean to run go test? \n\t%v\n", err.Error())
		return
	}

	cases, _ := ParseCases(filein)

	fileout, _ := os.Create(PROBLEM_NAME + "_output.txt")
	defer fileout.Close()

	for i := 0; i < len(cases); i++ {
		sol := cases[i].Solve()

		str := OutputResult(i, sol)
		fmt.Fprint(fileout, str)
		fmt.Print(str)
	}
}
