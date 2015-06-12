package main

import (
	"log"
	"time"

	"github.com/gorilla/feeds"
)

func encodeFeed(feed *feeds.Feed, useAtom bool) (string, error) {
	var encoded string
	var err error

	if useAtom {
		encoded, err = feed.ToAtom()
	} else {
		encoded, err = feed.ToRss()
	}
	if err != nil {
		log.Print("Error:", err)
		return "", err
	}
	return encoded, nil
}

func WeatherFeed(useAtom bool) (string, error) {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Daves Weather",
		Link:        &feeds.Link{Href: "http://scriptonomicon.org"},
		Description: "Weather Links relevant to Dave!",
		Author:      &feeds.Author{Name: "David Frascone", Email: "dave@frascone.com"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Detailed Radar Image",
			Link:        &feeds.Link{Href: "http://worker.frascone.net/images/radar.jpg"},
			Description: "Northern Utah Weather Radar (combined)",
			Author:      &feeds.Author{Name: "David Frascone", Email: "dave@frascone.com"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Radar Image",
			Link:        &feeds.Link{Href: "http://radar.weather.gov/lite/N0R/MTX_0.png"},
			Description: "Northern Utah Weather Radar",
			Author:      &feeds.Author{Name: "David Frascone", Email: "dave@frascone.com"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Traffic",
			Link:        &feeds.Link{Href: "http://dev.virtualearth.net/REST/V1/Imagery/Map/Road/40.568550%2C-111.890546/10?mapSize=640,400&mapLayer=TrafficFlow&format=png&key=Ai_LkqQXw6AkNW30JlZggReAiw4jgWlfeTuvIN7WfviMCTAVx0t3XljxeV4sxTpO"},
			Description: "Commute Traffic",
			Author:      &feeds.Author{Name: "David Frascone", Email: "dave@frascone.com"},
			Created:     now,
		},
	}

	return encodeFeed(feed, useAtom)
}

func ExampleFeed(useAtom bool) (string, error) {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "jmoiron.net blog",
		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&feeds.Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}

	return encodeFeed(feed, useAtom)
}
