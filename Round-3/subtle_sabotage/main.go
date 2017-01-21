package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const PROBLEM_NAME = "subtle_sabotage"

// Represents all of the information contained in the input file regarding a single
// test case.
type TestCase struct {
	M, N, K int
}

// Represents all of the information regarding a single test case
// that needs to be provided in the output file.
type TestResult int

// Parse a single test case from an io.Reader, including consuming the newline
// following the last line.
func ParseCase(filein *bufio.Reader) (t TestCase, err error) {
	_, err = fmt.Fscanf(filein, "%v %v %v\n", &t.N, &t.M, &t.K)
	if err != nil {
		return
	}

	return
}

func littleSolve(M, N, K int) (opts []int) {
	if 2*K+3 <= M {
		// Top and bottom touch case
		opts = append(opts, int(math.Ceil(float64(N)/float64(K))))
	}

	if (2*K+1 <= N) && (2*K+3 <= M) {
		sol := 4
		if K == 1 {
			sol = 5
		}
		opts = append(opts, sol)
	}
	return

}

// Solve a test case, returning a TestResult
func (t TestCase) Solve() TestResult {
	sols := littleSolve(t.M, t.N, t.K)

	for _, val := range littleSolve(t.N, t.M, t.K) {
		sols = append(sols, val)
	}

	if len(sols) == 0 {
		return -1
	}

	min := math.MaxInt64
	for _, val := range sols {
		if val < min {
			min = val
		}
	}
	return TestResult(min)
}

func OutputResult(testID int, result TestResult) string {
	return fmt.Sprintf("Case #%d: %v\n", testID+1, result)
}

func ParseCases(filein *bufio.Reader) (ts []TestCase, err error) {
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
	bufin := bufio.NewReader(filein)

	if err != nil {
		fmt.Printf("Did you mean to run go test? \n\t%v\n", err.Error())
		return
	}

	cases, _ := ParseCases(bufin)

	fileout, _ := os.Create(PROBLEM_NAME + "_output.txt")
	defer fileout.Close()

	for i := 0; i < len(cases); i++ {
		sol := cases[i].Solve()

		str := OutputResult(i, sol)
		fmt.Fprint(fileout, str)
		fmt.Print(str)
	}
}
