package ctrl

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"image/color"
	"strconv"
	"time"

	display "github.com/therebelrobot/tinygo-sandbox/utils/xiaoexpansion/display"
)

func Display() {

	screen := display.Init()

	x := int16(0)
	y := int16(0)
	deltaX := int16(2)
	deltaY := int16(1)
	blink := false
	for {

		// millivolts, percent, status := utils.BatteryStatus()

		// utils.Log("millivolts" + strconv.FormatInt(int64(millivolts), 10))
		// utils.Log("percent" + strconv.FormatInt(int64(percent), 10))
		// utils.Log("status" + strconv.FormatInt(int64(status), 10))

		// pixel := display.GetPixel(x, y)
		// c := color.RGBA{255, 255, 255, 255}
		// if pixel {
		// 	c = color.RGBA{0, 0, 0, 255}
		// }
		// display.SetPixel(x, y, c)

		x += deltaX
		y += deltaY

		if x <= 0 || x >= 127 {
			deltaX = -deltaX
		}

		if y == 0 || y == 63 {
			deltaY = -deltaY
		}

		if y%10 == 0 {
			blink = !blink
		}
		display.Clear(screen)

		display.Text(
			screen,
			2,
			15,
			"hewo: "+
				strconv.FormatInt(int64(x), 10)+
				"\n"+
				"Say hello!\n"+
				"thenewcryptid")

		if blink {
			display.DottedRect(screen, 118, 126, 40, 62, color.RGBA{255, 255, 255, 255})
		} else {
			display.DottedRect(screen, 118, 126, 40, 62, color.RGBA{0, 0, 0, 255})
		}
		display.Rect(screen, 117, 127, 0, 63, color.RGBA{255, 255, 255, 255})
		display.FilledBolt(screen, 122, 31, color.RGBA{255, 255, 255, 255})
		// display.Circle(display, 110, 60, 10, color.RGBA{255, 255, 255, 255})
		// display.FilledCircle(display, 120, 20, 15, color.RGBA{255, 255, 255, 255})

		// --------
		// ----

		screen.Display()

		// var currentPin machine.Pin
		// for currentPin = 0; currentPin < 64; currentPin++ {
		// 	// currentPin = p
		// 	value := machine.ADC{currentPin}.Get()
		// 	// if value < 4000 && value > 3000 {
		// 	utils.Log(strconv.FormatInt(int64(currentPin), 10) + "\t" + strconv.FormatInt(int64(value), 10))
		// 	// }
		// }
		// utils.Log("ADC_BATTERY: " + strconv.FormatInt(int64(machine.ADC{machine.BATTERY}.Get()), 10))

		time.Sleep(1 * time.Millisecond)
	}
}
