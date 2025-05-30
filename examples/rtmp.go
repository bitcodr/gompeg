package main

import (
	"log"

	"github.com/bitcodr/gompeg"
)

// RTMPExample demonstrates RTMP streaming with gompeg.
func RTMPExample() {
	err := gompeg.Stream().
		Input("video.mp4").
		Preset("veryfast").
		Output("rtmp://localhost/live/stream").
		RealTime().
		Run()
	if err != nil {
		log.Fatal(err)
	}
}
