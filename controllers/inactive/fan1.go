package ctrlinactive

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6

import (
	"machine"

	"github.com/therebelrobot/tinygo-bluno/utils"
)


func Fan1() {

    fan := machine.D3
    fan.Configure(machine.PinConfig{Mode: machine.PinOutput})

    for {
		fan.High()
        utils.Delay(10000)
        fan.Low()
        utils.Delay(10000)
    }
}