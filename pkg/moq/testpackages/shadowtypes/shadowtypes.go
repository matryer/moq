package shadowtypes

import (
	"github.com/rewardStyle/moq/pkg/moq/testpackages/shadowtypes/types"
)

type ShadowTypes interface {
	ShadowString(string, types.String)

	ShadowInt(int, types.Int)
	ShadowInt8(int8, types.Int8)
	ShadowInt16(int16, types.Int16)
	ShadowInt32(int32, types.Int32)
	ShadowInt64(int64, types.Int64)

	ShadowUint(uint, types.Uint)
	ShadowUint8(uint8, types.Uint8)
	ShadowUint16(uint16, types.Uint16)
	ShadowUint32(uint32, types.Uint32)
	ShadowUint64(uint64, types.Uint64)

	ShadowFloat32(float32, types.Float32)
	ShadowFloat64(float64, types.Float64)

	ShadowByte(byte, types.Byte)

	ShadowRune(rune, types.Rune)

	ShadowBool(bool, types.Bool)

	ShadowComplex64(complex64, types.Complex64)
	ShadowComplex128(complex128, types.Complex128)

	ShadowUintptr(uintptr, types.Uintptr)
}
