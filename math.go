package stdlib

import "math"

func floaty(a any) float64 {
	switch t := a.(type) {
	case float64:
		return t
	case float32:
		return float64(t)
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case uint:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	default:
		return 0.0
	}
}

func ext_degrees(rad any) float64 {
	return floaty(rad) * (180 / math.Pi)
}

func ext_pi() float64 {
	return math.Pi
}

func ext_radians(deg any) float64 {
	return floaty(deg) * (math.Pi / 180)
}

// Rounds toward zero
func ext_trunc(a any) float64 {
	f := floaty(a)

	if f >= 0 {
		return math.Floor(f)
	}

	r := math.Floor(math.Abs(f))
	if r == 0 {
		return 0
	}

	return -1 * r
}

func floaty1Float64(f func(float64) float64) func(a any) float64 {
	return func(a any) float64 {
		return f(floaty(a))
	}
}

func floaty2Float64(f func(float64, float64) float64) func(a any, b any) float64 {
	return func(a any, b any) float64 {
		return f(floaty(a), floaty(b))
	}
}

var mathFunctions = map[string]any{
	"acos":  floaty1Float64(math.Acos),
	"acosh": floaty1Float64(math.Acosh),
	"asin":  floaty1Float64(math.Asin),
	"asinh": floaty1Float64(math.Asinh),
	"atan":  floaty1Float64(math.Atan),
	// TODO: atan2
	"atanh":   floaty1Float64(math.Atanh),
	"ceil":    floaty1Float64(math.Ceil),
	"ceiling": floaty1Float64(math.Ceil),
	"cos":     floaty1Float64(math.Cos),
	"cosh":    floaty1Float64(math.Cosh),
	"degrees": ext_degrees,
	"exp":     floaty1Float64(math.Exp),
	"floor":   floaty1Float64(math.Floor),
	"ln":      floaty1Float64(math.Log),
	"log":     floaty1Float64(math.Log),
	"log10":   floaty1Float64(math.Log10),
	// TODO: support log(B, X)
	"log2":     floaty1Float64(math.Log2),
	"mod":      floaty2Float64(math.Mod),
	"pi":       ext_pi,
	"pow":      floaty2Float64(math.Pow),
	"power":    floaty2Float64(math.Pow),
	"radians":  ext_radians,
	"sin":      floaty1Float64(math.Sin),
	"sinh":     floaty1Float64(math.Sinh),
	"sqrt":     floaty1Float64(math.Sqrt),
	"tan":      floaty1Float64(math.Tan),
	"tanh":     floaty1Float64(math.Tanh),
	"trunc":    ext_trunc,
	"truncate": ext_trunc,
}
