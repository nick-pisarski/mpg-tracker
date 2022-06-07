package utilities

func Map[T, K any](from []T, fn func(T) K) []K {
	result := make([]K, len(from))
	for i, v := range from {
		result[i] = fn(v)
	}
	return result
}
