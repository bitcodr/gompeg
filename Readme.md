# gompeg  

Fluent, cross-platform FFmpeg CLI wrapper for Go

[![build](https://github.com/bitcodr/gompeg/actions/workflows/ci.yml/badge.svg)](https://github.com/bitcodr/gompeg/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bitcodr/gompeg.svg)](https://pkg.go.dev/github.com/bitcodr/gompeg)

---

## Features

* **Fluent API** – chain methods to build complex FFmpeg commands without string concatenation.  
* **Common workflows** – RTMP streaming, MP4/HLS transcoding, screenshots, audio-only, custom flags.  
* **Pipe support** – use `io.Reader` / `io.Writer` for stdin/stdout streams.  
* **Cross-platform** – works on Linux, macOS, and Windows.  FFmpeg auto-detected or custom path.  
* **Production-ready** – 90 %+ test coverage, MIT license, GitHub Actions CI, semver releases.  

---

## Installation

```bash
go get github.com/bitcodr/gompeg@v1
```

FFmpeg required – install it first:
* Ubuntu: `sudo apt install ffmpeg`
* macOS: `brew install ffmpeg`
* Windows: `choco install ffmpeg`

---

## Usage

### Basic example

```go
package main
import (
    "log"
    "github.com/bitcodr/gompeg"
)

func main() {
    err := gompeg.New().
        Input("input.mp4").
        VideoCodec("libx264").
        Output("output.mp4").
        Run()
    if err != nil { log.Fatal(err) }
}
```

### Preview the command without running

```go
cmd := gompeg.New().Input("in.mp4").Output("out.mkv")
fmt.Println(cmd.String()) // prints full ffmpeg CLI
```

### With context / timeouts

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
if err := gompeg.New().
        Input("in.mp4").
        Output("out.mp4").
        RunWithContext(ctx); err != nil {
    log.Fatal(err)
}
```

### Pipe data through stdin / stdout

```go
r, w := io.Pipe()         // example reader/writer pair
// write data to w in another goroutine …
err := gompeg.New().
    PipeInput(r).                    // ffmpeg reads from stdin ("-i -")
    OutputFormat("mp3").
    AudioCodec("libmp3lame").
    PipeOutput(os.Stdout).           // ffmpeg writes MP3 to stdout ("-")
    Run()
```

### Add custom ffmpeg flags

```go
err := gompeg.New().
    Input("input.mp4").
    Extra("-vf", "transpose=1").     // raw flags
    Output("rotated.mp4").
    Run()
```

---

## Example Programs

The `examples/` directory contains real-world usage programs:

* **HLS segmenting:** `examples/hls/main.go`
* **RTMP streaming:** `examples/rtmp.go`
* **Screenshot from video:** `examples/screenshot.go`
* **Audio extraction:** `examples/audio_extract.go`

Each file contains a function you can call or copy into your own project. To run an example, copy the function into a `main()` or adapt as needed.

---

## Common recipes

1 • Transcode to MP4 (H.264 + AAC)

```go
err := gompeg.New().
    Input("input.mkv").
    VideoCodec("libx264").
    AudioCodec("aac").
    Output("output.mp4").
    Run()
```

2 • Live-stream file to RTMP

```go
err := gompeg.Stream().
    Input("video.mp4").
    Preset("veryfast").
    Output("rtmp://localhost/live/stream").
    RealTime().                      // adds -re
    Run()
```

3 • Generate HLS

```go
err := gompeg.HLS().
    Input("movie.mp4").
    SegmentTime(5).                  // set HLS segment duration (seconds)
    Output("stream.m3u8").
    Run()
```

4 • Extract audio (MP3)

```go
err := gompeg.New().
    Input("podcast.wav").
    NoVideo().
    AudioCodec("libmp3lame").
    AudioBitrate(192).
    Output("podcast.mp3").
    Run()
```

5 • Single-frame screenshot

```go
err := gompeg.New().
    Input("video.mp4").
    Seek("00:00:10").
    VFrames(1).
    OutputFormat("image2").
    Output("shot.jpg").
    Run()
```

---

## API Reference

Every builder method returns the same `*gompeg.Command`, so you can chain arbitrarily:

### Presets

- `gompeg.New()` – start a new command
* `gompeg.HLS()` – set format to HLS (`-f hls`)
* `gompeg.Stream()` – set format to FLV (for RTMP streaming)

### Chainable Methods

- `.Input(path string)` – add input file
* `.Output(path string)` – add output file
* `.Format(fmt string)` – set output format (e.g. "mp4", "hls", "flv")
* `.OutputFormat(fmt string)` – alias for `.Format`
* `.VideoCodec(codec string)` – set video codec (e.g. "libx264")
* `.AudioCodec(codec string)` – set audio codec (e.g. "aac")
* `.VideoBitrate(kbps int)` – set video bitrate (kbps)
* `.AudioBitrate(kbps int)` – set audio bitrate (kbps)
* `.Preset(preset string)` – set encoder preset (e.g. "fast")
* `.Seek(timestamp string)` – seek to timestamp (e.g. "00:00:10")
* `.VFrames(n int)` – set number of video frames to output
* `.RealTime()` – add `-re` for real-time input
* `.NoVideo()` – disable video stream
* `.NoAudio()` – disable audio stream
* `.PipeInput(r io.Reader)` – set stdin
* `.PipeOutput(w io.Writer)` – set stdout
* `.Logs(w io.Writer)` – set stderr
* `.Extra(args ...string)` – add custom ffmpeg flags
* `.SegmentTime(seconds int)` – set HLS segment duration (only for HLS)

### Execution

- `.Run()` – run the command
* `.RunWithContext(ctx)` – run with context (for cancellation/timeouts)
* `.String()` – print the full ffmpeg command

---

## Development & Contributing

```sh
go test ./...        # run unit + integration tests
go test -cover       # view coverage
```

- FFmpeg must be on `$PATH` for integration tests.
* CI runs on Linux, macOS, Windows.
* Ensure `go fmt ./...` is clean and new code has tests.

---

## License

MIT – see LICENSE.
