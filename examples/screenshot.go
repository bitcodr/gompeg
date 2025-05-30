package main

import (
	"log"

	"github.com/bitcodr/gompeg"
)

// ScreenshotExample demonstrates taking a screenshot from a video with gompeg.
func ScreenshotExample() {
	err := gompeg.New().
		Input("video.mp4").
		Seek("00:00:10").
		VFrames(1).
		OutputFormat("image2").
		Output("shot.jpg").
		Run()
	if err != nil {
		log.Fatal(err)
	}
}
