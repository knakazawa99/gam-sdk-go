package typeutil

func ConvertToPointer[T any](v T) *T {
	return &v
}
