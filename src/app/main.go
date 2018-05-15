package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"app/controller"
	"app/model"
)

const (
	usgsDayURL            = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_day.geojson"
	usgsDataFetchInterval = 5 * time.Second
)

func main() {
	templates := loadTemplates()
	connectToUSGS()
	controller.Startup(templates)
	err := http.ListenAndServe(":8000", http.DefaultServeMux)
	if err != nil {
		log.Println("Failed to start the HTTP server:", err)
	}
}

// loadTemplates prepares all the html templates (from the ./templates subdirectory)
// and returns a map with names of templates as keys.
// Names of the templates are the same as files in 'content' and 'fragments' subdirectories
// Raises a panic if any fault loading the templates occurs.
func loadTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	contentFiles := filesInDir(basePath + "/content")
	for _, contentFile := range contentFiles {
		f, err := os.Open(basePath + "/content/" + contentFile.Name())
		if err != nil {
			panic("Failed to open template file \"" + contentFile.Name() + "\": " + err.Error())
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to load content from file \"" + contentFile.Name() + "\": " + err.Error())
		}
		f.Close()

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of \"" + contentFile.Name() + "\": " + err.Error())
		}

		result[contentFile.Name()] = tmpl
	}

	fragmentFiles := filesInDir(basePath + "/fragments")
	for _, fragmentFile := range fragmentFiles {
		fragment := template.Must(template.ParseFiles(basePath + "/fragments/" + fragmentFile.Name()))
		result[fragmentFile.Name()] = fragment
	}

	return result
}

// filesInDir fetches information about all the files in the directory given by path
func filesInDir(path string) []os.FileInfo {
	dir, err := os.Open(path)
	if err != nil {
		panic("Failure opening templates directory: " + err.Error())
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failure reading file from content directory: " + err.Error())
	}

	return files
}

// connectToUSGS fetches the USGS earthquake data from the USGS API and starts
// a goroutine which will consistently download the data at a regular time
// interval
func connectToUSGS() {
	err := fetchUSGSData()
	if err != nil {
		panic("Failure fetching initial USGS earthquake data: " + err.Error())
	}
	// start goroutine that will continually fetch data from USGS site
	go intervalUSGSDataFetcher()
}

func intervalUSGSDataFetcher() {
	timer := time.Tick(usgsDataFetchInterval)
	// loop each tick forever
	for range timer {
		err := fetchUSGSData()
		if err != nil {
			// better luck next tick
			log.Printf("Failure fetching data from USGS: %v", err)
		}

		// Push new data to all subscribed clients through a websocket
		controller.PushDataToClients()
	}
}

// fetchUSGSData performs a GET request to the USGS API to fetch the
// earthquake data for the last 24 hours and stores the data to the
// usgsdata model
func fetchUSGSData() error {
	var usgsClient = &http.Client{Timeout: 10 * time.Second}
	r, err := usgsClient.Get(usgsDayURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	jsonData, err := ioutil.ReadAll(r.Body)
	err = model.InitUSGSData(jsonData)
	if err != nil {
		return err
	}

	return nil
}
