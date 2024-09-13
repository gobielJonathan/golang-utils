package utils

func Filter[T comparable](data []T, predict func(val T, index int) bool) []T {

	var result []T

	for idx, v := range data {
		if predict(v, idx) {
			result = append(result, v)
		}
	}

	return result
}
