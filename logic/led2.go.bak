package logic

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6

import (
	"machine"

	"github.com/therebelrobot/tinygo-bluno/utils"
)


func Led2() {
    led := machine.D3
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.High()
        utils.Delay(250)
        led.Low()
        utils.Delay(250)
    }
}
