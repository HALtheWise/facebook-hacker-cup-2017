package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PROBLEM_NAME = "fighting_the_zombie"

type DiceRoll struct {
	numDice, diceSize int
	extra             int
}

type MemoKey struct {
	d      DiceRoll
	target int
}

var diceMemoTable map[MemoKey]float64

func init() {
	diceMemoTable = make(map[MemoKey]float64)
}

func (d *DiceRoll) probability(target int) float64 {
	memokey := MemoKey{d: *d,
		target: target}
	if prob, ok := diceMemoTable[memokey]; ok {
		return prob
	}

	if d.extra != 0 {
		d2 := *d
		d2.extra = 0
		return d2.probability(target - d.extra)
	}

	if target <= 0 {
		return 1
	} else {
		if d.numDice == 0 {
			return 0
		}
	}

	if target > d.diceSize*d.numDice {
		return 0
	}

	var sum float64 = 0
	for roll := 1; roll <= d.diceSize; roll++ {
		d2 := *d
		d2.numDice -= 1
		prob := d2.probability(target - roll)
		//		fmt.Printf("Rolled %#v %d got %f\n", d2, target-roll, prob)
		sum += (1.0 / float64(d.diceSize)) * prob
	}
	diceMemoTable[memokey] = sum
	return sum
}

func (d *DiceRoll) parseFromString(s string) {
	fmt.Sscanf(s, "%dd%d", &d.numDice, &d.diceSize)
	if strings.Contains(s, "+") {
		fmt.Sscanf(strings.Split(s, "+")[1], "%d", &d.extra)
	} else if strings.Contains(s, "-") {
		fmt.Sscanf(strings.Split(s, "-")[1], "%d", &d.extra)
		d.extra *= -1
	}
}

// Represents all of the information contained in the input file regarding a single
// test case.
type TestCase struct {
	zombie int
	spells []DiceRoll
}

// Represents all of the information regarding a single test case
// that needs to be provided in the output file.
type TestResult float64

// Parse a single test case from an io.Reader, including consuming the newline
// following the last line.
func ParseCase(filein *bufio.Reader) (t TestCase, err error) {
	var numSpells int
	_, err = fmt.Fscanf(filein, "%d %d\n", &t.zombie, &numSpells)
	if err != nil {
		return
	}

	t.spells = make([]DiceRoll, numSpells)

	spellString, _ := bufio.NewReader(filein).ReadString('\n')

	spellStrings := strings.Split(spellString, " ")

	for i, str := range spellStrings {
		t.spells[i].parseFromString(str)
	}
	return
}

// Solve a test case, returning a TestResult
func (t TestCase) Solve() TestResult {
	maxProb := 0.0
	for _, spell := range t.spells {
		prob := spell.probability(t.zombie)
		if prob > maxProb {
			maxProb = prob
		}
	}
	return TestResult(maxProb)
}

func OutputResult(testID int, result TestResult) string {
	return fmt.Sprintf("Case #%d: %.8f\n", testID+1, result)
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
	if err != nil {
		fmt.Printf("Did you mean to run go test? \n\t%v\n", err.Error())
		return
	}

	cases, _ := ParseCases(bufio.NewReader(filein))

	fileout, _ := os.Create(PROBLEM_NAME + "_output.txt")
	defer fileout.Close()

	for i := 0; i < len(cases); i++ {
		sol := cases[i].Solve()

		str := OutputResult(i, sol)
		fmt.Fprint(fileout, str)
		fmt.Print(str)
	}
}
