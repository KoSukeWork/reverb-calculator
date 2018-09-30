package controller

import (
	"calculator"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/view/templates/index.html")

	c := calculator.NewCalculator()

	if r.Method == http.MethodPost {
		r.ParseForm()

		bpm, _ := strconv.ParseFloat(r.Form.Get("bpm"), 64)
		bars, _ := strconv.Atoi(r.Form.Get("bars"))
		resolution, _ := strconv.Atoi(r.Form.Get("resolution"))

		c.SetBpm(bpm)
		c.SetBars(bars)
		c.SetResolution(resolution)
	}

	c.Calculate()

	err := t.Execute(w, IndexPageData{
		Bpm:        c.Bpm(),
		Bars:       c.Bars(),
		Resolution: c.Resolution(),
		Errors:     c.Errors(),
		Data:       c.Data(),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
