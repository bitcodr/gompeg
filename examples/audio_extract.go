package main

import (
	"log"

	"github.com/bitcodr/gompeg"
)

// AudioExtractExample demonstrates extracting audio from a video with gompeg.
func AudioExtractExample() {
	err := gompeg.New().
		Input("video.mp4").
		NoVideo().
		AudioCodec("libmp3lame").
		AudioBitrate(192).
		Output("audio.mp3").
		Run()
	if err != nil {
		log.Fatal(err)
	}
}
