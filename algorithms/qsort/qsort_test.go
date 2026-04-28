package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{3, 2, 1, 5, 4}
	QuickSort(values, 0, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		if values[i] > values[i+1] {
			t.Error("QuickSort failed. Got", values)
		}
	}
}

func TestQuickSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	QuickSort(values, 0, len(values)-1)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("QuickSort failed. Got", values)
	}
}

func TestQuickSort3(t *testing.T) {
	values := []int{5}
	QuickSort(values, 0, len(values)-1)
	if values[0] != 5 {
		t.Error("QuickSort failed. Got", values)
	}
}
