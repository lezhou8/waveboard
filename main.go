package main

import (
	"image/color"
	"log"
	"math"
	"math/cmplx"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"gonum.org/v1/gonum/dsp/fourier"
)

const (
	width           = 690
	height          = 500
	columnWidth     = 10
	sampleRate      = 44100
	columns     int = width / columnWidth
)

var (
	hasStarted bool
	titleFont  font.Face
	fft        fourier.FFT
)

type game struct{}

func (g *game) Update() error {
	if !hasStarted {
		if len(inpututil.AppendJustReleasedKeys(nil)) != 0 {
			hasStarted = true
			start = time.Now()
		}
		return nil
	}

	for _, key := range filterOutIgnoredKeys(inpututil.AppendJustPressedKeys(nil)) {
		keyNum := ebitenKeyKeyboardKeyMap[key]
		freq := keyBoardKeyNumToFreq(keyNum)
		freqs.Add(freq)
		keyLog(freq, keyNum, press)
	}
	for _, key := range filterOutIgnoredKeys(inpututil.AppendJustReleasedKeys(nil)) {
		keyNum := ebitenKeyKeyboardKeyMap[key]
		freq := keyBoardKeyNumToFreq(keyNum)
		freqs.Remove(freq)
		keyLog(freq, keyNum, release)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	if !hasStarted {
		text.Draw(screen, "Press any key to start", titleFont, 30, 50, color.White)
		return
	}

	amps := make([]float64, columns)
	for i := range amps {
		sum := 0.0
		for freq := range freqs.Iter() {
			sum += math.Sin(twoPi * freq.(float64) * float64(i) / float64(columns))
		}
		amps[i] = sum
	}
	coeffs := fft.Coefficients(nil, amps)
	coeffsLen := len(coeffs)

	columnWidth := float32(width / coeffsLen)
	for i := 0; i < coeffsLen; i++ {
		columnLen := float32(cmplx.Abs(coeffs[i])) * 10.0
		vector.DrawFilledRect(
			screen,
			float32(i)*columnWidth,
			height-columnLen-20,
			columnWidth,
			columnLen,
			color.RGBA{0xff, uint8((float32(i) / float32(coeffsLen)) * 180), 0x00, 0xff},
			false,
		)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	sr := beep.SampleRate(sampleRate)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(sound())

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	titleFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	hasStarted = false
	fft = *fourier.NewFFT(columns)

	ebiten.SetWindowTitle("Wave Board")
	ebiten.SetWindowSize(width, height)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
