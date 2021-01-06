package typeUtils

import (
	"unsafe"
)

const SizeOfInt = unsafe.Sizeof(int(0))
const MaxInt = 1<<(SizeOfInt*4-1) - 1
const MinInt = -1 << (SizeOfInt*4 - 1)

func ItfToInt(src interface{}) (dst int, err error) {
	switch src.(type) {
	case int:
		return src.(int), nil
	case int8:
		return Int64ToInt(src.(int64))
	case int16:
		return Int64ToInt(src.(int64))
	case int32:
		return Int64ToInt(src.(int64))
	case int64:
		return Int64ToInt(src.(int64))
	case uint:
		return UInt64ToInt(src.(uint64))
	case uint8:
		return UInt64ToInt(src.(uint64))
	case uint16:
		return UInt64ToInt(src.(uint64))
	case uint32:
		return UInt64ToInt(src.(uint64))
	case uint64:
		return UInt64ToInt(src.(uint64))
	default:
		return 0, ErrInvalidSourceType
	}
}

func ItfToIntConvertible(src interface{}) error {
	switch src.(type) {
	case int,
		int8,
		int16,
		int32,
		int64,
		uint,
		uint8,
		uint16,
		uint32,
		uint64:
		return nil
	default:
		return ErrInvalidSourceType
	}
}

func Int64ToInt(src int64) (dst int, err error) {
	if src > int64(MaxInt) || src < int64(MinInt) {
		return 0, ErrUnsafeConversion
	}
	return int(src), nil
}

func UInt64ToInt(src uint64) (dst int, err error) {
	if src > uint64(MaxInt) {
		return 0, ErrUnsafeConversion
	}
	return int(src), nil
}
