package main

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6

import (
	ctrl "github.com/therebelrobot/tinygo-sandbox/controllers"
	utils "github.com/therebelrobot/tinygo-sandbox/utils/shared"
)

func main() {
	utils.LogSetup()
	// utils.BLE("LED Screen!")
	// logic.LedMatrixTrace()
	// ctrl.Display()
	// ctrl.PlaySong()
	// ctrl.SdCard()
	// ctrl.Halo()
	// ctrl.Mask()
	ctrl.HunterBox()
	// logic.LedMatrixMessage()
	// logic.Fan1() // goroutine
	// logic.Led2()

}
