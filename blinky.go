package main

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6

import (
	// "github.com/therebelrobot/tinygo-bluno/utils"
	"github.com/therebelrobot/tinygo-bluno/logic"
)

// go:embed scenes/*
// var scenes embed.FS

func main() {
	// utils.BLE("LED Screen!")
	logic.LedMatrix()
	// logic.Fan1() // goroutine
	// logic.Led2()
}
