package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/CodeMonk/CarwingsFeeds/radar"
)

func dumpRequest(req *http.Request) {
	data, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		return
	}
	log.Print("Error: Could not DumpRequestOut:", err)
	data, err = httputil.DumpRequest(req, true)
	if err != nil {
		log.Print("Error: Could not DumpRequest:", err)
		return
	}

	log.Printf("Req: %s", string(data))
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	dumpRequest(req)
	io.WriteString(w, "Hello World!\n")
}
func WeatherHandler(w http.ResponseWriter, req *http.Request) {
	dumpRequest(req)
	rss, err := WeatherFeed(false)
	if err != nil {
		log.Print("Error:", err)
		return
	}
	io.WriteString(w, rss)
}
func ExampleHandler(w http.ResponseWriter, req *http.Request) {
	dumpRequest(req)
	rss, err := ExampleFeed(false)
	if err != nil {
		log.Print("Error:", err)
		return
	}
	io.WriteString(w, rss)
}

func RadarHandler(w http.ResponseWriter, req *http.Request) {
	dumpRequest(req)
	// get our image
	r := radar.New("MTX")
	image, err := r.GetImageBlob()
	if err != nil {
		log.Print("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(image)
}
