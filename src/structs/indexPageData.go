package structs

type IndexPageData struct {
    Bpm float64
    Bars int
    Resolution int
    Errors []string
    StepData StepDataList
}