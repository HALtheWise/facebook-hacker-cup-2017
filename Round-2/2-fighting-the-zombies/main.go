package main

import (
	"fmt"
	"io"
	"os"
)

const PROBLEM_NAME = "fighting_the_zombies"

type Zombie struct {
	x, y int
}

type KillBox struct {
	r          int
	xmin, ymin int
}

func (z *Zombie) KilledBy(k *KillBox) bool {
	if z.x >= k.xmin && z.x <= k.xmin+k.r {
		if z.y >= k.ymin && z.y <= k.ymin+k.r {
			return true
		}
	}
	return false
}

func numEitherKills(zs []Zombie, k1 *KillBox, k2 *KillBox) (kills int) {
	for _, z := range zs {
		if z.KilledBy(k1) || z.KilledBy(k2) {
			kills++
		}
	}
	return
}

func (k *KillBox) MakeFrom(z1 *Zombie, z2 *Zombie, r int) {
	k.xmin = z1.x
	k.ymin = z2.y
	k.r = r
}

// Solve a test case, returning a TestResult
func (t TestCase) Solve() (kills TestResult) {
	nz := len(t.zombies)

	// Make killboxes
	killboxes := make([]KillBox, nz*nz)
	for i := 0; i < len(killboxes); i++ {
		killboxes[i].MakeFrom(&t.zombies[i%nz], &t.zombies[i/nz], t.r)
	}

	// Find the best combination of killboxes
	for _, kb1 := range killboxes {
		for _, kb2 := range killboxes {
			numkilled := TestResult(numEitherKills(t.zombies, &kb1, &kb2))
			if numkilled > kills {
				kills = numkilled
			}
		}
	}

	return
}

// Represents all of the information contained in the input file regarding a single
// test case.
type TestCase struct {
	zombies []Zombie
	r       int
}

// Represents all of the information regarding a single test case
// that needs to be provided in the output file.
type TestResult int

// Parse a single test case from an io.Reader, including consuming the newline
// following the last line.
func ParseCase(filein io.Reader) (t TestCase, err error) {
	var numZombies int
	_, err = fmt.Fscanf(filein, "%v %v\n", &numZombies, &t.r)
	if err != nil {
		return
	}

	t.zombies = make([]Zombie, numZombies)

	for i := 0; i < numZombies; i++ {
		_, err = fmt.Fscanf(filein, "%v %v\n", &t.zombies[i].x, &t.zombies[i].y)
		if err != nil {
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
