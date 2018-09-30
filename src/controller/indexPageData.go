package controller

import "calculator"

type IndexPageData struct {
	Bpm        float64
	Bars       int
	Resolution int
	Errors     []string
	Data       calculator.StepDataList
}
