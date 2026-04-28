package bubblesort

import "testing"

func TestBubblesort1(t *testing.T) {
	values := []int{3, 2, 1, 5, 4}
	Bubblesort(values)
	for i := 0; i < len(values)-1; i++ {
		if values[i] > values[i+1] {
			t.Error("Bubblesort failed. Got", values)
		}
	}
}

func TestBubblesort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	Bubblesort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("Bubblesort failed. Got", values)
	}
}

func TestBubblesort3(t *testing.T) {
	values := []int{5}
	Bubblesort(values)
	if values[0] != 5 {
		t.Error("Bubblesort failed. Got", values)
	}
}
