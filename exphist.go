package exphist

import "math"

type ExpHist struct {
	windowSize uint
	mergeSize  uint

	time    uint
	total   uint
	last    uint
	buckets [][]uint
}

func New(windowSize int, epsilon float64) *ExpHist {
	return &ExpHist{
		windowSize: uint(windowSize),
		mergeSize:  uint(math.Ceil(math.Ceil(1.0/epsilon)/2.0)) + 2,
		last:       1,
		buckets:    [][]uint{[]uint{}},
	}
}

func (e *ExpHist) Add(x uint) {
}

func (e *ExpHist) Count() float64 {
	return 0.0
}
