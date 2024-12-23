package utils

func CartesianProduct[T any](slices ...[]T) [][]T {
	if len(slices) == 0 {
		return [][]T{{}}
	}
	first := slices[0]
	remaining := CartesianProduct(slices[1:]...)

	var result [][]T

	for _, element := range first {
		for _, combination := range remaining {
			newCombination := append([]T{element}, combination...)
			result = append(result, newCombination)
		}
	}

	return result
}
