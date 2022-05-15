package stdlib

import (
	"bytes"
	"math"
	"sort"
)

type mode struct {
	counts   map[any]int
	top      any
	topCount int
}

func newMode() *mode {
	return &mode{
		counts: map[any]int{},
	}
}

func (m *mode) Step(x any) {
	m.counts[x]++
	c := m.counts[x]
	if c > m.topCount {
		m.top = x
		m.topCount = c
	}
}

func (m *mode) Done() any {
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

type sqliteValueKind uint

const (
	sqliteNull sqliteValueKind = iota
	sqliteInt
	sqliteString
	sqliteReal
	sqliteBlob
)

type sqliteValue struct {
	kind sqliteValueKind
	i    int64
	s    string
	r    float64
	b    []byte
}

type sqliteValues []sqliteValue

func (svs *sqliteValues) Len() int {
	return len(*svs)
}

func (svs *sqliteValues) Less(i, j int) bool {
	ie := (*svs)[i]
	je := (*svs)[j]
	if ie.kind != je.kind {
		// TODO: support mixed value types?
		return false
	}

	switch ie.kind {
	case sqliteInt:
		return ie.i < je.i
	case sqliteString:
		return ie.s < je.s
	case sqliteReal:
		return ie.r < je.r
	case sqliteBlob:
		return bytes.Compare(ie.b, je.b) < 0
	}

	return false
}

func (svs *sqliteValues) Swap(i, j int) {
	(*svs)[i], (*svs)[j] = (*svs)[j], (*svs)[i]
}

type median struct {
	xs sqliteValues
}

func newMedian() *median {
	return &median{}
}

func (m *median) Step(x any) {
	v := sqliteValue{kind: sqliteNull}
	switch t := x.(type) {
	case int64:
		v.kind = sqliteInt
		v.i = t
	case int:
		v.kind = sqliteInt
		v.i = int64(t)
	case string:
		v.kind = sqliteString
		v.s = t
	case float64:
		v.kind = sqliteReal
		v.r = t
	case []byte:
		v.kind = sqliteBlob
		v.b = t
	}
	m.xs = append(m.xs, v)
}

func (m *median) Done() any {
	if len(m.xs) == 0 {
		return nil
	}

	sort.Sort(&m.xs)
	e := m.xs[int(math.Floor(float64(len(m.xs))/2))]
	switch e.kind {
	case sqliteInt:
		return e.i
	case sqliteString:
		return e.s
	case sqliteReal:
		return e.r
	case sqliteBlob:
		return e.b
	}

	return nil
}

var aggregateFunctions = map[string]any{
	"stddev":     newStddev,
	"stdev":      newStddev,
	"stddev_pop": newStddev,
	"mode":       newMode,
	"median":     newMedian,
}
