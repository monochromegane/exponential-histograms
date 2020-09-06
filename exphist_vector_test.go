package exphist

import (
	"testing"
)

func TestExpHistVectorAdd(t *testing.T) {
	hist := NewForVector(2)
	stream := [][]float64{{1.1, 0.11}, {2.2, 0.22}, {3.3, 0.33}, {4.4, 0.44}, {5.5, 0.55}, {6.6, 0.66}, {7.7, 0.77}}
	for _, x := range stream {
		hist.Add(x)
	}

	if actual := hist.Size(); actual != len(stream) {
		t.Errorf("ExpHistVector.Size should return %d, but %d", len(stream), actual)
	}

	expect := make([]float64, 2)
	for _, s := range stream {
		for i, _ := range s {
			expect[i] += s[i]
		}
	}
	sum := hist.Sum()
	for i, _ := range sum {
		if actual := sum[i]; actual != expect[i] {
			t.Errorf("ExpHistVector.Sum should return %f, but %f", expect[i], actual)
		}
	}

	// [7.7, 0.77](cap:1) [12.1, 1.21](cap:2) [11.0, 1.1](cap:4)
	if actual := len(hist.buckets); actual != 3 {
		t.Errorf("ExpHistVector should have %d buckets, but %d", 3, actual)
	}
}

func TestExpHistVectorTail(t *testing.T) {
	hist := NewForVector(2)

	hist.Add([]float64{1.1, 0.11}) // [1.1, 0.11](cap:1)
	buckets := hist.Tail()
	if actual := buckets.Size(); actual != 1 {
		t.Errorf("Tail should return buckets that has %d bucket, but %d", 1, actual)
	}
	sum := buckets.Sum()
	if actual := sum[0]; actual != 1.1 {
		t.Errorf("Tail should return buckets that has %f value, but %f", 1.1, actual)
	}
	if actual := sum[1]; actual != 0.11 {
		t.Errorf("Tail should return buckets that has %f value, but %f", 0.11, actual)
	}

	hist.Add([]float64{2.2, 0.22}) // [1.1, 0.11](cap:1) [2.2, 0.22](cap:1)
	buckets = hist.Tail()
	if actual := buckets.Size(); actual != 1 {
		t.Errorf("Tail should return buckets that has %d bucket, but %d", 1, actual)
	}
	sum = buckets.Sum()
	if actual := sum[0]; actual != 1.1 {
		// The oldest bucket has 1.1 value, not 2.2
		t.Errorf("Tail should return buckets that has %f value, but %f", 1.1, actual)
	}
	if actual := sum[1]; actual != 0.11 {
		// The oldest bucket has 0.11 value, not 0.22
		t.Errorf("Tail should return buckets that has %f value, but %f", 0.11, actual)
	}

	hist.Add([]float64{100.0, 10.0}) // [100.0, 10.0](cap:1) [3.3, 0.33](cap:2)
	buckets = hist.Tail()
	if actual := buckets.Size(); actual != 2 {
		t.Errorf("Tail should return buckets that has %d bucket, but %d", 2, actual)
	}
	sum = buckets.Sum()
	if actual := sum[0]; !almostEqual(actual, 3.3) {
		// The oldest bucket has 1.1 + 2.2 value, not 100.0
		t.Errorf("Tail should return buckets that has %f value, but %f", 3.3, actual)
	}
	if actual := sum[1]; !almostEqual(actual, 0.33) {
		// The oldest bucket has 0.11 + 0.22 value, not 10.0
		t.Errorf("Tail should return buckets that has %f value, but %f", 0.33, actual)
	}
}

func TestExpHistVector(t *testing.T) {
	// [1.1, 0.11](cap:1) [2.2, 0.22](cap:1)
	hist := NewForVector(2)
	hist.Add([]float64{1.1, 0.11})
	hist.Add([]float64{2.2, 0.22})
	hist.Drop()
	if actual := hist.Size(); actual != 1 {
		t.Errorf("ExpHistVector.Size should return %d after Drop, but %d", 1, actual)
	}
	sum := hist.Sum()
	if actual := sum[0]; actual != 2.2 {
		t.Errorf("ExpHistVector.Sum should return %f after Drop, but %f", 2.2, actual)
	}
	if actual := sum[1]; actual != 0.22 {
		t.Errorf("ExpHistVector.Sum should return %f after Drop, but %f", 0.22, actual)
	}

	// // [100.0, 10.0](cap:1) [3.3, 0.33](cap:2)
	hist = NewForVector(2)
	hist.Add([]float64{1.1, 0.11})
	hist.Add([]float64{2.2, 0.22})
	hist.Add([]float64{100.0, 10.0})
	hist.Drop()
	if actual := hist.Size(); actual != 1 {
		t.Errorf("ExpHistVector.Size should return %d after Drop, but %d", 1, actual)
	}
	sum = hist.Sum()
	if actual := sum[0]; actual != 100.0 {
		t.Errorf("ExpHistVector.Sum should return %f after Drop, but %f", 100.0, actual)
	}
	if actual := sum[1]; actual != 10.0 {
		t.Errorf("ExpHistVector.Sum should return %f after Drop, but %f", 10.0, actual)
	}
}
