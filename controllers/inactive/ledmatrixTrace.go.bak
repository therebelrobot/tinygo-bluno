package ctrlinactive

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var leds [128]color.RGBA

func LedMatrixTrace() {

	neo := machine.D0
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)
	currentPixel := 0

	for {
		if currentPixel < 329 {
			currentPixel += 1
		} else {
			currentPixel = 0
		}

		for i := range leds {
			if i == currentPixel {
				// Alpha channel is not supported by WS2812 so we leave it out
				leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
			} else {
				leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			}
		}

		ws.WriteColors(leds[:])
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
}
