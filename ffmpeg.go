package gompeg

import (
	"fmt"
	"strconv"
	"strings"
)

type Media struct {
	aspect                string
	resolution            string
	videoBitRate          int
	videoBitRateTolerance int
	videoMaxBitrate       int
	videoMinBitrate       int
	videoCodec            string
	vFrames               int
	frameRate             int
	audioRate             int
	maxKeyframe           int
	minKeyframe           int
	keyframeInterval      int
	audioCodec            string
	audioBitrate          int
	audioChannels         int
	bufferSize            int
	preset                string
	quality               int
	inputPath             string
	outputPath            string
	outputFormat          string
	nativeFramerateInput  bool
	pixelFormat           string
	//threads               int
	//tune                  string
	//audioProfile          string
	//videoProfile          string
	//target                string
	//duration              string
	//durationInput         string
	//seekTime              string
	//strict                int
	//muxDelay              string
	//seekUsingTsInput      bool
	//seekTimeInput         string
	//hideBanner            bool
	//copyTs                bool
	//inputInitialOffset    string
	//rtmpLive              string
	//hlsPlaylistType       string
	//hlsListSize           int
	//hlsSegmentDuration    int
	//httpMethod            string
	//httpKeepAlive         bool
	//streamIds             map[int]string
	//filter                string
	//skipVideo             bool
	//skipAudio             bool
}

////////////////////////////////////////////////////////
/////////////////////Getters///////////////////////////
//////////////////////////////////////////////////////

//video aspect ratio like "4:3" , "16:9"
//aspect ratio describe image width and height relationship
func (m *Media) SetAspect(v string) {
	m.aspect = v
}

//set resolution ex: "100x100"
func (m *Media) SetResolution(v string) {
	m.resolution = v
}

//number of bits can processed in a unit of time
func (m *Media) SetVideoBitRate(v int) {
	m.videoBitRate = v
}

func (m *Media) SetVideoMaxBitrate(v int) {
	m.videoMaxBitrate = v
}

func (m *Media) SetVideoMinBitrate(v int) {
	m.videoMinBitrate = v
}

func (m *Media) SetVideoBitRateTolerance(v int) {
	m.videoBitRateTolerance = v
}

//for compress and decompress videos, it convert decompress video to compress video
func (m *Media) SetVideoCodec(v string) {
	m.videoCodec = v
}

//number of video frames to output
//frame are number of sequence pictures that published in a second
//ex : for example 24 FPS is 24 picture in a second
func (m *Media) SetVFrames(v int) {
	m.vFrames = v
}

func (m *Media) SetFrameRate(v int) {
	m.frameRate = v
}

func (m *Media) SetAudioRate(v int) {
	m.audioRate = v
}

func (m *Media) SetAudioBitRate(v int) {
	m.audioRate = v
}

func (m *Media) SetMaxKeyframe(v int) {
	m.maxKeyframe = v
}

func (m *Media) SetNativeFramerateInput(v bool) {
	m.nativeFramerateInput = v
}

func (m *Media) SetInputPath(v string) {
	m.inputPath = v
}

func (m *Media) SetPreset(v string) {
	m.preset = v
}

func (m *Media) SetBufferSize(v int) {
	m.bufferSize = v
}

//Videos, images and other visual media contain a value called Pixel Format.
// It describes the layout of each and every pixel in the image data of a picture or video. ...
// Usually, the value given for the Pixel Format in a file includes the bits per pixel (bpp) and the color channel or model.
func (m *Media) SetPixelFormat(v string) {
	m.pixelFormat = v
}

//The shorter the Interval, the better chance you have of video being of better quality.
// This is a very subjective subject for many reasons. These cameras use compression.
// Basically, a "Key Frame" is an entire and complete and total image,
// that is used as a reference for other frames ("images") that the camera generates.
func (m *Media) SetKeyframeInterval(v int) {
	m.keyframeInterval = v
}

func (m *Media) SetAudioCodec(v string) {
	m.audioCodec = v
}

//Sound Channel refers to the independent audio signal which is collected or playback when the sound is recording or playback in different spatial position.
// Therefore, the number of channel is the amount of sound source when the sound is recording or the relevant speaker number when it is playback.
func (m *Media) SetAudioChannels(v int) {
	m.audioChannels = v
}

func (m *Media) SetOutputFormat(v string) {
	m.outputFormat = v
}

func (m *Media) SetQuality(v int) {
	m.quality = v
}


func (m *Media) SetOutputPath(v string){
	m.outputPath = v
}

////////////////////////////////////////////////////////
/////////////////////Getters///////////////////////////
//////////////////////////////////////////////////////

func (m *Media) Aspect() []string {
	if m.resolution != "" {
		resolution := strings.Split(m.resolution, "x")
		if len(resolution) != 0 {
			width, _ := strconv.ParseFloat(resolution[0], 64)
			height, _ := strconv.ParseFloat(resolution[1], 64)
			return []string{"-aspect", fmt.Sprintf("%f", width/height)}
		}
	}
	if m.aspect != "" {
		return []string{"-aspect", m.aspect}
	}
	return nil
}

func (m *Media) VideoBitRate() []string {
	if m.videoBitRate != 0 {
		return []string{"-b:v", fmt.Sprintf("%d", m.videoBitRate)}
	}
	return nil
}

func (m *Media) VideoMaxBitRate() []string {
	if m.videoMaxBitrate != 0 {
		return []string{"-maxrate", fmt.Sprintf("%dk", m.videoMaxBitrate)}
	}
	return nil
}

func (m *Media) VideoMinBitRate() []string {
	if m.videoMinBitrate != 0 {
		return []string{"-minrate", fmt.Sprintf("%dk", m.videoMinBitrate)}
	}
	return nil
}

func (m *Media) VideoBitRateTolerance() []string {
	if m.videoBitRateTolerance != 0 {
		return []string{"-bt", fmt.Sprintf("%dk", m.videoBitRateTolerance)}
	}
	return nil
}

func (m *Media) VideoCodec() []string {
	if m.videoCodec != "" {
		return []string{"-c:v", m.videoCodec}
	}
	return nil
}

func (m *Media) VFrames() []string {
	if m.vFrames != 0 {
		return []string{"-frames:v", fmt.Sprintf("%d", m.vFrames)}
	}
	return nil
}

func (m *Media) FrameRate() []string {
	if m.frameRate != 0 {
		return []string{"-r", fmt.Sprintf("%d", m.frameRate)}
	}
	return nil
}

func (m *Media) AudioRate() []string {
	if m.audioRate != 0 {
		return []string{"-ar", fmt.Sprintf("%d", m.audioRate)}
	}
	return nil
}

func (m *Media) AudioBitrate() []string {
	if m.audioRate != 0 {
		return []string{"-b:a", fmt.Sprintf("%d", m.audioRate)}
	}
	return nil
}

func (m *Media) NativeFramerateInput() []string {
	if m.nativeFramerateInput {
		return []string{"-re"}
	}
	return nil
}

func (m *Media) InputPath() []string {
	if m.inputPath != "" {
		return []string{"-i", m.inputPath}
	}
	return nil
}

func (m *Media) Preset() []string {
	if m.preset != "" {
		return []string{"-preset", m.preset}
	}
	return nil
}

func (m *Media) BufferSize() []string {
	if m.bufferSize != 0 {
		return []string{"-bufsize", fmt.Sprintf("%dk", m.bufferSize)}
	}
	return nil
}

func (m *Media) PixelFormat() []string {
	if m.pixelFormat != "" {
		return []string{"-pix_fmt", m.pixelFormat}
	}
	return nil
}

func (m *Media) KeyFrameInterval() []string {
	if m.keyframeInterval != 0 {
		return []string{"-g", fmt.Sprintf("%d", m.keyframeInterval)}
	}
	return nil
}

func (m *Media) AudioCodec() []string {
	if m.audioCodec != "" {
		return []string{"-c:a", m.audioCodec}
	}
	return nil
}

func (m *Media) AudioChannels() []string {
	if m.audioChannels != 0 {
		return []string{"-ac", fmt.Sprintf("%d", m.audioChannels)}
	}
	return nil
}

func (m *Media) OutputFormat() []string {
	if m.outputFormat != "" {
		return []string{"-f", m.outputFormat}
	}
	return nil
}

func (m *Media) Quality() []string {
	if m.quality != 0 {
		return []string{"-crf", fmt.Sprintf("%d", m.quality)}
	}
	return nil
}

func (m *Media) OutputPath() []string{
	if m.outputPath !=""{
		return []string{m.outputPath}
	}
	return nil
}




