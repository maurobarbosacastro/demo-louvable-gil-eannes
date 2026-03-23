package utils

import "github.com/google/uuid"

func StringPointer(s string) *string {
	return &s
}

func PointerString(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}
func FloatPointer(s float64) *float64 {
	return &s
}

func UuidPointer(uuid uuid.UUID) *uuid.UUID { return &uuid }

func BoolPointer(b bool) *bool { return &b }

func Int64Pointer(n int64) *int64 {
	return &n
}

func PointerInt64(value *int64) int64 {
	if value != nil {
		return *value
	}
	return 0
}
