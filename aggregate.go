package stdlib

import "math"

type mode struct {
	counts   map[int]int
	top      int
	topCount int
}

func newMode() *mode {
	return &mode{
		counts: map[int]int{},
	}
}

func (m *mode) Step(x int) {
	m.counts[x]++
	c := m.counts[x]
	if c > m.topCount {
		m.top = x
		m.topCount = c
	}
}

func (m *mode) Done() int {
	return m.top
}

type modestr struct {
	counts   map[string]int
	top      string
	topCount int
}

func newModestr() *modestr {
	return &modestr{
		counts: map[string]int{},
	}
}

func (m *modestr) Step(x string) {
	m.counts[x]++
	c := m.counts[x]
	if c > m.topCount {
		m.top = x
		m.topCount = c
	}
}

func (m *modestr) Done() string {
	return m.top
}

// SOURCE: https://github.com/mattn/go-sqlite3/blob/master/_example/custom_func/main.go
type stddev struct {
	xs []int64
	// Running average calculation
	sum int64
	n   int64
}

func newStddev() *stddev { return &stddev{} }

func (s *stddev) Step(x int64) {
	s.xs = append(s.xs, x)
	s.sum += x
	s.n++
}

func (s *stddev) Done() float64 {
	mean := float64(s.sum) / float64(s.n)
	var sqDiff []float64
	for _, x := range s.xs {
		sqDiff = append(sqDiff, math.Pow(float64(x)-mean, 2))
	}
	var dev float64
	for _, x := range sqDiff {
		dev += x
	}
	dev /= float64(len(sqDiff))
	return math.Sqrt(dev)
}

var aggregateFunctions = map[string]any{
	"stddev":     newStddev,
	"stdev":      newStddev,
	"stddev_pop": newStddev,
	"mode":       newMode,
	"modestr":    newModestr,
}
