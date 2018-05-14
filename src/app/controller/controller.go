package controller

import (
	"earthquakes/src/app/model"
	"html/template"
	"net/http"
)

var (
	homeController         home
	earthquakesController  earthquakes
	clientPusherController *clientPusher
)

// Startup initializes all the application controllers and registers all the routes
func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	earthquakesController.earthquakeTemplate = templates["earthquake.html"]

	clientPusherController = newClientPusher()

	homeController.registerRoutes()
	earthquakesController.registerRoutes()
	clientPusherController.registerRoutes()

	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
}

// PushDataToClients is used to push any USGS data to
// "subscribed" clients (if they exist). The earthquake data is fetched
// from the model, but the model is updated with new data at a regular interval
// by a separate goroutine
func PushDataToClients() {
	earthquakes := model.GetEarthquakes()
	clientPusherController.pushDataToClients(earthquakes)
}
