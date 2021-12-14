package logic

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6

import (
	"machine"

	"github.com/therebelrobot/tinygo-bluno/utils"
)

func Led1() {
    led := machine.D5
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.High()
        utils.Delay(1000)
        led.Low()
        utils.Delay(1000)
    }
}