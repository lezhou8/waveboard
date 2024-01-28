package main

import (
	"log"
	"math"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/eiannone/keyboard"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/speaker"
)

const (
	baseFreq           = 440.0
	freqIncreaseFactor = 20.0
	asciiShiftDown     = 32
	asciiBottomLimit   = 33
	asciiTopLimit      = 125
)

var freqSet = mapset.NewSet()

type playLengthMode int

const (
	short playLengthMode = iota
	long
	extraLong
)

func inIgnoredKeys(char rune) bool {
	if char < asciiBottomLimit || asciiTopLimit < char || char == '`' ||
		char == '+' || char == '-' || char == '=' || char == '_' {
		return true
	}
	return false
}

func asciiToFreq(note int) float64 {
	return math.Pow(2, (float64(note)-49)/12) * 440
}

func freqErp(time float64) float64 {
	freqSum := 0.0
	for f := range freqSet.Iter() {
		freqSum += math.Sin(2 * math.Pi * f.(float64) * time)
	}
	return freqSum
}

func charKeyMap(char rune) int {
	ascii := int(char)
	switch {
	case 96 < ascii:
		ascii -= 5
	case 61 < ascii:
		ascii -= 3
	case 45 < ascii:
		ascii -= 2
	case 43 < ascii:
		ascii -= 1
	}
	return ascii - asciiShiftDown
}

func noise(deltaTime float64) beep.Streamer {
	t := 0.0
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			y := freqErp(t)
			samples[i][0], samples[i][1] = y, y
			t += deltaTime
		}
		return len(samples), true
	})
}

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	ctrl := &beep.Ctrl{Streamer: noise(sr.D(1).Seconds()), Paused: false}
	volume := &effects.Volume{Streamer: ctrl, Base: 2, Volume: 0, Silent: false}
	speaker.Play(volume)
	keyPressLength := short

	if err := keyboard.Open(); err != nil {
		log.Fatalf("Error while trying to open keyboard: %s", err)
	}
	defer keyboard.Close()

	numNoteMap := map[int]string{
		1:  "A0",
		2:  "A#0/Bb0",
		3:  "B0",
		4:  "C1",
		5:  "C#1/Db1",
		6:  "D1",
		7:  "D#1/Eb1",
		8:  "E1",
		9:  "F1",
		10: "F#1/Gb1",
		11: "G1",
		12: "G#1/Ab1",
		13: "A1",
		14: "A#1/Bb1",
		15: "B1",
		16: "C2",
		17: "C#2/Db2",
		18: "D2",
		19: "D#2/Eb2",
		20: "E2",
		21: "F2",
		22: "F#2/Gb2",
		23: "G2",
		24: "G#2/Ab2",
		25: "A2",
		26: "A#2/Bb2",
		27: "B2",
		28: "C3",
		29: "C#3/Db3",
		30: "D3",
		31: "D#3/Eb3",
		32: "E3",
		33: "F3",
		34: "F#3/Gb3",
		35: "G3",
		36: "G#3/Ab3",
		37: "A3",
		38: "A#3/Bb3",
		39: "B3",
		40: "C4",
		41: "C#4/Db4",
		42: "D4",
		43: "D#4/Eb4",
		44: "E4",
		45: "F4",
		46: "F#4/Gb4",
		47: "G4",
		48: "G#4/Ab4",
		49: "A4",
		50: "A#4/Bb4",
		51: "B4",
		52: "C5",
		53: "C#5/Db5",
		54: "D5",
		55: "D#5/Eb5",
		56: "E5",
		57: "F5",
		58: "F#5/Gb5",
		59: "G5",
		60: "G#5/Ab5",
		61: "A5",
		62: "A#5/Bb5",
		63: "B5",
		64: "C6",
		65: "C#6/Db6",
		66: "D6",
		67: "D#6/Eb6",
		68: "E6",
		69: "F6",
		70: "F#6/Gb6",
		71: "G6",
		72: "G#6/Ab6",
		73: "A6",
		74: "A#6/Bb6",
		75: "B6",
		76: "C7",
		77: "C#7/Db7",
		78: "D7",
		79: "D#7/Eb7",
		80: "E7",
		81: "F7",
		82: "F#7/Gb7",
		83: "G7",
		84: "G#7/Ab7",
		85: "A7",
		86: "A#7/Bb7",
		87: "B7",
		88: "C8",
	}

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Printf("Error while getting keypress: %s", err)
		}

		if key == keyboard.KeyEsc {
			return
		} else if key == keyboard.KeyCtrlC {
			return
		} else if key == keyboard.KeySpace {
			ctrl.Paused = true
			log.Println("Rest")
			continue
		}

		if char == '+' {
			volume.Volume += 0.1
			log.Printf("Volume: %f\n", volume.Volume)
		} else if char == '-' {
			volume.Volume -= 0.1
			log.Printf("Volume: %f\n", volume.Volume)
		} else if char == '=' {
			volume.Volume = 0.0
			log.Printf("Volume: %f\n", volume.Volume)
		} else if char == '~' {
			if keyPressLength == short || keyPressLength == extraLong {
				keyPressLength = long
				log.Printf("Key press length: long\n")
			} else {
				keyPressLength = short
				log.Printf("Key press length: short\n")
			}
		} else if char == '_' {
			if keyPressLength == short || keyPressLength == long {
				keyPressLength = extraLong
				log.Printf("Key press length: extra long\n")
			} else {
				keyPressLength = short
				log.Printf("Key press length: short\n")
			}

		}

		if inIgnoredKeys(char) {
			continue
		}

		ctrl.Paused = false
		ascii := charKeyMap(char)
		freq := asciiToFreq(ascii)
		freqSet.Add(freq)

		go func() {
			if keyPressLength == extraLong {
				time.Sleep(1000 * time.Millisecond)
			} else if keyPressLength == long {
				time.Sleep(500 * time.Millisecond)
			} else {
				time.Sleep(250 * time.Millisecond)
			}
			freqSet.Remove(freq)
		}()

		log.Printf("Playing %09.4fHz (%s)\n", freq, numNoteMap[ascii])
	}
}
