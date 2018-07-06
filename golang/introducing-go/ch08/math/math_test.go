package math

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]float64{2,4})
	if v != 3 {
		t.Error("Expected 3, got ", v)
	} else {
		t.Log("OK!")
	}
}

