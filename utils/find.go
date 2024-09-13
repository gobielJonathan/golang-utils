package utils

func Find[T comparable](data []T, interate func(val T) bool) *T {

	for _, v := range data {
		if interate(v) {
			return &v
		}
	}

	return nil
}
