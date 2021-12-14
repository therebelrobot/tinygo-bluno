package main

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
import (
	"github.com/therebelrobot/tinygo-bluno/logic"
)

func main() {
    go logic.Led1() // goroutine
    logic.Led2()
}
