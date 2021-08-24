package main

import (
	"github.com/blackjack/webcam"
	"github.com/slynickel/videostream"
)

func main() {
	cfg := videostream.GetConfig()

	cam, err := webcam.Open(cfg.Device)
	if err != nil {
		panic(err.Error())
	}
	defer cam.Close()

	formats := videostream.SupportedFormats(cam)
	formats.PrintAll()
	frameSizes := videostream.SupportedFrameSizes(cam, formats)

}
