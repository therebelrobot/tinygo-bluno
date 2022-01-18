package ctrl

import (
	"machine"
	"strconv"
	"time"

	utils "github.com/therebelrobot/tinygo-sandbox/utils/shared"
	xiaobutton "github.com/therebelrobot/tinygo-sandbox/utils/xiaoexpansion/button"
	display "github.com/therebelrobot/tinygo-sandbox/utils/xiaoexpansion/display"
	"tinygo.org/x/drivers/sdcard"
)

var (
	spi    machine.SPI
	sckPin machine.Pin
	sdoPin machine.Pin
	sdiPin machine.Pin
	csPin  machine.Pin
	ledPin machine.Pin
)
var sd = sdcard.New(machine.SPI0, machine.SPI0_SCK_PIN, machine.SPI0_SDO_PIN, machine.SPI0_SDI_PIN, machine.D3)

func initSd() {
	err := sd.Configure()
	if err != nil {
		utils.Log(err.Error())
		for {
			time.Sleep(time.Hour)
		}
	}
	// littlefs.New(&sd)
}

func SdCard() {
	// waitSerial()
	screen := display.Init()

	button := xiaobutton.Init()
	initSd()
	for {

		size := sd.Size()

		// // Configure FATFS with sector size (must match value in ff.h - use 512)
		// filesystem.Configure(&fatfs.Config{
		// 	SectorSize: 512,
		// })

		// // FAT is read-only for the time being; so just try to mount it
		// // and panic if it fails
		// if err := filesystem.Mount(); err != nil {
		// 	utils.Log("could not mount")
		// }

		isPressed := !(button.Get())
		if isPressed {
			// var _ tinyfs.Filesystem = filesystem

			// ========================================================
			// f, err := fs.Open("/")
			// if err != nil {
			// 	utils.Log("error fsopen")
			// 	return
			// }
			// files, err := f.Readdir(0)
			// if err != nil {
			// 	utils.Log("error freaddir")
			// 	return
			// }

			// for _, v := range files {
			// 	utils.Log(v.Name() + " " + strconv.FormatBool(v.IsDir()))
			// }
			utils.Log("pressing")

			display.Clear(screen)

			display.Text(
				screen,
				2,
				15,
				"size: "+
					strconv.FormatInt(int64(size), 10)+
					"\n"+
					"Say hello!\n"+
					"thenewcryptid")

			screen.Display()
		}
		time.Sleep(1 * time.Millisecond)

	}
}

// // Wait for user to open serial console
// func waitSerial() {
// 	for !machine.Serial.DTR() {
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }
