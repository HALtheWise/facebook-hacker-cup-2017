package main

import (
	"fmt"
	"math"
	"os"
)

func handleFile(inputfile, outputfile string) {
	filein, err := os.Open(inputfile)
	defer filein.Close()
	fileout, err := os.Create(outputfile)
	defer fileout.Close()
	if err != nil {
		return
	}

	var numCases int
	fmt.Fscanf(filein, "%v\n", &numCases)
	fmt.Printf("numCases: %v\n", numCases)

	for i := 0; i < numCases; i++ {
		var percent, x, y float64
		fmt.Fscanf(filein, "%v %v %v\n", &percent, &x, &y)

		color := isColored(x, y, percent)
		fmt.Fprintf(fileout, "Case #%d: %v\n", i+1, color)
	}

}

func main() {
	handleFile("input.txt", "output.txt")
}

type MyColor bool

const (
	Black MyColor = true
	White MyColor = false
)

func (c MyColor) String() string {
	if c == Black {
		return "black"
	} else {
		return "white"
	}
}

func isColored(x, y, percent float64) (isblack MyColor) {
	x -= 50
	y -= 50
	percent /= 100

	fmt.Printf("%v %v %v\n", x, y, percent)
	// Measured clockwise from Up in fractional circles
	angleTo := math.Atan2(x, y) / (2 * math.Pi)
	if angleTo < 0 {
		angleTo += 1
	}

	distTo := math.Sqrt(x*x + y*y)

	fmt.Printf("%v\t%v\n", angleTo, percent)
	if distTo > 50 {
		return White
	}

	if angleTo > percent {
		return White
	} else {
		return Black
	}
}
