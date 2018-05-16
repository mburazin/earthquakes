# Earthquake tracker
This project features a web application consisting of frontend and backend parts and is used to keep track of earthquakes taken from [USGS website](https://www.usgs.gov/).
The earthquake data is downloaded by the backend part of the service from one of the USGS APIs that keeps track of earthquakes (in the last 24hrs) and is refreshed every 5 minutes. The backend performs synchronization of earthquake data every 5 minutes to stay updated and uses websockets to push this data to the client (frontend part), so the browser gets automatically refreshed with any possible new earthquakes without needing to explicitly send extra requests.

## How to use
When accessing the server through a web browser, a simple web page opens up consisting of a Google Map showing all the earthquakes in the last 24 hours and their locations denoted by markers. When any of the markers are clicked, more information is fetched from the backend and presented in the earthquake details section of the webpage.
Every ~5 mins the map gets updated with fresh earthquake data pushed from the backend service.

## How to run the web service
In order to build the code for the web app (that starts the HTTP server) you have to have _GO_ installed on your system.
Instructions:

1) Install GO and set up `GOPATH` environment variable to point to your GO workspace
2) Clone this project into `$GOPATH/src/`
3) Install project dependency that handles websockets (Gorilla websocket):<br>
    `
    $ go get github.com/gorilla/websocket
    `
4) Go to the project directory:<br>
    `
    $ cd $GOPATH/src/earthquakes
    `
5) Build the binary from source code:<br>
    `
    $ ./build.sh
    `
6) The binary is placed in the path `./bin/app`, so run it:<br>
    `
    $ ./bin/app
    `

Now the HTTP server is started. Open up your web browser and navigate to http://localhost:8000 to open up the webpage and start using the service.

### Licensing information
This project is licensed under the terms of the MIT license. For more information, see `LICENSE` file.

### Attribution

* USGS earthquake data API (24hrs)
    - https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_day.geojson
* Gorilla Websocket (for handling websockets):
    - https://github.com/gorilla/websocket
* Google Maps:
    - https://cloud.google.com/maps-platform/
* GO programming language:
-    - https://golang.org/