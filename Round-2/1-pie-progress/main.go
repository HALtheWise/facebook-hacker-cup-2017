package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

const PROBLEM_NAME = "pie_progress"

// Represents all of the information contained in the input file regarding a single
// test case.
type TestCase struct {
	//	testID              int
	numDays, piesPerDay int
	costPerPie          [][]int
}

type Circumstance struct {
	t           *TestCase
	piesOwned   int
	daysElapsed int
}

var memoTable map[Circumstance]int

// Note: we assume prices are sorted in ascending order
func bestPossible(c Circumstance) (moneytospend int) {
	moneytospend = math.MaxInt64

	if memo, ok := memoTable[c]; ok {
		return memo
	}

	if c.piesOwned >= c.t.numDays {
		// Already bought enough pies
		//		fmt.Printf("Win on %#v\n", c)
		return 0
	}
	if c.daysElapsed > c.t.numDays {
		//Time is out! You lose!
		return math.MaxInt64
	}
	if c.daysElapsed > c.piesOwned {
		// You die of starvation!
		return math.MaxInt64
	}

	currentDay := c.t.costPerPie[c.daysElapsed]

	boughtCost := 0
	for buy := 0; buy <= c.t.piesPerDay; buy++ {
		// Try buying this many pies
		if buy > 0 {
			boughtCost += currentDay[buy-1]
		}
		newCircumstance := c
		newCircumstance.piesOwned += buy
		newCircumstance.daysElapsed++
		cost := boughtCost + bestPossible(newCircumstance) + buy*buy
		if cost < moneytospend {
			// We found a better solution!
			moneytospend = cost
		}
	}

	//	fmt.Println(c, moneytospend)

	memoTable[c] = moneytospend
	return moneytospend
}

// Solve a test case, returning a TestResult
func (t TestCase) Solve() (cost TestResult) {
	memoTable = make(map[Circumstance]int)
	c := Circumstance{
		t:           &t,
		piesOwned:   0,
		daysElapsed: 0,
	}
	return TestResult(bestPossible(c))
}

// Parse a single test case from an io.Reader, including consuming the newline
// following the last line.
func ParseCase(filein *bufio.Reader) (t TestCase, err error) {
	//	t.testID = rand.Int()
	_, err = fmt.Fscanf(filein, "%v %v\n", &t.numDays, &t.piesPerDay)
	if err != nil {
		return
	}

	t.costPerPie = make([][]int, t.numDays)

	for i := range t.costPerPie {
		t.costPerPie[i] = make([]int, t.piesPerDay)
	}

	for i := 0; i < t.numDays; i++ {
		dayString, _ := filein.ReadString('\n')
		costStrings := strings.Split(dayString, " ")
		for j := 0; j < t.piesPerDay; j++ {
			_, err = fmt.Sscanf(strings.TrimSpace(costStrings[j]), "%v", &t.costPerPie[i][j])
			if err != nil {
				return
			}
		}
		sort.Ints(t.costPerPie[i])
	}
	return
}

// Represents all of the information regarding a single test case
// that needs to be provided in the output file.
type TestResult int

func OutputResult(testID int, result TestResult) string {
	return fmt.Sprintf("Case #%d: %v\n", testID+1, result)
}

func ParseCases(file io.Reader) (ts []TestCase, err error) {
	filein := bufio.NewReader(file)
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
