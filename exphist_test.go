package exphist

import (
	"testing"
)

func TestNew(t *testing.T) {
	hist := New(1, 0.5)
	if hist.mergeSize != 3 {
		t.Errorf("ExpHist should have %d as mergeSize, but %d", 3, hist.mergeSize)
	}

	hist = New(1, 0.01)
	if hist.mergeSize != 52 {
		t.Errorf("ExpHist should have %d as mergeSize, but %d", 52, hist.mergeSize)
	}
}

func TestBits(t *testing.T) {
	hist := New(10, 0.5)
	stream := []uint{1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0}
	for _, x := range stream {
		hist.Add(x)
	}

	if hist.time != uint(len(stream)) {
		t.Errorf("Time should be %d, but %d", len(stream), hist.time)
	}

	if hist.total != 8 {
		t.Errorf("Total should be %d, but %d", 8, hist.total)
	}

	if hist.last != 4 {
		t.Errorf("Last should be %d, but %d", 4, hist.last)
	}

	if count := hist.Count(); count != 6.0 {
		t.Errorf("Count should return %f, but %f", 6.0, count)
	}

	// Expectation of the buckets
	// Bucket:      [    ] [  ] [ ] [ ]
	// Bucket size:    2^2  2^1 2^0 2^0
	// TimeStamp:      ~10  ~12 ~13 ~14

	bucketNum := 0
	for i, _ := range hist.buckets {
		bucketNum += len(hist.buckets[i])
	}
	if bucketNum != 4 {
		t.Errorf("Number of buckets should be %d, but %d", 4, bucketNum)
	}

	// 0 of 2^0
	if time := hist.buckets[0][0]; time != 13 {
		t.Errorf("TimeStamp at 0 of 2^0 should be %d, but %d", 13, time)
	}

	// 1 of 2^0
	if time := hist.buckets[0][1]; time != 14 {
		t.Errorf("TimeStamp at 1 of 2^0 should be %d, but %d", 14, time)
	}

	// 0 of 2^1
	if time := hist.buckets[1][0]; time != 12 {
		t.Errorf("TimeStamp at 0 of 2^1 should be %d, but %d", 12, time)
	}

	// 0 of 2^2
	if time := hist.buckets[2][0]; time != 10 {
		t.Errorf("TimeStamp at 0 of 2^2 should be %d, but %d", 10, time)
	}
}

func TestPositiveIntegers(t *testing.T) {
	hist := New(200, 0.01)
	for i := 1; i <= 200; i++ {
		hist.Add(uint(i))
	}

	if hist.time != 200 {
		t.Errorf("Time should be %d, but %d", 200, hist.time)
	}
	if hist.total != 20100 {
		t.Errorf("Total should be %d, but %d", 20100, hist.total)
	}
	if hist.last != 256 {
		t.Errorf("Last should be %d, but %d", 256, hist.last)
	}

	if count := hist.Count(); count != 19972.0 {
		t.Errorf("Count should return %f, but %f", 19972.0, count)
	}
}
