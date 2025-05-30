package main

import "github.com/bitcodr/gompeg"

// HLSExample demonstrates HLS segmenting with gompeg.
func HLSExample() {
	err := gompeg.HLS().Input("input.mp4").SegmentTime(4).Output("stream.m3u8").Run()
	if err != nil {
		panic(err)
	}
}
