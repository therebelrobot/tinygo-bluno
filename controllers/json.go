package ctrl

import (

	// "strconv"

	"fmt"

	"github.com/buger/jsonparser"
	utils "github.com/therebelrobot/tinygo-sandbox/utils/shared"
	xiaobutton "github.com/therebelrobot/tinygo-sandbox/utils/xiaoexpansion/button"
)

var data = []byte(`{
  "version": "v0.0.1-alpha",
  "author": "Aster (@therebelrobot)",
  "project": {
    "frameDelay": 300,
    "looping": true,
    "frameDrop": true,
    "sections": [
      {
        "id": 0,
        "h": 8,
        "w": 8,
        "serpentine": true,
        "display": {
          "x": 25,
          "y": 1,
          "rotation": 90
        }
      },
      {
        "id": 1,
        "h": 8,
        "w": 8,
        "serpentine": true,
        "display": {
          "x": 17,
          "y": 17,
          "rotation": -90
        }
      },
      {
        "id": 2,
        "h": 16,
        "w": 16,
        "serpentine": true,
        "display": {
          "x": 17,
          "y": 17,
          "rotation": 180
        }
      }
    ],
    "frames": [
      [
        {
          "section": 0,
          "row": 0,
          "col": 0,
          "r": 255,
          "g": 255,
          "b": 0
        },
        {
          "section": 0,
          "row": 1,
          "col": 0,
          "r": 255,
          "g": 255,
          "b": 0
        },
        {
          "section": 1,
          "row": 2,
          "col": 0,
          "r": 255,
          "g": 255,
          "b": 0
        }
      ],
      [
        {
          "section": 0,
          "row": 0,
          "col": 0,
          "r": 255,
          "g": 255,
          "b": 0
        },
        {
          "section": 0,
          "row": 1,
          "col": 0,
          "r": 255,
          "g": 255,
          "b": 0
        },
        {
          "section": 1,
          "row": 2,
          "col": 0,
          "r": 255,
          "g": 255,
          "b": 0
        }
      ]
    ]
  }
}`)

func Json() {

	button := xiaobutton.Init()
	for {

		isPressed := !(button.Get())
		if isPressed {

			var frameDelay int64
			var err error
			if value, err := jsonparser.GetInt(data, "project", "frameDelay"); err == nil {
				frameDelay = value
			}
			if err != nil {
				utils.Log("Error")
			}

			utils.Log("frameDelay: " + fmt.Sprintf("%v", frameDelay))
		}
	}
}
