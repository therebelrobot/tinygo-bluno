package main

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6

import (
	ctrl "github.com/therebelrobot/tinygo-bluno/controllers"
	utils "github.com/therebelrobot/tinygo-bluno/utils/shared"
)

func main() {
	utils.LogSetup()
	// utils.BLE("LED Screen!")
	// logic.LedMatrixTrace()
	ctrl.Display()
	// logic.LedMatrixMessage()
	// logic.Fan1() // goroutine
	// logic.Led2()
}
