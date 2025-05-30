package gompeg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Command builds and runs an ffmpeg invocation. Use New() or preset helpers (e.g., HLS()).
// All exported methods return the same pointer for fluent chaining.
type Command struct {
	ffmpegPath string

	// basic config
	inputs  []string
	outputs []string

	// flags
	videoCodec   string
	audioCodec   string
	videoBitrate string
	audioBitrate string
	preset       string
	format       string

	// mux
	seek     string
	vframes  int
	realTime bool
	noVideo  bool
	noAudio  bool

	// hls/streaming
	hlsSegmentTime int

	// piping + logs
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer

	// extra raw args
	extra []string
}

// New returns a fresh Command and checks ffmpeg availability.
func New() *Command {
	path, _ := exec.LookPath("ffmpeg") // ignore error; handled on Run()
	return &Command{ffmpegPath: path}
}

// HLS returns a Command preset for HLS output (format = "hls").
func HLS() *Command { return New().Format("hls") }

// Stream returns a Command preset for RTMP/FLV streaming (format = "flv").
func Stream() *Command { return New().Format("flv") }

// Input adds an input file or stream.
func (c *Command) Input(path string) *Command { c.inputs = append(c.inputs, path); return c }

// Output adds an output file or stream.
func (c *Command) Output(path string) *Command { c.outputs = append(c.outputs, path); return c }

// Format sets the output format (e.g., "mp4", "hls", "flv").
func (c *Command) Format(f string) *Command { c.format = f; return c }

// OutputFormat is an alias for Format for API compatibility.
func (c *Command) OutputFormat(f string) *Command { return c.Format(f) }

// VideoCodec sets the video codec (e.g., "libx264").
func (c *Command) VideoCodec(v string) *Command { c.videoCodec = v; return c }

// AudioCodec sets the audio codec (e.g., "aac").
func (c *Command) AudioCodec(a string) *Command { c.audioCodec = a; return c }

// VideoBitrate sets the video bitrate in kbps.
func (c *Command) VideoBitrate(k int) *Command { c.videoBitrate = fmt.Sprintf("%dk", k); return c }

// AudioBitrate sets the audio bitrate in kbps.
func (c *Command) AudioBitrate(k int) *Command { c.audioBitrate = fmt.Sprintf("%dk", k); return c }

// Preset sets the encoder preset (e.g., "fast").
func (c *Command) Preset(p string) *Command { c.preset = p; return c }

// Seek sets the seek timestamp (e.g., "00:00:10").
func (c *Command) Seek(ts string) *Command { c.seek = ts; return c }

// VFrames sets the number of video frames to output.
func (c *Command) VFrames(n int) *Command { c.vframes = n; return c }

// RealTime adds the -re flag for real-time input.
func (c *Command) RealTime() *Command { c.realTime = true; return c }

// NoVideo disables video stream (-vn).
func (c *Command) NoVideo() *Command { c.noVideo = true; return c }

// NoAudio disables audio stream (-an).
func (c *Command) NoAudio() *Command { c.noAudio = true; return c }

// PipeInput sets the stdin for ffmpeg.
func (c *Command) PipeInput(r io.Reader) *Command { c.stdin = r; return c }

// PipeOutput sets the stdout for ffmpeg.
func (c *Command) PipeOutput(w io.Writer) *Command { c.stdout = w; return c }

// Logs sets the stderr for ffmpeg.
func (c *Command) Logs(w io.Writer) *Command { c.stderr = w; return c }

// Extra appends custom ffmpeg arguments.
func (c *Command) Extra(args ...string) *Command { c.extra = append(c.extra, args...); return c }

// SegmentTime sets the HLS segment duration in seconds (only for HLS format).
func (c *Command) SegmentTime(seconds int) *Command { c.hlsSegmentTime = seconds; return c }

// BuildArgs converts the struct into ffmpeg CLI arguments.
func (c *Command) BuildArgs() ([]string, error) {
	if len(c.inputs) == 0 {
		return nil, errors.New("no input specified")
	}
	if len(c.outputs) == 0 {
		return nil, errors.New("no output specified")
	}
	var args []string
	if c.realTime {
		args = append(args, "-re")
	}
	for _, in := range c.inputs {
		args = append(args, "-i", in)
	}
	if c.seek != "" {
		args = append(args, "-ss", c.seek)
	}
	if c.noVideo {
		args = append(args, "-vn")
	}
	if c.noAudio {
		args = append(args, "-an")
	}
	if c.videoCodec != "" {
		args = append(args, "-c:v", c.videoCodec)
	}
	if c.audioCodec != "" {
		args = append(args, "-c:a", c.audioCodec)
	}
	if c.videoBitrate != "" {
		args = append(args, "-b:v", c.videoBitrate)
	}
	if c.audioBitrate != "" {
		args = append(args, "-b:a", c.audioBitrate)
	}
	if c.preset != "" {
		args = append(args, "-preset", c.preset)
	}
	if c.vframes > 0 {
		args = append(args, "-frames:v", fmt.Sprint(c.vframes))
	}
	if c.format != "" {
		args = append(args, "-f", c.format)
	}
	if c.hlsSegmentTime > 0 && c.format == "hls" {
		args = append(args, "-hls_time", fmt.Sprint(c.hlsSegmentTime))
	}
	// append extras before outputs
	args = append(args, c.extra...)
	args = append(args, c.outputs...)
	return args, nil
}

// String returns the full ffmpeg command for logging or debugging.
func (c *Command) String() string {
	a, err := c.BuildArgs()
	if err != nil {
		return "<invalid command>"
	}
	return "ffmpeg " + strings.Join(a, " ")
}

// Run executes the command. Use RunWithContext for cancellation.
func (c *Command) Run() error { return c.RunWithContext(context.Background()) }

// RunWithContext executes the command with a context for cancellation/timeouts.
func (c *Command) RunWithContext(ctx context.Context) error {
	if c.ffmpegPath == "" {
		return errors.New("ffmpeg binary not found in PATH; please install or SetPath")
	}
	args, err := c.BuildArgs()
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, c.ffmpegPath, args...)
	cmd.Stdin = c.stdin
	if c.stdout != nil {
		cmd.Stdout = c.stdout
	}
	if c.stderr != nil {
		cmd.Stderr = c.stderr
	}
	return cmd.Run()
}

// SetPath globally overrides the ffmpeg binary location.
func SetPath(p string) { defaultPath = p }

var defaultPath string
