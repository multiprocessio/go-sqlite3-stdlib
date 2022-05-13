package stdlib

import "math"

func ext_degrees(rad float64) float64 {
	return rad * (180 / math.Pi)
}

func ext_modulus(y, x float64) float64 {
	return y % x
}

func ext_pi() float64 {
	return math.Pi
}

func ext_radians(deg float64) float64 {
	return deg * (math.Pi / 180)
}
