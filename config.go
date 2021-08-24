package videostream

import (
	"flag"
	"fmt"
	"sort"

	"github.com/blackjack/webcam"
)

const (
	V4L2_PIX_FMT_PJPG = 0x47504A50
	V4L2_PIX_FMT_YUYV = 0x56595559
)

var supportedFormats = map[webcam.PixelFormat]bool{
	V4L2_PIX_FMT_PJPG: true,
	V4L2_PIX_FMT_YUYV: true,
}

type formats map[webcam.PixelFormat]string

// FrameSizes provides a sortable list of framesizes
type FrameSizes []webcam.FrameSize

// Len provides the lenth of a framesizes array
func (slice FrameSizes) Len() int {
	return len(slice)
}

// Less provides a way to compare two frame sizes
func (slice FrameSizes) Less(i, j int) bool {
	ls := slice[i].MaxWidth * slice[i].MaxHeight
	rs := slice[j].MaxWidth * slice[j].MaxHeight
	return ls < rs
}

// Swap provides a way to reorder frames in Framesizes
func (slice FrameSizes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SupportedFormats(cam *webcam.Webcam) formats {
	return cam.GetSupportedFormats()
}

func SupportedFrameSizes(cam *webcam.Webcam, pf webcam.PixelFormat) []webcam.FrameSize {
	frames := FrameSizes(cam.GetSupportedFrameSizes(pf))
	sort.Sort(frames)
	return frames
}

func (f formats) PrintAll() {
	for pf, st := range f {
		fmt.Printf("Name: %s, Linux Identifer: %d\n", st, pf)
		supported := false
		if supportedFormats[pf] {
			supported = true
		}
		fmt.Printf("\tIs supported format: %t\n", supported)
	}
}

func (fs FrameSizes) PrintAll() {
	for _, fr := range fs {
		fmt.Printf("\tFrame Size: %s\n", fr.GetString())
	}
}

type VideoConfig struct {
	Device       string
	VideoFormat  string
	FrameSize    string
	DisplayVideo bool
	Port         int
	PrintFPS     bool
}

func GetConfig() VideoConfig {
	dev := flag.String("d", "/dev/video0", "video device to use")
	fmtstr := flag.String("f", "", "video format to use, default first supported")
	szstr := flag.String("s", "", "frame size to use, default largest one")
	single := flag.Bool("m", false, "single image http mode, default jpeg video")
	addr := flag.Int("l", 8080, "addr to listen")
	fps := flag.Bool("p", false, "print fps info")
	flag.Parse()

	return VideoConfig{
		Device:       *dev,
		VideoFormat:  *fmtstr,
		FrameSize:    *szstr,
		DisplayVideo: *single,
		Port:         *addr,
		PrintFPS:     *fps,
	}
}
