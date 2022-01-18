package ctrl

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"bytes"
	"image/color"
	"strconv"
	"time"

	sharedutils "github.com/therebelrobot/tinygo-sandbox/utils/shared"
	tea5767 "github.com/therebelrobot/tinygo-sandbox/utils/tea5767"

	display "github.com/therebelrobot/tinygo-sandbox/utils/xiaoexpansion/display"
)

func uniqueFloat(floatSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range floatSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func floatArrayToString(A []float64, delim string) string {

	var buffer bytes.Buffer
	for i := 0; i < len(A); i++ {
		buffer.WriteString(strconv.FormatFloat(A[i], 'f', 1, 64))
		if i != len(A)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

func HunterBox() {
	sharedutils.Log("EMF")

	screen := display.Init()

	radio := tea5767.Init()

	currentFreq := radio.Constants.FreqRangeUs[0]
	currentFrame := 0
	goingUp := true
	numberOfCycles := 1

	tea5767.SetFrequency(radio, currentFreq)

	highLvl := make(map[int64][]float64)
	totalHighLevel := 0
	var maxLevel int64
	maxLevel = 0

	for {
		isInitializing := numberOfCycles <= 4
		var finalFreqs []float64
		if _, ok := highLvl[maxLevel-2]; ok {
			totalHighLevel = len(highLvl[maxLevel]) + len(highLvl[maxLevel-1]) + len(highLvl[maxLevel-2])
		} else if _, ok := highLvl[maxLevel-1]; ok {
			totalHighLevel = len(highLvl[maxLevel]) + len(highLvl[maxLevel-1])
		} else if _, ok := highLvl[maxLevel]; ok {
			totalHighLevel = len(highLvl[maxLevel])
		}
		finalFreqs = make([]float64, totalHighLevel)
		if isInitializing {
			if goingUp && currentFreq <= radio.Constants.FreqRangeUs[1] {
				currentFreq += 0.1
			} else if goingUp && currentFreq > radio.Constants.FreqRangeUs[1] {
				currentFreq = radio.Constants.FreqRangeUs[1]
				goingUp = false
				numberOfCycles += 1
			} else if !goingUp && currentFreq >= radio.Constants.FreqRangeUs[0] {
				currentFreq -= 0.3
			} else {
				numberOfCycles += 1
				currentFreq = radio.Constants.FreqRangeUs[0]
				goingUp = true
			}
		} else {
			if len(finalFreqs) != totalHighLevel || finalFreqs[0] != highLvl[maxLevel][0] {
				if _, ok := highLvl[maxLevel-2]; ok {
					finalFreqs = append(highLvl[maxLevel], highLvl[maxLevel-1]...)
					finalFreqs = append(finalFreqs, highLvl[maxLevel-2]...)
				} else if _, ok := highLvl[maxLevel-1]; ok {
					finalFreqs = append(highLvl[maxLevel], highLvl[maxLevel-1]...)
				} else if _, ok := highLvl[maxLevel]; ok {
					finalFreqs = highLvl[maxLevel]
				}
			}

			if goingUp && currentFrame < len(finalFreqs)-1 {
				currentFrame += 1
			} else if goingUp && currentFrame == len(finalFreqs)-1 {
				currentFrame -= 1
				goingUp = false
			} else if !goingUp && currentFrame > 0 {
				currentFrame -= 1
			} else {
				currentFrame += 1
				goingUp = true
			}
			sharedutils.Log("3 " + strconv.FormatInt(int64(currentFrame+1), 10) + " " + strconv.FormatFloat(finalFreqs[currentFrame], 'f', 1, 64))

			// if _, ok := finalFreqs[currentFrame]; ok {
			currentFreq = finalFreqs[currentFrame]
			// }
		}

		response := tea5767.SetFrequency(radio, currentFreq)

		var signalStrength int64

		if isInitializing {

			signalStrength = int64(response[3] >> 4)
			if maxLevel < signalStrength {
				maxLevel = signalStrength

				// cleanup the lower-signal
				for i := 0; i < (int(maxLevel) - 2); i++ {
					if _, ok := highLvl[maxLevel]; ok {
						delete(highLvl, maxLevel)
					}
				}
			}

			// only add if in the top channels
			if signalStrength >= (maxLevel - 2) {
				highLvl[signalStrength] = append(highLvl[signalStrength], currentFreq)
				highLvl[signalStrength] = uniqueFloat(highLvl[signalStrength])
			}

		}

		display.Clear(screen)

		if isInitializing {
			display.Text(
				screen,
				2,
				15,
				"Initializing...")
			display.Text(
				screen,
				2,
				35,
				"# signal: "+strconv.FormatInt(int64(totalHighLevel), 10))
			display.Text(
				screen,
				2,
				55,
				"Freq: "+strconv.FormatFloat(currentFreq, 'f', 1, 64))
			if numberOfCycles > 1 {
				display.FilledRect(
					screen,
					107,
					117,
					1,
					11,
					color.RGBA{100, 100, 100, 255},
				)
			} else {
				display.Rect(
					screen,
					107,
					117,
					1,
					11,
					color.RGBA{100, 100, 100, 255},
				)
			}
			if numberOfCycles > 2 {
				display.FilledRect(
					screen,
					117,
					127,
					1,
					11,
					color.RGBA{100, 100, 100, 255},
				)

			} else {
				display.Rect(
					screen,
					117,
					127,
					1,
					11,
					color.RGBA{100, 100, 100, 255},
				)
			}
			if numberOfCycles > 3 {
				display.FilledRect(
					screen,
					107,
					117,
					11,
					21,
					color.RGBA{100, 100, 100, 255},
				)

			} else {
				display.Rect(
					screen,
					107,
					117,
					11,
					21,
					color.RGBA{100, 100, 100, 255},
				)
			}

			display.Rect(
				screen,
				117,
				127,
				11,
				21,
				color.RGBA{100, 100, 100, 255},
			)
		} else {

			display.FilledRect(
				screen,
				107,
				127,
				1,
				21,
				color.RGBA{100, 100, 100, 255},
			)
			display.FilledBolt(
				screen,
				117,
				11,
				color.RGBA{0, 0, 0, 255},
			)

			var dirStr string
			if goingUp {
				dirStr = "up"
			} else {
				dirStr = "dn"
			}

			display.Text(
				screen,
				2,
				15,
				"Scan "+dirStr+" "+strconv.FormatInt(int64(currentFrame), 10))
			display.Text(
				screen,
				2,
				35,
				"# channels: "+strconv.FormatInt(int64(totalHighLevel), 10))
			display.Text(
				screen,
				2,
				55,
				"Freq: "+strconv.FormatFloat(currentFreq, 'f', 1, 64))
		}

		screen.Display()
		if isInitializing {
			time.Sleep(50 * time.Millisecond)
		} else {
			time.Sleep(400 * time.Millisecond)
		}
	}

}
