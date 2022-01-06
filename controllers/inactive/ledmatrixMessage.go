package ctrlinactive

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var leds [250]color.RGBA
var coloredLeds [68]int

func setLeds() {
	coloredLeds[0] = 256
	coloredLeds[1] = 255
	coloredLeds[2] = 224
	coloredLeds[3] = 223
	coloredLeds[4] = 192
	coloredLeds[5] = 191
	coloredLeds[6] = 160
	coloredLeds[7] = 159
	coloredLeds[8] = 128
	coloredLeds[9] = 260
	coloredLeds[10] = 251
	coloredLeds[11] = 228
	coloredLeds[12] = 219
	coloredLeds[13] = 196
	coloredLeds[14] = 187
	coloredLeds[15] = 164
	coloredLeds[16] = 155
	coloredLeds[17] = 132
	coloredLeds[18] = 193
	coloredLeds[19] = 194
	coloredLeds[20] = 195
	coloredLeds[21] = 262
	coloredLeds[22] = 263
	coloredLeds[23] = 264
	coloredLeds[24] = 265
	coloredLeds[25] = 266
	coloredLeds[26] = 247
	coloredLeds[27] = 232
	coloredLeds[28] = 215
	coloredLeds[29] = 200
	coloredLeds[30] = 183
	coloredLeds[31] = 168
	coloredLeds[32] = 151
	coloredLeds[33] = 134
	coloredLeds[34] = 135
	coloredLeds[35] = 136
	coloredLeds[36] = 137
	coloredLeds[37] = 138
	coloredLeds[38] = 268
	coloredLeds[39] = 242
	coloredLeds[40] = 238
	coloredLeds[41] = 240
	coloredLeds[42] = 63
	coloredLeds[43] = 209
	coloredLeds[44] = 206
	coloredLeds[45] = 177
	coloredLeds[46] = 174
	coloredLeds[47] = 145
	coloredLeds[48] = 142
	coloredLeds[49] = 31
	coloredLeds[50] = 95
	coloredLeds[51] = 81
	coloredLeds[52] = 111
	coloredLeds[53] = 113
	coloredLeds[54] = 114
	coloredLeds[55] = 115
	coloredLeds[56] = 116
	coloredLeds[57] = 117
	coloredLeds[58] = 118
	coloredLeds[59] = 119
	coloredLeds[60] = 87
	coloredLeds[61] = 86
	coloredLeds[62] = 85
	coloredLeds[63] = 84
	coloredLeds[64] = 83
	coloredLeds[65] = 92
	coloredLeds[66] = 99
	coloredLeds[67] = 108
}

func contains(s [68]int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func LedMatrixMessage() {

	neo := machine.D2
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)
	currentPixel := 0
	setLeds()
	for {
		if currentPixel < 329 {
			currentPixel += 1
		} else {
			currentPixel = 0
		}

		for i := range leds {
			if contains(coloredLeds, i) {
				leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
			}
			// else if i == currentPixel {
			// 	// Alpha channel is not supported by WS2812 so we leave it out
			// 	leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
			// } else {
			// 	leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			// }
		}

		ws.WriteColors(leds[:])
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
