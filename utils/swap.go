package utils

func Swap[T any](slice []T, a, b int) {
	slice[a], slice[b] = slice[b], slice[a]
}
