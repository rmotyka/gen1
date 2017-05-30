package crossing

import (
	"testing"
	"fmt"
)

func TestCrossingIndexes(t *testing.T) {
	indexes := crossingIndexes(100, 1000)

	fmt.Println(indexes)
	if len(indexes) != 100 {
		t.Error("Wrong number of output items ", len(indexes))
	}

	for i := 0; i < len(indexes); i++  {
		if indexes[i] > 1000 {
			t.Error("Max index exceeded")
		}
	}

	for i := 0; i < len(indexes); i++  {
		for j := 0; j < len(indexes); j++  {
			if i != j && indexes[i] == indexes[j] {
				t.Error("Found duplicates")
			}
		}
	}
}