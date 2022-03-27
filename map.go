package gsf

func MapSlice[T any, U any](slice []T, mapper func(item T) U) []U {
	newSlice := []U{}
	for _, item := range slice {
		newSlice = append(newSlice, mapper(item))
	}
	return newSlice
}
