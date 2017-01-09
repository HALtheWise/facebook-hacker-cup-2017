package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

		var numBoxes int
		fmt.Fscanf(filein, "%v\n", &numBoxes)

		var boxes []int

		for j := 0; j < numBoxes; j++ {
			var weight int
			fmt.Fscanf(filein, "%v\n", &weight)
			boxes = append(boxes, weight)
		}

		trips := TripsNeeded(boxes)
		fmt.Fprintf(fileout, "Case #%d: %v\n", i+1, trips)
	}

}

func main() {
	handleFile("input.txt", "output.txt")
}

func TripsNeeded(boxWeights []int) (trips int) {
	numBoxes := len(boxWeights)
	boxStackHeights := make([]int, numBoxes)

	for i := 0; i < numBoxes; i++ {
		boxStackHeights[i] = int(math.Ceil(50.0 / float64(boxWeights[i])))
	}

	sort.Ints(boxStackHeights)

	fmt.Println(boxStackHeights)

	usedBoxes := 0

	for i := 0; i < numBoxes; i++ {
		thisbox := boxStackHeights[i]
		if usedBoxes+thisbox <= numBoxes {
			trips++
			usedBoxes += thisbox
		} else {
			return
		}
	}

	return
}
