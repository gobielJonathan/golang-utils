package utils

func Reduce[T comparable, DV comparable](data []T, interate func(curr DV, val T) DV, defaultValue DV) DV {

	var result DV = defaultValue

	for _, v := range data {
		result = interate(result, v)
	}

	return result
}
