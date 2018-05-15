package viewmodel

import (
	"app/model"
	"time"
)

// Earthquake provides the data used when rendering the
// earthquake fragment view (template)
type Earthquake struct {
	Place        string
	Time         time.Time
	Magnitude    float32
	Significance int
	MoreInfoURL  string
}

// NewEarthquake creates a new viewmodel that is used when
// rendering the earthquake fragment view (used in AJAX reply)
func NewEarthquake(earthquake model.Earthquake) Earthquake {
	return Earthquake{
		Place:        earthquake.Place,
		Time:         earthquake.Time,
		Magnitude:    earthquake.Magnitude,
		Significance: earthquake.Significance,
		MoreInfoURL:  earthquake.MoreInfoURL,
	}
}
