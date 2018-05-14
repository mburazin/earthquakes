/**
 * @file Contains the code for handling the Google Map
 * @author Marko Burazin <marko.burazin1@gmail.com>
 */
(function() {
    var map;
    var markers = [];
    var infoWindow;
  
    /**
     * Initialize the Google Map used to show earthquake locations
     * @function setPosition
     */
    function init() {
      // create a new map
      map = new google.maps.Map(document.getElementById('map'), {
        minZoom: 2,
        gestureHandling: 'greedy'
      });

      // Perform a HTTP GET request to the API to get earthquake data
      $.get("/api/earthquakes", function(data, status) {
        arr = $.parseJSON(data);
        _refreshEarthquakes(arr);

        // create a websocket to the server to update the map with new earthquakes
        var ws = new WebSocket('ws://' + window.location.host + '/ws');

        // wait for new earthquake data on the websocket
        ws.addEventListener('message', function(e) {
            var arr = $.parseJSON(e.data);
            while(markers.length) { markers.pop().setMap(null); }
            _refreshEarthquakes(arr);
        });
        ws.addEventListener('error', function(e) {
            console.log("Error on a websocket: " + e.message);
        });
      });
  
    }

    /**
     * Add all earthquakes to the map as markers
     * @function _refreshEarthquakes
     * @param {Array} earthquakes contains all earthquakes
     */
    function _refreshEarthquakes(earthquakes) {
      var bounds = new google.maps.LatLngBounds();

      earthquakes.forEach(function(quake) {
        var marker = new google.maps.Marker({
          position: {lat: quake["lat"], lng: quake["lng"]},
          title: quake["place"],
          animation: google.maps.Animation.DROP
        });
    

        marker.id = quake.ID;
        marker.addListener('click', function() {
          $.get("/earthquake/" + marker.id, function(data, status) {
            $('#quake-details-list').replaceWith(data);
          });
        });

        markers.push(marker);
        bounds.extend(marker.position);
        map.fitBounds(bounds);
        marker.setMap(map);
      });
    }

    // Exported public functions
    window.GoogleMap = {
        init: init
    };
})();
