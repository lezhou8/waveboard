package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var ebitenKeyKeyboardKeyMap = map[ebiten.Key]int{
	ebiten.KeyControlLeft:  19,
	ebiten.KeyAltLeft:      20,
	ebiten.KeySpace:        21,
	ebiten.KeyAltRight:     22,
	ebiten.KeyControlRight: 23,
	ebiten.KeyPageUp:       24,
	ebiten.KeyPageDown:     25,
	ebiten.KeyLeft:         26,
	ebiten.KeyDown:         27,
	ebiten.KeyRight:        28,
	ebiten.KeyUp:           29,
	ebiten.KeyShiftLeft:    30,
	ebiten.KeyZ:            31,
	ebiten.KeyX:            32,
	ebiten.KeyC:            33,
	ebiten.KeyV:            34,
	ebiten.KeyB:            35,
	ebiten.KeyN:            36,
	ebiten.KeyM:            37,
	ebiten.KeyComma:        38,
	ebiten.KeyPeriod:       39,
	ebiten.KeySlash:        40,
	ebiten.KeyShiftRight:   41,
	ebiten.KeyA:            42,
	ebiten.KeyS:            44,
	ebiten.KeyD:            45,
	ebiten.KeyF:            46,
	ebiten.KeyG:            47,
	ebiten.KeyH:            48,
	ebiten.KeyJ:            49,
	ebiten.KeyK:            50,
	ebiten.KeyL:            51,
	ebiten.KeySemicolon:    52,
	ebiten.KeyQuote:        53,
	ebiten.KeyEnter:        54,
	ebiten.KeyTab:          56,
	ebiten.KeyQ:            57,
	ebiten.KeyW:            58,
	ebiten.KeyE:            59,
	ebiten.KeyR:            60,
	ebiten.KeyT:            61,
	ebiten.KeyY:            62,
	ebiten.KeyU:            63,
	ebiten.KeyI:            64,
	ebiten.KeyO:            65,
	ebiten.KeyP:            66,
	ebiten.KeyBracketLeft:  67,
	ebiten.KeyBracketRight: 68,
	ebiten.KeyBackslash:    69,
	ebiten.KeyGraveAccent:  70,
	ebiten.KeyDigit1:       71,
	ebiten.KeyDigit2:       72,
	ebiten.KeyDigit3:       73,
	ebiten.KeyDigit4:       74,
	ebiten.KeyDigit5:       75,
	ebiten.KeyDigit6:       76,
	ebiten.KeyDigit7:       77,
	ebiten.KeyDigit8:       78,
	ebiten.KeyDigit9:       79,
	ebiten.KeyDigit0:       80,
	ebiten.KeyMinus:        81,
	ebiten.KeyEqual:        82,
	ebiten.KeyBackspace:    83,
	ebiten.KeyEscape:       84,
	ebiten.KeyHome:         85,
	ebiten.KeyEnd:          86,
	ebiten.KeyInsert:       87,
	ebiten.KeyDelete:       88,
}

var keyNumNoteMap = map[int]string{
	19: "D#2/Eb2",
	20: "E2",
	21: "F2",
	22: "F#2/Bb2",
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

func keyBoardKeyNumToFreq(key int) float64 {
	return 440 * math.Pow(2, (float64(key)-49)/12)
}

func filterOutIgnoredKeys(keys []ebiten.Key) []ebiten.Key {
	var filtered []ebiten.Key
	for _, key := range keys {
		if key == ebiten.KeyShift || key == ebiten.KeyMeta ||
			key == ebiten.KeyMetaLeft || key == ebiten.KeyMetaRight ||
			key == ebiten.KeyMax || key == ebiten.KeyContextMenu ||
			key == ebiten.KeyCapsLock || key == ebiten.KeyPause ||
			key == ebiten.KeyScrollLock || key == ebiten.KeyPrintScreen ||
			key == ebiten.KeyNumLock || key == ebiten.KeyNumpad0 ||
			key == ebiten.KeyNumpad1 || key == ebiten.KeyNumpad2 ||
			key == ebiten.KeyNumpad3 || key == ebiten.KeyNumpad4 ||
			key == ebiten.KeyNumpad5 || key == ebiten.KeyNumpad6 ||
			key == ebiten.KeyNumpad7 || key == ebiten.KeyNumpad8 ||
			key == ebiten.KeyNumpad9 || key == ebiten.KeyNumpadDivide ||
			key == ebiten.KeyNumpadMultiply ||
			key == ebiten.KeyNumpadSubtract ||
			key == ebiten.KeyNumpadAdd || key == ebiten.KeyNumpadDecimal ||
			key == ebiten.KeyNumpadEnter || key == ebiten.KeyNumpadEqual ||
			key == ebiten.KeyF1 || key == ebiten.KeyF2 ||
			key == ebiten.KeyF3 || key == ebiten.KeyF4 ||
			key == ebiten.KeyF5 || key == ebiten.KeyF6 ||
			key == ebiten.KeyF7 || key == ebiten.KeyF8 ||
			key == ebiten.KeyF9 || key == ebiten.KeyF10 ||
			key == ebiten.KeyF11 || key == ebiten.KeyF12 {
			continue
		}
		filtered = append(filtered, key)
	}
	return filtered
}
