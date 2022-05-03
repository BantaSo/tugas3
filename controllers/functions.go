package controllers

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var (
		file        = path.Join("view", "index.html")
		tmpl, err   = template.ParseFiles(file)
		waterstatus string
		windstatus  string
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	water := rand.Intn(80)
	wind := rand.Intn(80)
	if water <= 5 {
		waterstatus = "aman"
	} else if water > 5 && water <= 15 {
		waterstatus = "siaga"
	} else {
		waterstatus = "bahaya"
	}
	if wind <= 6 {
		windstatus = "aman"
	} else if wind > 6 && wind <= 15 {
		windstatus = "siaga"
	} else {
		windstatus = "bahaya"
	}
	massage1 := fmt.Sprintf("Ketinggian air %d meter ", water)
	massage2 := fmt.Sprintf("kecepatan angin %d per detik ", wind)
	data := map[string]interface{}{
		"water":        massage1,
		"wind":         massage2,
		"status_water": waterstatus,
		"status_wind":  windstatus,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
