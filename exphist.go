package exphist

import "math"

type ExpHist struct {
	windowSize uint
	mergeSize  int

	time    uint
	total   uint
	last    uint
	buckets [][]uint
}

func New(windowSize int, epsilon float64) *ExpHist {
	return &ExpHist{
		windowSize: uint(windowSize),
		mergeSize:  int(math.Ceil(math.Ceil(1.0/epsilon)/2.0)) + 2,
		last:       1,
		buckets:    [][]uint{[]uint{}},
	}
}

func (e *ExpHist) Add(x uint) {
	e.time += 1
	e.drop()
	for i := 0; i < int(x); i++ {
		e.total += 1
		e.buckets[0] = append(e.buckets[0], e.time)

		e.merge()
	}
}

func (e *ExpHist) Count() float64 {
	return 0.0
}

func (e *ExpHist) drop() {
	if !e.hasExpiredBucket() {
		return
	}
	for {
		e.tail()
		if !e.hasExpiredBucket() {
			break
		}
	}
}

func (e *ExpHist) hasExpiredBucket() bool {
	if e.total == 0 || e.time <= e.windowSize {
		return false
	}
	expiryTime := e.time - e.windowSize
	return e.buckets[len(e.buckets)-1][0] < expiryTime
}

func (e *ExpHist) tail() {
	size := uint(math.Pow(2.0, float64(len(e.buckets)-1)))
	if len(e.buckets[len(e.buckets)-1]) == 1 {
		e.buckets = e.buckets[0 : len(e.buckets)-1]
		e.last = uint(math.Pow(2.0, float64(len(e.buckets)-1)))
	} else {
		e.buckets[len(e.buckets)-1] = e.buckets[len(e.buckets)-1][1:]
	}
	e.total -= uint(size)
}

func (e *ExpHist) merge() {
	for i, _ := range e.buckets {
		if len(e.buckets[i]) < e.mergeSize {
			continue
		}
		if i == len(e.buckets)-1 {
			e.buckets = append(e.buckets, []uint{})
			e.last = uint(math.Pow(2.0, float64(len(e.buckets)-1)))
		}
		e.buckets[i+1] = append(e.buckets[i+1], e.buckets[i][1])
		e.buckets[i] = e.buckets[i][2:]
	}
}
