package controller

import (
	"earthquakes/src/app/model"
	"earthquakes/src/app/viewmodel"
	"encoding/json"
	"html/template"
	"net/http"
	"regexp"
)

type earthquakes struct {
	earthquakeTemplate *template.Template
}

func (e earthquakes) registerRoutes() {
	http.HandleFunc("/api/earthquakes", e.handleAPIEarthquakes)
	http.HandleFunc("/earthquake/", e.handleGetEarthquake)
}

// handleAPIEarthquakes is a handler for the server's REST API
func (e earthquakes) handleAPIEarthquakes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// HTTP GET request
		// send all the earthquakes data to the client using JSON format
		quakes := model.GetEarthquakes()
		enc := json.NewEncoder(w) // JSON encoder will write to the socket
		err := enc.Encode(quakes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// handleGetEarthquake handles an AJAX get request for the unordered list containing
// single earthquake information identified by an ID in the URL
func (e earthquakes) handleGetEarthquake(w http.ResponseWriter, r *http.Request) {
	// Extract a unique earthquake ID from the URL, e.g. /earthquake/ci38172392
	earthquakePattern, _ := regexp.Compile(`/earthquake/(\w.+)`)
	matches := earthquakePattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		// html for specific earthquake info
		earthquakeID := matches[1]
		earthquake := model.GetEarthquake(earthquakeID)

		// check if earthquake with given ID found
		if earthquake == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			vm := viewmodel.NewEarthquake(*earthquake)
			// render the view (HTML fragment)
			e.earthquakeTemplate.Execute(w, vm)
		}
	} else {
		// URL format to extract ID not valid
		// report as not found
		w.WriteHeader(http.StatusNotFound)
	}
}
