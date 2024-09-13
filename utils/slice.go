package utils

func Slice[T comparable](data []T, start int, end int) []T {
	var result = []T{}

	for idx, v := range data {
		if idx >= start && idx <= end {
			result = append(result, v)
		}
	}

	return result
}
