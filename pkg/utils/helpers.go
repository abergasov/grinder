package utils

func SliceHasIntersections(need, slice []int64) bool {
	for i := range slice {
		for j := range need {
			if need[j] == slice[i] {
				return true
			}
		}
	}
	return false
}

func Int64InSlice(need int64, slice []int64) bool {
	for i := range slice {
		if need == slice[i] {
			return true
		}
	}
	return false
}
