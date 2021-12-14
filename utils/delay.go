package utils

import (
	"time"
)

func Delay(t int64) {
    time.Sleep(time.Duration(1000000 * t))
}