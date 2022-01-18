package rgbastrip

import (
	"machine"
)

func Init() (red machine.Pin, blue machine.Pin, green machine.Pin, white machine.Pin) {
	red = machine.D2
	blue = machine.D3
	green = machine.D4
	white = machine.D5
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})
	white.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return red, blue, green, white
}
