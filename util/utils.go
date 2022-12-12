package util

// Contains check if a slice contains the searched item
func Contains[T comparable](container []T, item T) bool {
	for _, i := range container {
		if i == item {
			return true
		}
	}

	return false
}
