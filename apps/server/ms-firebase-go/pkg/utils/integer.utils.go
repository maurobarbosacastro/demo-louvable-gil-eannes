package utils

func Int64Pointer(n int64) *int64 {
	return &n
}

func PointerInt64(value *int64) int64 {
	if value != nil {
		return *value
	}
	return 0
}
