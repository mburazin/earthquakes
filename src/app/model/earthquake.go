package model

import (
	"time"
)

// Earthquake structure defines all the data that describes the
// earthquake. All the fields and their values are also used to
// send earthquake info back to the client in JSON format. So, the
// mapping to JSON key names is also specified.
type Earthquake struct {
	ID           string    `json:"id"`
	Place        string    `json:"place"`
	Time         time.Time `json:"time"`
	Magnitude    float32   `json:"mag"`
	Significance int       `json:"sig"`
	MoreInfoURL  string    `json:"infoUrl"`
	Longitude    float64   `json:"lng"`
	Latitude     float64   `json:"lat"`
	Depth        float64   `json:"depth"`
}

// GetEarthquakes returns a slice of all the earthquakes currently stored
// that were the most recently fetched from the USGS site
func GetEarthquakes() []Earthquake {
	earthquakeDetails := earthquakeData()
	earthquakes := []Earthquake{}

	for _, locFeature := range earthquakeDetails.Features {
		earthquake := featureToEarthquakeFormat(locFeature)
		earthquakes = append(earthquakes, *earthquake)
	}

	return earthquakes
}

// GetEarthquake returns a specific earthquake (identified by id)
// among the most recently fetched from the USGS site
func GetEarthquake(id string) *Earthquake {
	earthquakeDetails := earthquakeData()

	for _, locFeature := range earthquakeDetails.Features {
		if locFeature.ID == id {
			return featureToEarthquakeFormat(locFeature)
		}
	}

	return nil
}

// featureToEarthquakeFormat maps data format from USGS specific
// to the Earthquake format that is used in this application and
// sent back to the client
func featureToEarthquakeFormat(l locationFeature) *Earthquake {
	return &Earthquake{
		ID:           l.ID,
		Place:        l.Properties.Place,
		Time:         time.Unix(0, l.Properties.Time*int64(time.Millisecond)),
		Magnitude:    l.Properties.Mag,
		Significance: l.Properties.Sig,
		MoreInfoURL:  l.Properties.URL,
		Longitude:    l.Geometry.Coordinates[0],
		Latitude:     l.Geometry.Coordinates[1],
		Depth:        l.Geometry.Coordinates[2],
	}
}
