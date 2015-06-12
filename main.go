package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/rss/example.xml", ExampleHandler)
	r.HandleFunc("/rss/weather.xml", WeatherHandler)
	r.HandleFunc("/images/radar.jpg", RadarHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
