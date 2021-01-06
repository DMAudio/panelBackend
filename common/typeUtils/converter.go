package typeUtils

import "github.com/DMAudio/panelBackend/package/errgo/fmt/errors"

var ErrInvalidTargetType = errors.New("invalid target type")
var ErrInvalidSourceType = errors.New("invalid source type")
var ErrUnsafeConversion = errors.New("unsafe conversion")

type ValueType string

const (
	Interface ValueType = "interface"
	String    ValueType = "string"
	Bytes     ValueType = "bytes"
	Int       ValueType = "int"
	UInt      ValueType = "uint"
)

func ItfToType(src interface{}, dstType ValueType) (interface{}, error) {
	switch dstType {
	case String:
		return ItfToString(src)
	case Bytes:
		return ItfToBytes(src)
	case Int:
		return ItfToInt(src)
	case Interface:
		return src, nil
	default:
		return nil, errors.Becausef(nil, ErrInvalidTargetType, "type:", dstType)
	}
}
