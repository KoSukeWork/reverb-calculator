package main

import (
	"calculator"
	"flag"
	"fmt"
)

// Params
var bpm float64
var bars int
var resolution int

func main() {
	parseFlags()

	c := calculator.NewCalculator()
	c.SetBpm(bpm)
	c.SetBars(bars)
	c.SetResolution(resolution)
	c.Calculate()
	printOutput(c.Bpm(), c.Data())
}

func parseFlags() {
	flag.Float64Var(&bpm, "bpm", calculator.DefaultBpm, "Beats per Minute")
	flag.IntVar(&resolution, "resolution", calculator.DefaultResolution, "The lowest divider of one note to show.")
	flag.IntVar(&bars, "bars", calculator.DefaultBars, "The amount of bars to display")
	flag.Parse()
}

func printOutput(bpm float64, stepData calculator.StepDataList) {
	fmt.Printf("%+7v %10.2f %v\n", "Tempo:", bpm, "BPM")
	for _, step := range stepData {
		fmt.Printf("%+7v %10.2f ms %10.3f s \n", step.BarPosition, step.StepValueMs, step.StepValueSec)
	}
}
