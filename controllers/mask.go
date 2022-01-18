package ctrl

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"image/color"
	"machine"
	"strconv"
	"time"

	scenes "github.com/therebelrobot/tinygo-sandbox/scenes"
	sharedutils "github.com/therebelrobot/tinygo-sandbox/utils/shared"
	"tinygo.org/x/drivers/ws2812"
)

const NUM_LEDS = 384

var maskleds [NUM_LEDS]color.RGBA

func Mask() {

	matrixPin := machine.D0
	matrixPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	matrix := ws2812.New(matrixPin)
	currentPixel := 0
	currentFrame := 0

	for {
		sharedutils.Log("currentFrame: " + strconv.FormatInt(int64(currentFrame), 10))
		if currentPixel < NUM_LEDS {
			currentPixel += 1
		} else {
			currentPixel = 0
		}

		for i := range maskleds {
			if sharedutils.ContainsInt(scenes.LED_ADDRESSES[currentFrame], i) {
				index := sharedutils.GetIndexInt(
					len(scenes.LED_ADDRESSES[currentFrame]),
					func(ii int) bool { return scenes.LED_ADDRESSES[currentFrame][ii] == i })
				maskleds[i] = scenes.LED_COLORS[currentFrame][index]
			} else if i == currentPixel {
				// Alpha channel is not supported by WS2812 so we leave it out
				maskleds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
			} else {
				maskleds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			}
		}

		matrix.WriteColors(maskleds[:])

		if currentFrame < len(scenes.LED_ADDRESSES)-1 {
			currentFrame += 1
		} else {
			currentFrame = 0
		}
		time.Sleep(time.Duration(scenes.FRAME_DELAY) * time.Millisecond)
	}
}
