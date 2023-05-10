package slices

func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func MapErr[T, U any](slice []T, fn func(T) (U, error)) ([]U, error) {
	result := make([]U, len(slice))
	for i, v := range slice {
		var err error
		result[i], err = fn(v)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}