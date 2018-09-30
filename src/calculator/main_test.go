package calculator_test

import (
	"calculator"
	"reflect"
	"testing"
)

func TestCalculator_SetBpm(t *testing.T) {
	var tests = []struct {
		param    float64
		expected float64
	}{
		{calculator.MinBpm - 1, calculator.DefaultBpm},
		{calculator.MaxBpm + 1, calculator.DefaultBpm},
		{calculator.MinBpm, calculator.MinBpm},
		{calculator.MaxBpm, calculator.MaxBpm},
	}

	c := calculator.NewCalculator()
	for _, test := range tests {
		c.SetBpm(test.param)
		actual := c.Bpm()
		if actual != test.expected {
			t.Errorf("Expected SetBpm(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestCalculator_SetBars(t *testing.T) {
	var tests = []struct {
		param    int
		expected int
	}{
		{calculator.MinBars - 1, calculator.DefaultBars},
		{calculator.MaxBars + 1, calculator.DefaultBars},
		{calculator.MinBars, calculator.MinBars},
		{calculator.MaxBars, calculator.MaxBars},
	}

	c := calculator.NewCalculator()
	for _, test := range tests {
		c.SetBars(test.param)
		actual := c.Bars()
		if actual != test.expected {
			t.Errorf("Expected SetBars(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestCalculator_SetResolution(t *testing.T) {
	var tests = []struct {
		param    int
		expected int
	}{
		{calculator.MinResolution - 1, calculator.DefaultResolution},
		{calculator.MaxResolution + 1, calculator.DefaultResolution},
		{calculator.MinResolution, calculator.MinResolution},
		{calculator.MaxResolution, calculator.MaxResolution},
	}

	c := calculator.NewCalculator()
	for _, test := range tests {
		c.SetResolution(test.param)
		actual := c.Resolution()
		if actual != test.expected {
			t.Errorf("Expected SetResolution(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestCalculator_Calculate(t *testing.T) {
	var tests = []struct {
		bpm        float64
		bars       int
		resolution int
		expected   calculator.StepDataList
	}{
		{
			calculator.DefaultBpm,
			calculator.DefaultBars,
			calculator.DefaultResolution,
			calculator.StepDataList{
				{"4/1", 8000, 8},
				{"3/1", 6000, 6},
				{"2/1", 4000, 4},
				{"1/1", 2000, 2},
				{"1/2", 1000, 1},
				{"1/4", 500, 0.5},
				{"1/8", 250, 0.25},
				{"1/16", 125, 0.125},
				{"1/32", 62.5, 0.063},
				{"1/64", 31.25, 0.031},
				{"1/128", 15.63, 0.016},
				{"1/256", 7.8100000000000005, 0.008},
				{"1/512", 3.91, 0.004},
			},
		},
	}

	c := calculator.NewCalculator()
	for _, test := range tests {
		c.SetBpm(test.bpm)
		c.SetBars(test.bars)
		c.SetResolution(test.resolution)
		c.Calculate()
		actual := c.Data()

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(
				"Expected Calculator() with: bpm: %v, bars: %v, resolution: %v to be \n%v , got \n%v",
				test.bpm, test.bars, test.resolution, test.expected, actual,
			)
		}

	}

}
