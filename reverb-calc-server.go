package main

import (
	"net/http"
	"html/template"
	"log"
    "calculator"
    "fmt"
    "strconv"
    "structs"
)

const defaultBpm = 120
const defaultBars = 4
const defaultResolution =  10

const maxBpm = 300
const maxBars = 30
const maxResolution =  50

// Params
type TemplateData struct {
    Bpm float64
    Bars int
    Resolution int
    Errors []string
    StepData structs.StepDataList
}

func main() {
	http.HandleFunc("/", indexController) // set router
    http.HandleFunc("/reverb-calc", indexController) // set router
    http.Handle("/reverb-calc/css/", http.StripPrefix("/reverb-calc/css/", http.FileServer(http.Dir("./assets/css"))))
    http.Handle("/reverb-calc/js/", http.StripPrefix("/reverb-calc/js/", http.FileServer(http.Dir("./assets/js"))))


    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func indexController(w http.ResponseWriter, r *http.Request) {
    var error string

    t, _ := template.ParseFiles("src/view/index.html")

    // set defaults
    PageData := TemplateData {
        Bpm: defaultBpm,
        Bars: defaultBars,
        Resolution: defaultResolution,
    }

    // get form data
    if r.Method == http.MethodPost {
        r.ParseForm()

        // validate form data
        bpm, err := strconv.ParseFloat(r.Form.Get("bpm"), 64)
        if err != nil || bpm < 0.1 || bpm > maxBpm {
            error = fmt.Sprintf("bpm must be between %v and %v", 0.1, maxBpm)
            PageData.Errors = append(PageData.Errors, error)
        } else {
            PageData.Bpm = bpm
        }

        bars, err := strconv.Atoi(r.Form.Get("bars"))
        if err != nil || bars < 1 || bars > maxBars {
            error = fmt.Sprintf("bars must be between %v and %v", 1, maxBars)
            PageData.Errors = append(PageData.Errors, error)
        } else {
            PageData.Bars = bars
        }

        resolution, err := strconv.Atoi(r.Form.Get("resolution"))
        if err != nil || resolution < 1 || resolution > maxResolution {
            error = fmt.Sprintf("resolution must be between %v and %v", 1, maxResolution)
            PageData.Errors = append(PageData.Errors, error)
        } else {
            PageData.Resolution = resolution
        }

    }

    PageData.StepData = calculator.GetStepDataByBpm(PageData.Bpm, PageData.Bars, PageData.Resolution)
    t.Execute(w, PageData)
}

