package utils

func Some[T comparable](data []T, interate func(val T) bool) bool {

	for _, v := range data {
		if interate(v) {
			return true
		}
	}

	return false
}
