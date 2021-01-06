package typeUtils

import "fmt"

func ItfToString(src interface{}) (dst string, err error) {
	if src == nil {
		return "", nil
	}
	switch src.(type) {
	case []byte:
		dst = string(src.([]byte)[:])
	case string:
		dst = src.(string)
	case fmt.Stringer:
		dst = src.(fmt.Stringer).String()
	default:
		return "", ErrInvalidSourceType
	}
	return
}

func ItfToStringConvertible(src interface{}) error {
	switch src.(type) {
	case []byte, string, fmt.Stringer:
		return nil
	default:
		return ErrInvalidSourceType
	}
}

func StringFallback(val, fallback string) string {
	if val != "" {
		return val
	}
	return fallback
}

func IndexOfStringSlice(raw []string, target string) int {
	var i int
	for i = len(raw) - 1; i >= 0; i-- {
		if target == raw[i] {
			return i
		}
	}
	return i
}
