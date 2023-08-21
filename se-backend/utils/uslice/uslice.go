package uslice

// Contains verifica se o valor existe no slice.
func Contains[T comparable](slice []T, element T) bool {
	for _, s := range slice {
		if element == s {
			return true
		}
	}

	return false
}
