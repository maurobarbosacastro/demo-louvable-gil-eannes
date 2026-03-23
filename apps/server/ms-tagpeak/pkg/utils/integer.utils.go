package utils

import "math"

func Round2Digits(value float64) float64 {
	return math.Round(value*100) / 100
}
