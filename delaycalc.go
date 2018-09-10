package main

import (
	"fmt"
	"math"
	"flag"
)

const defaultdepth =  11
const defaultBars = 4
const defaultBpm = 120

var bpm float64
var depth int
var bars int

func main() {

    initFlags()
	quarter := float64(60 / bpm) 
	whole := float64(4 * quarter)
	wholeMs := whole * 1000

    printHeader(bpm)
	printNotes(wholeMs, depth, bars)
}

func initFlags() {
    flag.Float64Var(&bpm, "bpm", 120, "Beats per Minute")
    flag.IntVar(&depth, "depth", 11, "The lowest divider of one note to show.")
    flag.IntVar(&bars, "bars", 4, "The amount of bars to display")
    flag.Parse()
}

func printHeader(bpm float64) {
    fmt.Printf("%7v %10v\n", "BPM:", bpm)
    fmt.Println("==================================")
}

func printNotes(whole float64, depth int, upperBar int) {
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
		
		barPosition := fmt.Sprintf("%v/%v:",currentBar, subBar)
		stepValueMs := Round(stepValue, 0.01)
		stepValueSec := Round((stepValueMs / 1000), 0.001)
		currentBar = upperBar - step

		fmt.Printf("%+7v %10.2f ms %10.3f s \n", barPosition, stepValueMs, stepValueSec)
	}
}

func Round(x, unit float64) float64 {
    return math.Round(x/unit) * unit
}
