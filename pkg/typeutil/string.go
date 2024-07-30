package typeutil

func StringToPtr(s string) *string {
	return &s
}

func StringSliceToPtr(s []string) *[]string {
	return &s
}
