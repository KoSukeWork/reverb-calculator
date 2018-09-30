package calculator

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"math"
)

type calculator struct {
	bpm        float64
	bars       int
	resolution int
	data       StepDataList
	errors     []string
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
	if govalidator.InRangeFloat64(bpm, MinBpm, MaxBpm) {
		c.bpm = bpm
	} else {
		err := fmt.Sprintf("bpm must be between %v and %v", MinBpm, MaxBpm)
		c.appendError(err)
	}
}

func (c *calculator) Bars() int {
	return c.bars
}

func (c *calculator) SetBars(bars int) {
	if govalidator.InRangeInt(bars, MinBars, MaxBars) {
		c.bars = bars
	} else {
		err := fmt.Sprintf("bars must be between %v and %v", MinBars, MaxBars)
		c.appendError(err)
	}
}

func (c *calculator) Resolution() int {
	return c.resolution
}

func (c *calculator) SetResolution(resolution int) {
	if govalidator.InRangeInt(resolution, MinResolution, MaxResolution) {
		c.resolution = resolution
	} else {
		err := fmt.Sprintf("resolution must be between %v and %v", MinResolution, MaxResolution)
		c.appendError(err)
	}
}

func (c *calculator) Data() StepDataList {
	return c.data
}

func (c *calculator) setData(data StepDataList) {
	c.data = data
}

func (c *calculator) Errors() []string {
	return c.errors
}

func (c *calculator) appendError(err string) {
	c.errors = append(c.errors, err)
}
