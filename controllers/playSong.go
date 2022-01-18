package ctrl

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	utils "github.com/therebelrobot/tinygo-sandbox/utils/shared"
	xiaobutton "github.com/therebelrobot/tinygo-sandbox/utils/xiaoexpansion/button"
	xiaobuzzer "github.com/therebelrobot/tinygo-sandbox/utils/xiaoexpansion/buzzer"
)

func PlaySong() {
	// length := 28 // the number of notes
	// notes := "GGAGcB GGAGdc GGxecBA yyecdc"
	// beats := make([]int, length)
	// beats[0] = 2
	// beats[1] = 2
	// beats[2] = 8
	// beats[3] = 8
	// beats[4] = 8
	// beats[5] = 16
	// beats[6] = 1
	// beats[7] = 2
	// beats[8] = 2
	// beats[9] = 8
	// beats[10] = 8
	// beats[11] = 8
	// beats[12] = 16
	// beats[13] = 1
	// beats[14] = 2
	// beats[15] = 2
	// beats[16] = 8
	// beats[17] = 8
	// beats[18] = 8
	// beats[19] = 8
	// beats[20] = 16
	// beats[21] = 1
	// beats[22] = 2
	// beats[23] = 2
	// beats[24] = 8
	// beats[25] = 8
	// beats[26] = 8
	// beats[27] = 16
	// tempo := 150

	button := xiaobutton.Init()
	speaker, bzPin := xiaobuzzer.Init()
	for {
		// utils.Log("looping")
		isPressed := !(button.Get())

		if isPressed {
			bzPin.High()
			speaker.On()
			utils.Log("pressing ")
			// xiaobuzzer.PlaySong(buzzer, notes, beats, tempo)

		} else {
			bzPin.Low()
			speaker.Off()
			// utils.Log("not")

		}
	}

}
