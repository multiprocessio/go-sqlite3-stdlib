package stdlib

import "math"

func ext_degrees(rad float64) float64 {
	return rad * (180 / math.Pi)
}

func ext_pi() float64 {
	return math.Pi
}

func ext_radians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// Rounds toward zero
func ext_trunc(f float64) float64 {
	if f >= 0 {
		return math.Floor(f)
	}

	r := math.Floor(math.Abs(f))
	if r == 0 {
		return 0
	}

	return -1 * r
}

var mathFunctions = map[string]any{
	"acos":  math.Acos,
	"acosh": math.Acosh,
	"asin":  math.Asin,
	"asinh": math.Asinh,
	"atan":  math.Atan,
	// TODO: atan2
	"atanh":   math.Atanh,
	"ceil":    math.Ceil,
	"ceiling": math.Ceil,
	"cos":     math.Cos,
	"cosh":    math.Cosh,
	"degrees": ext_degrees,
	"exp":     math.Exp,
	"floor":   math.Floor,
	"ln":      math.Log,
	"log":     math.Log,
	"log10":   math.Log10,
	// TODO: support log(B, X)
	"log2":     math.Log2,
	"mod":      math.Mod,
	"pi":       ext_pi,
	"pow":      math.Pow,
	"power":    math.Pow,
	"radians":  ext_radians,
	"sin":      math.Sin,
	"sinh":     math.Sinh,
	"sqrt":     math.Sqrt,
	"tan":      math.Tan,
	"tanh":     math.Tanh,
	"trunc":    ext_trunc,
	"truncate": ext_trunc,
}
