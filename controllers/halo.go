package ctrl

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var haloleds [47]color.RGBA

func Halo() {

	haloPin := machine.D0
	haloPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	halo := ws2812.New(haloPin)
	currentPixel := 0

	for {
		if currentPixel < 47 {
			currentPixel += 1
		} else {
			currentPixel = 0
		}

		for i := range haloleds {
			if i == currentPixel {
				// Alpha channel is not supported by WS2812 so we leave it out
				haloleds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
			} else {
				haloleds[i] = color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa}
			}
		}

		halo.WriteColors(haloleds[:])
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
}
