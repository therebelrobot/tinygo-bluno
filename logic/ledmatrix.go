package logic

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"log"
	"machine"
	"time"

	"github.com/therebelrobot/tinygo-bluno/scenes"
	"github.com/therebelrobot/tinygo-bluno/utils"

	"tinygo.org/x/drivers/ws2812"
)

// var leds [128]color.RGBA

func LedMatrix() {

	// data, err := utils.SLed("blinky.sled")

	// if err != nil {

	// 	log.Fatal(err)
	// }

	// log.Print(data)

	neo := machine.D2
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)
	total := 0
	frame := 0

	config := scenes.GetConfig()

	for {
		if frame == total {
			frame = 0
		} else {
			frame += 1
		}

		leds, totalFrames, frameDelay, err := utils.SLed(config, frame)

		if err != nil {
			log.Fatal(err)
		}
		total = totalFrames

		ws.WriteColors(leds[:])
		led.Set(frame%4 == 0)
		time.Sleep(time.Duration(frameDelay) * time.Millisecond)
	}
}
