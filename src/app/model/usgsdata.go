package model

import (
	"encoding/json"
	"fmt"
	"sync"
)

var earthquakeUsgsData usgsData

// usgsData holds data retrieved from the USGS API and
// uses mutex to synchronize access to it between goroutines
// to prevent data corruption
type usgsData struct {
	details earthquakeDetails
	mux     sync.Mutex
}

// earthquakeDetails contain details about the earthquakes
type earthquakeDetails struct {
	Features []locationFeature
}

// locationFeature describes the structure of the data
// as it corresponds to data format retrieved from the
// USGS API describing locations of earthquakes
type locationFeature struct {
	Properties struct {
		Mag   float32
		Place string
		Time  int64
		URL   string
		Sig   int
	}
	Geometry struct {
		Coordinates []float64
	}
	ID string
}

// InitUSGSData takes the JSON data describing the earthquakes
// that was fetched from the USGS site API and stores it into the
// local "storage"
func InitUSGSData(jsonData []byte) error {
	// allow only one goroutine at a time to change USGS data variable(storage)
	earthquakeUsgsData.mux.Lock()
	defer earthquakeUsgsData.mux.Unlock()

	// parse JSON data received from USGS site
	err := json.Unmarshal(jsonData, &earthquakeUsgsData.details)
	if err != nil {
		return fmt.Errorf("Failed parsing fetched USGS data: %v", err)
	}

	return nil
}

func earthquakeData() earthquakeDetails {
	// allow only one goroutine at a time to access USGS data variable(storage)
	earthquakeUsgsData.mux.Lock()
	defer earthquakeUsgsData.mux.Unlock()

	return earthquakeUsgsData.details
}
