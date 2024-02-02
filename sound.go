package main

import (
	"math"

	mapset "github.com/deckarep/golang-set"
	"github.com/gopxl/beep"
)

const twoPi = 2.0 * math.Pi

var freqs = mapset.NewSet()

func sineFreqCalc(t float64) float64 {
	sum := 0.0
	for freq := range freqs.Iter() {
		sum += math.Sin(twoPi * freq.(float64) * t)
	}
	return sum
}

func sound() beep.Streamer {
	sr := beep.SampleRate(sampleRate)
	t := 0.0
	dt := sr.D(1).Seconds()
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			v := sineFreqCalc(t)
			samples[i][0], samples[i][1] = v, v
			t += dt
		}
		return len(samples), true
	})
}
