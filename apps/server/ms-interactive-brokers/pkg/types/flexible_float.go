package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FlexibleFloat64 handles JSON fields that might be strings or numbers
// This is useful when APIs return numeric values as strings
type FlexibleFloat64 float64

// UnmarshalJSON implements json.Unmarshaler interface
// It handles both numeric and string JSON values, converting them to float64
func (f *FlexibleFloat64) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a numeric value first
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*f = FlexibleFloat64(num)
		return nil
	}

	// Try to unmarshal as a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// Handle empty string as 0
	if str == "" {
		*f = 0
		return nil
	}

	// Parse string to float
	parsed, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("cannot parse %q as float64: %w", str, err)
	}

	*f = FlexibleFloat64(parsed)
	return nil
}

// MarshalJSON implements json.Marshaler interface
// It marshals the float64 value as a JSON number
func (f FlexibleFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}

// Float64 returns the float64 value
func (f FlexibleFloat64) Float64() float64 {
	return float64(f)
}
