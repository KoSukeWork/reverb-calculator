package main
import (
	"fmt"
	"math"
)

const maxDepth =  11
const upperBar = 4
const defaultBpm = 120

func main() {
	bpm := float64(defaultBpm)
	quarter := float64(60 / bpm) 
	whole := float64(4 * quarter)
	wholeMs := whole * 1000	
	printNotes(wholeMs, maxDepth)
}

func Round(x, unit float64) float64 { 
    return math.Round(x/unit) * unit 
} 

func printNotes(whole float64, depth int) {
	subBar := 1
	stepValue := whole
	currentBar := upperBar
	
	for step := 1; step < (depth + upperBar ); step++ {
		if currentBar <= 1 {
			currentBar = 1
			divider := math.Pow(2, float64(step - upperBar))
			stepValue = whole / divider
			subBar = int(divider)
			
		} else {
			stepValue = whole * float64(currentBar)
		}
		
		barPointer := fmt.Sprintf("%v/%v:",currentBar, subBar)
		stepValueMs := Round(stepValue, 0.01)
		stepValueSec := Round((stepValueMs / 1000), 0.001)
		fmt.Printf("%+7v %10.2f ms %10.3f s\n", barPointer, stepValueMs, stepValueSec)
		currentBar = upperBar - step 
	}
}

