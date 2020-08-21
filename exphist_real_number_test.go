package exphist

import (
	"math"
	"testing"
)

func TestExpHistRealNumberAdd(t *testing.T) {
	hist := NewForRealNumber(2)
	stream := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	for _, x := range stream {
		hist.Add(x)
	}

	if actual := hist.Size(); actual != len(stream) {
		t.Errorf("ExpHistRealNumber.Size should return %d, but %d", len(stream), actual)
	}

	expect := 0.0
	for _, s := range stream {
		expect += s
	}
	if actual := hist.Sum(); actual != expect {
		t.Errorf("ExpHistRealNumber.Sum should return %f, but %f", expect, actual)
	}

	// [7.7](cap:1) [12.1](cap:2) [11.0](cap:4)
	if actual := len(hist.buckets); actual != 3 {
		t.Errorf("ExpHistRealNumber should have %d buckets, but %d", 3, actual)
	}
}

func TestExpHistRealNumberTail(t *testing.T) {
	hist := NewForRealNumber(2)

	hist.Add(1.1) // [1.1](cap:1)
	buckets := hist.Tail()
	if actual := buckets.Size(); actual != 1 {
		t.Errorf("Tail should return buckets that has %d bucket, but %d", 1, actual)
	}
	if actual := buckets.Sum(); actual != 1.1 {
		t.Errorf("Tail should return buckets that has %f value, but %f", 1.1, actual)
	}

	hist.Add(2.2) // [1.1](cap:1) [2.2](cap:1)
	buckets = hist.Tail()
	if actual := buckets.Size(); actual != 1 {
		t.Errorf("Tail should return buckets that has %d bucket, but %d", 1, actual)
	}
	if actual := buckets.Sum(); actual != 1.1 {
		// The oldest bucket has 1.1 value, not 2.2
		t.Errorf("Tail should return buckets that has %f value, but %f", 1.1, actual)
	}

	hist.Add(100.0) // [100.0](cap:1) [3.3](cap:2)
	buckets = hist.Tail()
	if actual := buckets.Size(); actual != 2 {
		t.Errorf("Tail should return buckets that has %d bucket, but %d", 2, actual)
	}
	if actual := buckets.Sum(); !almostEqual(actual, 3.3) {
		// The oldest bucket has 1.1 + 2.2 value, not 100.0
		t.Errorf("Tail should return buckets that has %f value, but %f", 3.3, actual)
	}
}

func TestExpHistRealNumberDrop(t *testing.T) {
	// [1.1](cap:1) [2.2](cap:1)
	hist := NewForRealNumber(2)
	hist.Add(1.1)
	hist.Add(2.2)
	hist.Drop()
	if actual := hist.Size(); actual != 1 {
		t.Errorf("ExpHistRealNumber.Size should return %d after Drop, but %d", 1, actual)
	}
	if actual := hist.Sum(); actual != 2.2 {
		t.Errorf("ExpHistRealNumber.Sum should return %f after Drop, but %f", 2.2, actual)
	}

	// // [100.0](cap:1) [3.3](cap:2)
	hist = NewForRealNumber(2)
	hist.Add(1.1)
	hist.Add(2.2)
	hist.Add(100.0)
	hist.Drop()
	if actual := hist.Size(); actual != 1 {
		t.Errorf("ExpHistRealNumber.Size should return %d after Drop, but %d", 1, actual)
	}
	if actual := hist.Sum(); actual != 100.0 {
		t.Errorf("ExpHistRealNumber.Sum should return %f after Drop, but %f", 100.0, actual)
	}
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
