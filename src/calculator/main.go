package calculator

import (
	"fmt"
	"math"
)

type calculator struct {
	bpm        float64
	bars       int
	resolution int
	data       StepDataList
}

func NewCalculator() *calculator {
	c := new(calculator)
	c.SetBpm(DefaultBpm)
	c.SetBars(DefaultBars)
	c.SetResolution(DefaultResolution)
	return c
}

func (c *calculator) Calculate() {
	bpm := c.Bpm()
	bars := c.Bars()
	resolution := c.Resolution()
	data := c.Data()

	quarter := float64(60 / bpm)
	whole := float64(4 * quarter)
	wholeMs := whole * 1000

	subBar := 1
	stepValue := wholeMs
	currentBar := c.bars

	for step := 1; step < (resolution + bars); step++ {
		if currentBar <= 1 {
			currentBar = 1
			divider := math.Pow(2, float64(step-bars))
			stepValue = wholeMs / divider
			subBar = int(divider)
		} else {
			stepValue = wholeMs * float64(currentBar)
		}

		barPosition := fmt.Sprintf("%v/%v", currentBar, subBar)
		stepValueMs := c.round(stepValue, 0.01)
		stepValueSec := c.round((stepValueMs / 1000), 0.001)
		currentBar = bars - step

		data = append(data, StepData{
			barPosition,
			stepValueMs,
			stepValueSec,
		})
	}

	c.setData(data)
}

func (c *calculator) round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Getter/Setters
func (c *calculator) Bpm() float64 {
	return c.bpm
}

func (c *calculator) SetBpm(bpm float64) {
	c.bpm = bpm
}

func (c *calculator) Bars() int {
	return c.bars
}

func (c *calculator) SetBars(bars int) {
	c.bars = bars
}

func (c *calculator) Resolution() int {
	return c.resolution
}

func (c *calculator) SetResolution(resolution int) {
	c.resolution = resolution
}

func (c *calculator) Data() StepDataList {
	return c.data
}

func (c *calculator) setData(data StepDataList) {
	c.data = data
}
