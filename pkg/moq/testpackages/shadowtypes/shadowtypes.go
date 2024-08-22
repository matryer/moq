package shadowtypes

import (
	"github.com/matryer/moq/pkg/moq/testpackages/shadowtypes/types"
)

// ShadowTypes is a test interface.
type ShadowTypes interface {
	// ShadowString is a test method.
	ShadowString(string, types.String)

	// ShadowInt is a test method.
	ShadowInt(int, types.Int)
	// ShadowInt8 is a test method.
	ShadowInt8(int8, types.Int8)
	// ShadowInt16 is a test method.
	ShadowInt16(int16, types.Int16)
	// ShadowInt32 is a test method.
	ShadowInt32(int32, types.Int32)
	// ShadowInt64 is a test method.
	ShadowInt64(int64, types.Int64)

	// ShadowUint is a test method.
	ShadowUint(uint, types.Uint)
	// ShadowUint8 is a test method.
	ShadowUint8(uint8, types.Uint8)
	// ShadowUint16 is a test method.
	ShadowUint16(uint16, types.Uint16)
	// ShadowUint32 is a test method.
	ShadowUint32(uint32, types.Uint32)
	// ShadowUint64 is a test method.
	ShadowUint64(uint64, types.Uint64)

	// ShadowFloat32 is a test method.
	ShadowFloat32(float32, types.Float32)
	// ShadowFloat64 is a test method.
	ShadowFloat64(float64, types.Float64)

	// ShadowByte is a test method.
	ShadowByte(byte, types.Byte)

	// ShadowRune is a test method.
	ShadowRune(rune, types.Rune)

	// ShadowBool is a test method.
	ShadowBool(bool, types.Bool)

	// ShadowComplex64 is a test method.
	ShadowComplex64(complex64, types.Complex64)
	// ShadowComplex128 is a test method.
	ShadowComplex128(complex128, types.Complex128)

	// ShadowUintptr is a test method.
	ShadowUintptr(uintptr, types.Uintptr)
}
