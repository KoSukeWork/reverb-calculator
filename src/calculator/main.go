package calculator

import (
	"fmt"
    "math"
    "structs"
)

func GetStepDataByBpm(bpm float64, bars int, resolution int) structs.StepDataList {
    var stepDataList structs.StepDataList

    quarter := float64(60 / bpm)
	whole   := float64(4 * quarter)
	wholeMs := whole * 1000

	subBar := 1
	stepValue := wholeMs
	currentBar := bars

	for step := 1; step < (resolution + bars ); step++ {
		if currentBar <= 1 {
			currentBar = 1
			divider := math.Pow(2, float64(step - bars))
			stepValue = wholeMs / divider
			subBar = int(divider)
		} else {
			stepValue = wholeMs * float64(currentBar)
		}

		barPosition := fmt.Sprintf("%v/%v",currentBar, subBar)
		stepValueMs := Round(stepValue, 0.01)
		stepValueSec := Round((stepValueMs / 1000), 0.001)
		currentBar = bars - step

        stepDataList = append(stepDataList, structs.StepData{barPosition, stepValueMs, stepValueSec})
	}

    return stepDataList
}

func Round(x, unit float64) float64 {
    return math.Round(x/unit) * unit
}
