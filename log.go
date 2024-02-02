package main

import (
	"fmt"
	"time"
)

type pressOrRelease int

const (
	press = iota
	release
)

var start time.Time

func keyLog(freq float64, keyNum int, action pressOrRelease) {
	var actionStr string
	if action == press {
		actionStr = "pressed"
	} else {
		actionStr = "released"
	}

	now := time.Since(start)
	fmt.Printf(
		"%02d:%02d:%03d %-8s %-7s (%08.3fHz)\n",
		int(now.Minutes()),
		int(now.Seconds())%60,
		int(now.Milliseconds())%1000,
		actionStr,
		keyNumNoteMap[keyNum],
		freq,
	)
}
