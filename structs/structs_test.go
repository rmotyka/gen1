package structs

import "testing"

func TestRoute_Equals(t *testing.T) {
	r1 := &Route{ []int {1, 2, 3},  float64(3)}
	r2 := &Route{ []int {1, 2, 3},  float64(3)}

	e := r1.Equals(r2)
	if !e {
		t.Error("Shuld be equal")
	}
}

func TestRoute_Equals1(t *testing.T) {
	r1 := &Route{ []int {1, 2, 3},  float64(3)}
	r2 := &Route{ []int {1, 2, 4},  float64(3)}

	e := r1.Equals(r2)
	if e {
		t.Error("Shuld not be equal")
	}
}
