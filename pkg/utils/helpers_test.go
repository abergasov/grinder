package utils

import "testing"

func TestSliceHasIntersections(t *testing.T) {
	big := []int64{3, 4, 5, 6}
	small := []int64{1, 3, 9}
	if !SliceHasIntersections(small, big) {
		t.Error("unexpected false")
	}

	small = []int64{11, 22}
	if SliceHasIntersections(small, big) {
		t.Error("unexpected true")
	}

	big = []int64{3}
	small = []int64{1, 3, 9}
	if !SliceHasIntersections(small, big) {
		t.Error("unexpected false")
	}

}
