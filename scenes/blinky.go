package scenes

import (
	"github.com/therebelrobot/tinygo-bluno/utils"
)

type SLedConfig = utils.SLedConfig
type SLedProject = utils.SLedProject
type SLedSection = utils.SLedSection
type SLedFrame = utils.SLedFrame

func GetConfig() SLedConfig {
	var sections []SLedSection

	sections[0] = SLedSection{
		Id:         0,
		H:          8,
		W:          8,
		Serpentine: true,
	}
	sections[1] = SLedSection{
		Id:         1,
		H:          8,
		W:          8,
		Serpentine: true,
	}

	var frames [][]SLedFrame

	frames[0][0] = SLedFrame{
		Section: 0,
		Row:     0,
		Col:     0,
		R:       0xff,
		G:       0xff,
		B:       0x00,
	}
	frames[0][1] = SLedFrame{
		Section: 0,
		Row:     1,
		Col:     0,
		R:       0xff,
		G:       0xff,
		B:       0x00,
	}
	frames[0][2] = SLedFrame{
		Section: 0,
		Row:     3,
		Col:     0,
		R:       0xff,
		G:       0xff,
		B:       0x00,
	}

	frames[1][0] = SLedFrame{
		Section: 0,
		Row:     0,
		Col:     0,
		R:       0x00,
		G:       0xff,
		B:       0x00,
	}
	frames[1][1] = SLedFrame{
		Section: 0,
		Row:     1,
		Col:     0,
		R:       0x00,
		G:       0x00,
		B:       0x00,
	}
	frames[1][2] = SLedFrame{
		Section: 0,
		Row:     3,
		Col:     0,
		R:       0x00,
		G:       0x00,
		B:       0x00,
	}

	results := SLedConfig{
		Version: "v0.0.1-alpha",
		Author:  "Aster (@therebelrobot)",
		Project: SLedProject{
			FrameDelay: 300,
			Looping:    true,
			FrameDrop:  true,
			Sections:   sections,
			Frames:     frames,
		},
	}

	return results
}
