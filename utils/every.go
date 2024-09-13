package utils

func Every[T comparable](data []T, interate func(val T) bool) bool {

	var truthy = true

	for _, v := range data {
		truthy = truthy && interate(v)
	}

	return truthy
}
