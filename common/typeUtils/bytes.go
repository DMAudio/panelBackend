package typeUtils

import "fmt"

func ItfToBytes(src interface{}) (dst []byte, err error) {
	if src == nil {
		return nil, nil
	}
	switch src.(type) {
	case []byte:
		dst = src.([]byte)
	case string:
		dst = []byte(src.(string))
	case fmt.Stringer:
		dst = []byte(src.(fmt.Stringer).String())
	default:
		return nil, ErrInvalidSourceType
	}
	return
}

func ItfToBytesConvertible(src interface{}) error {
	switch src.(type) {
	case []byte, string, fmt.Stringer:
		return nil
	default:
		return ErrInvalidSourceType
	}
}
