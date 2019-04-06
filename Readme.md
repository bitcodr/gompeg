#gompeg

Simple And Productive Go FFMPEG Wrapper


How to use:
    
            package main
            
            import (
            	"github.com/amiraliio/gompeg"
            	"log"
            )
            
            func main() {
            	stream := new(gompeg.Media)
            	stream.SetInputPath("test.mp4")
            	stream.SetOutputPath("rtmp://localhost/live/test")
            	stream.SetNativeFramerateInput(true)
            	stream.SetVideoCodec("libx264")
            	stream.SetPreset("veryfast")
            	stream.SetVideoMaxBitrate(3000)
            	stream.SetBufferSize(6000)
            	stream.SetPixelFormat("yuv420p")
            	stream.SetKeyframeInterval(50)
            	stream.SetAudioCodec("aac")
            	stream.SetAudioBitRate(160)
            	stream.SetAudioChannels(2)
            	stream.SetAudioRate(44100)
            	stream.SetOutputFormat("flv")
            	if err := stream.Build(); err != nil {
            		log.Fatal(err.Error())
            	}
            }
            
            
 
            