package main

import (
	"calculator"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"structs"
)

const defaultServerPort = 9090

// Params
var serverPort int

func main() {

	initFlags()
	serverListenConfig := fmt.Sprintf(":%d", serverPort)

	http.HandleFunc("/", indexController)            // set router
	http.HandleFunc("/reverb-calc", indexController) // set router
	http.Handle("/reverb-calc/css/", http.StripPrefix("/reverb-calc/css/", http.FileServer(http.Dir("src/view/assets/css"))))
	http.Handle("/reverb-calc/js/", http.StripPrefix("/reverb-calc/js/", http.FileServer(http.Dir("src/view/assets/js"))))

	fmt.Printf("starting Server at %v\n", serverListenConfig)

	err := http.ListenAndServe(serverListenConfig, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initFlags() {
	flag.IntVar(&serverPort, "port", defaultServerPort, "The Port the server should listen.")
	flag.Parse()
}

func indexController(w http.ResponseWriter, r *http.Request) {
	var PageData structs.IndexPageData
	var pageError string

	t, _ := template.ParseFiles("src/view/templates/index.html")

	// set default
	c := calculator.NewCalculator()

	// get form data
	if r.Method == http.MethodPost {
		r.ParseForm()

		// validate form data
		bpm, err := strconv.ParseFloat(r.Form.Get("bpm"), 64)
		if err != nil || bpm < 0.1 || bpm > calculator.MaxBpm {
			pageError = fmt.Sprintf("bpm must be between %v and %v", 0.1, calculator.MaxBpm)
			PageData.Errors = append(PageData.Errors, pageError)
		} else {
			c.SetBpm(bpm)
		}

		bars, err := strconv.Atoi(r.Form.Get("bars"))
		if err != nil || bars < 1 || bars > calculator.MaxBars {
			pageError = fmt.Sprintf("bars must be between %v and %v", 1, calculator.MaxBars)
			PageData.Errors = append(PageData.Errors, pageError)
		} else {
			c.SetBars(bars)
		}

		resolution, err := strconv.Atoi(r.Form.Get("resolution"))
		if err != nil || resolution < 1 || resolution > calculator.MaxResolution {
			pageError = fmt.Sprintf("resolution must be between %v and %v", 1, calculator.MaxResolution)
			PageData.Errors = append(PageData.Errors, pageError)
		} else {
			c.SetResolution(resolution)
		}

	}

	c.Calculate()

	// Build Output
	PageData.Bpm = c.Bpm()
	PageData.Bars = c.Bars()
	PageData.Resolution = c.Resolution()
	PageData.StepData = c.Data()

	t.Execute(w, PageData)
}
