package gompeg

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestBuildArgs_Minimal(t *testing.T) {
	got, err := New().Input("in.mp4").Output("out.mp4").BuildArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"-i", "in.mp4", "out.mp4"}
	if len(got) != len(want) {
		t.Fatalf("args length mismatch: %v vs %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("arg %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestBuildArgs_AllOptions(t *testing.T) {
	cmd := New().
		Input("input.mp4").
		Output("output.mp4").
		VideoCodec("libx264").
		AudioCodec("aac").
		VideoBitrate(1000).
		AudioBitrate(128).
		Preset("fast").
		Seek("00:00:10").
		VFrames(5).
		RealTime().
		NoVideo().
		NoAudio().
		Extra("-vf", "scale=1280:720").
		Format("mp4")
	args, err := cmd.BuildArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	joined := strings.Join(args, " ")
	checks := []string{"-i input.mp4", "-ss 00:00:10", "-vn", "-an", "-c:v libx264", "-c:a aac", "-b:v 1000k", "-b:a 128k", "-preset fast", "-frames:v 5", "-f mp4", "-vf scale=1280:720", "output.mp4"}
	for _, c := range checks {
		if !strings.Contains(joined, c) {
			t.Errorf("missing %q in args: %v", c, args)
		}
	}
}

func TestBuildArgs_HLS_SegmentTime(t *testing.T) {
	cmd := HLS().Input("in.mp4").SegmentTime(7).Output("out.m3u8")
	args, err := cmd.BuildArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	joined := strings.Join(args, " ")
	if !strings.Contains(joined, "-f hls") {
		t.Error("missing -f hls")
	}
	if !strings.Contains(joined, "-hls_time 7") {
		t.Error("missing -hls_time 7")
	}
}

func TestBuildArgs_Stream(t *testing.T) {
	cmd := Stream().Input("in.mp4").Output("rtmp://localhost/live/stream")
	args, err := cmd.BuildArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	joined := strings.Join(args, " ")
	if !strings.Contains(joined, "-f flv") {
		t.Error("missing -f flv")
	}
}

func TestOutputFormatAlias(t *testing.T) {
	cmd := New().Input("a.wav").OutputFormat("mp3").Output("a.mp3")
	args, err := cmd.BuildArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	joined := strings.Join(args, " ")
	if !strings.Contains(joined, "-f mp3") {
		t.Error("OutputFormat did not set -f mp3")
	}
}

func TestErrorCases(t *testing.T) {
	_, err := New().Output("out.mp4").BuildArgs()
	if err == nil {
		t.Error("expected error for missing input")
	}
	_, err = New().Input("in.mp4").BuildArgs()
	if err == nil {
		t.Error("expected error for missing output")
	}
}

func TestString(t *testing.T) {
	cmd := New().Input("a").Output("b")
	str := cmd.String()
	if !strings.HasPrefix(str, "ffmpeg ") {
		t.Errorf("String() should start with 'ffmpeg ': %q", str)
	}
}

func TestPipeInputOutputLogs(t *testing.T) {
	cmd := New().Input("a").Output("b")
	var in bytes.Buffer
	var out, errOut bytes.Buffer
	cmd.PipeInput(&in).PipeOutput(&out).Logs(&errOut)
	if cmd.stdin != &in || cmd.stdout != &out || cmd.stderr != &errOut {
		t.Error("PipeInput, PipeOutput, or Logs did not set the correct fields")
	}
}

func TestRunWithContext_InvalidFFmpeg(t *testing.T) {
	cmd := New().Input("a").Output("b")
	cmd.ffmpegPath = "" // force error
	err := cmd.RunWithContext(context.Background())
	if err == nil {
		t.Error("expected error for missing ffmpeg binary")
	}
}
