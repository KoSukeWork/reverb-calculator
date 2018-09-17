package main

import (
	"fmt"
	"flag"
    "calculator"
    "structs"
)

const defaultBpm = 120
const defaultBars = 4
const defaultResolution =  11

// Params
var bpm float64
var bars int
var resolution int

func main() {
    var stepData structs.StepDataList
    initFlags()
    stepData = calculator.GetStepDataByBpm(bpm, bars, resolution)
    printList(stepData)
}

func initFlags() {
    flag.Float64Var(&bpm, "bpm", defaultBpm, "Beats per Minute")
    flag.IntVar(&resolution, "resolution", defaultResolution, "The lowest divider of one note to show.")
    flag.IntVar(&bars, "bars", defaultBars, "The amount of bars to display")
    flag.Parse()
}

func printList(stepData structs.StepDataList) {
    fmt.Printf("%+7v %10.2f %v\n", "Tempo:", bpm, "BPM")
    for _, step := range stepData {
        fmt.Printf("%+7v %10.2f ms %10.3f s \n", step.BarPosition, step.StepValueMs, step.StepValueSec)
    }
}
