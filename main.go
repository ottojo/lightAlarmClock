package main

import (
	"time"
	"github.com/ottojo/lights"
)

func main() {
	fade(lights.Color{0, 0, 0}, lights.Color{1, 0.55, 0}, 30*time.Second, 30*10)
	fade(lights.Color{1, 0.55, 0}, lights.Color{1, 1, 1}, 30*time.Second, 30*10)
}

func fade(current, target lights.Color, duration time.Duration, steps int) {
	rateR := target.R - current.R
	rateG := target.G - current.G
	rateB := target.B - current.B
	for current.R < target.R || current.G < target.G || current.B < target.B {
		current.R += rateR / float64(steps)
		current.G += rateG / float64(steps)
		current.B += rateB / float64(steps)
		go lights.SetAll(current)
		time.Sleep(time.Duration(duration.Nanoseconds()/int64(steps)) * time.Nanosecond)
	}
}
