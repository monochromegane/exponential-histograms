package exphist

import "testing"

func TestNew(t *testing.T) {
	e := New(1, 0.5)
	if e.mergeSize != 3 {
		t.Errorf("ExpHist should have %d as mergeSize, but %d", 3, e.mergeSize)
	}

	e = New(1, 0.01)
	if e.mergeSize != 52 {
		t.Errorf("ExpHist should have %d as mergeSize, but %d", 52, e.mergeSize)
	}
}
