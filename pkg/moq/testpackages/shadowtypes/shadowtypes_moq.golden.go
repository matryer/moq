// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package shadowtypes

import (
	"github.com/matryer/moq/pkg/moq/testpackages/shadowtypes/types"
	"sync"
)

// Ensure, that ShadowTypesMock does implement ShadowTypes.
// If this is not the case, regenerate this file with moq.
var _ ShadowTypes = &ShadowTypesMock{}

// ShadowTypesMock is a mock implementation of ShadowTypes.
//
//	func TestSomethingThatUsesShadowTypes(t *testing.T) {
//
//		// make and configure a mocked ShadowTypes
//		mockedShadowTypes := &ShadowTypesMock{
//			ShadowBoolFunc: func(b bool, boolMoqParam types.Bool)  {
//				panic("mock out the ShadowBool method")
//			},
//			ShadowByteFunc: func(v byte, byteMoqParam types.Byte)  {
//				panic("mock out the ShadowByte method")
//			},
//			ShadowComplex128Func: func(v complex128, complex128MoqParam types.Complex128)  {
//				panic("mock out the ShadowComplex128 method")
//			},
//			ShadowComplex64Func: func(v complex64, complex64MoqParam types.Complex64)  {
//				panic("mock out the ShadowComplex64 method")
//			},
//			ShadowFloat32Func: func(f float32, float32MoqParam types.Float32)  {
//				panic("mock out the ShadowFloat32 method")
//			},
//			ShadowFloat64Func: func(f float64, float64MoqParam types.Float64)  {
//				panic("mock out the ShadowFloat64 method")
//			},
//			ShadowIntFunc: func(n int, intMoqParam types.Int)  {
//				panic("mock out the ShadowInt method")
//			},
//			ShadowInt16Func: func(n int16, int16MoqParam types.Int16)  {
//				panic("mock out the ShadowInt16 method")
//			},
//			ShadowInt32Func: func(n int32, int32MoqParam types.Int32)  {
//				panic("mock out the ShadowInt32 method")
//			},
//			ShadowInt64Func: func(n int64, int64MoqParam types.Int64)  {
//				panic("mock out the ShadowInt64 method")
//			},
//			ShadowInt8Func: func(n int8, int8MoqParam types.Int8)  {
//				panic("mock out the ShadowInt8 method")
//			},
//			ShadowRuneFunc: func(n rune, runeMoqParam types.Rune)  {
//				panic("mock out the ShadowRune method")
//			},
//			ShadowStringFunc: func(s string, stringMoqParam types.String)  {
//				panic("mock out the ShadowString method")
//			},
//			ShadowUintFunc: func(v uint, uintMoqParam types.Uint)  {
//				panic("mock out the ShadowUint method")
//			},
//			ShadowUint16Func: func(v uint16, uint16MoqParam types.Uint16)  {
//				panic("mock out the ShadowUint16 method")
//			},
//			ShadowUint32Func: func(v uint32, uint32MoqParam types.Uint32)  {
//				panic("mock out the ShadowUint32 method")
//			},
//			ShadowUint64Func: func(v uint64, uint64MoqParam types.Uint64)  {
//				panic("mock out the ShadowUint64 method")
//			},
//			ShadowUint8Func: func(v uint8, uint8MoqParam types.Uint8)  {
//				panic("mock out the ShadowUint8 method")
//			},
//			ShadowUintptrFunc: func(v uintptr, uintptrMoqParam types.Uintptr)  {
//				panic("mock out the ShadowUintptr method")
//			},
//		}
//
//		// use mockedShadowTypes in code that requires ShadowTypes
//		// and then make assertions.
//
//	}

type ShadowTypesMock struct {
	// ShadowBoolFunc mocks the ShadowBool method.
	ShadowBoolFunc func(b bool, boolMoqParam types.Bool)

	// ShadowByteFunc mocks the ShadowByte method.
	ShadowByteFunc func(v byte, byteMoqParam types.Byte)

	// ShadowComplex128Func mocks the ShadowComplex128 method.
	ShadowComplex128Func func(v complex128, complex128MoqParam types.Complex128)

	// ShadowComplex64Func mocks the ShadowComplex64 method.
	ShadowComplex64Func func(v complex64, complex64MoqParam types.Complex64)

	// ShadowFloat32Func mocks the ShadowFloat32 method.
	ShadowFloat32Func func(f float32, float32MoqParam types.Float32)

	// ShadowFloat64Func mocks the ShadowFloat64 method.
	ShadowFloat64Func func(f float64, float64MoqParam types.Float64)

	// ShadowIntFunc mocks the ShadowInt method.
	ShadowIntFunc func(n int, intMoqParam types.Int)

	// ShadowInt16Func mocks the ShadowInt16 method.
	ShadowInt16Func func(n int16, int16MoqParam types.Int16)

	// ShadowInt32Func mocks the ShadowInt32 method.
	ShadowInt32Func func(n int32, int32MoqParam types.Int32)

	// ShadowInt64Func mocks the ShadowInt64 method.
	ShadowInt64Func func(n int64, int64MoqParam types.Int64)

	// ShadowInt8Func mocks the ShadowInt8 method.
	ShadowInt8Func func(n int8, int8MoqParam types.Int8)

	// ShadowRuneFunc mocks the ShadowRune method.
	ShadowRuneFunc func(n rune, runeMoqParam types.Rune)

	// ShadowStringFunc mocks the ShadowString method.
	ShadowStringFunc func(s string, stringMoqParam types.String)

	// ShadowUintFunc mocks the ShadowUint method.
	ShadowUintFunc func(v uint, uintMoqParam types.Uint)

	// ShadowUint16Func mocks the ShadowUint16 method.
	ShadowUint16Func func(v uint16, uint16MoqParam types.Uint16)

	// ShadowUint32Func mocks the ShadowUint32 method.
	ShadowUint32Func func(v uint32, uint32MoqParam types.Uint32)

	// ShadowUint64Func mocks the ShadowUint64 method.
	ShadowUint64Func func(v uint64, uint64MoqParam types.Uint64)

	// ShadowUint8Func mocks the ShadowUint8 method.
	ShadowUint8Func func(v uint8, uint8MoqParam types.Uint8)

	// ShadowUintptrFunc mocks the ShadowUintptr method.
	ShadowUintptrFunc func(v uintptr, uintptrMoqParam types.Uintptr)

	// calls tracks calls to the methods.
	calls struct {
		// ShadowBool holds details about calls to the ShadowBool method.
		ShadowBool []ShadowTypesMockShadowBoolCalls
		// ShadowByte holds details about calls to the ShadowByte method.
		ShadowByte []ShadowTypesMockShadowByteCalls
		// ShadowComplex128 holds details about calls to the ShadowComplex128 method.
		ShadowComplex128 []ShadowTypesMockShadowComplex128Calls
		// ShadowComplex64 holds details about calls to the ShadowComplex64 method.
		ShadowComplex64 []ShadowTypesMockShadowComplex64Calls
		// ShadowFloat32 holds details about calls to the ShadowFloat32 method.
		ShadowFloat32 []ShadowTypesMockShadowFloat32Calls
		// ShadowFloat64 holds details about calls to the ShadowFloat64 method.
		ShadowFloat64 []ShadowTypesMockShadowFloat64Calls
		// ShadowInt holds details about calls to the ShadowInt method.
		ShadowInt []ShadowTypesMockShadowIntCalls
		// ShadowInt16 holds details about calls to the ShadowInt16 method.
		ShadowInt16 []ShadowTypesMockShadowInt16Calls
		// ShadowInt32 holds details about calls to the ShadowInt32 method.
		ShadowInt32 []ShadowTypesMockShadowInt32Calls
		// ShadowInt64 holds details about calls to the ShadowInt64 method.
		ShadowInt64 []ShadowTypesMockShadowInt64Calls
		// ShadowInt8 holds details about calls to the ShadowInt8 method.
		ShadowInt8 []ShadowTypesMockShadowInt8Calls
		// ShadowRune holds details about calls to the ShadowRune method.
		ShadowRune []ShadowTypesMockShadowRuneCalls
		// ShadowString holds details about calls to the ShadowString method.
		ShadowString []ShadowTypesMockShadowStringCalls
		// ShadowUint holds details about calls to the ShadowUint method.
		ShadowUint []ShadowTypesMockShadowUintCalls
		// ShadowUint16 holds details about calls to the ShadowUint16 method.
		ShadowUint16 []ShadowTypesMockShadowUint16Calls
		// ShadowUint32 holds details about calls to the ShadowUint32 method.
		ShadowUint32 []ShadowTypesMockShadowUint32Calls
		// ShadowUint64 holds details about calls to the ShadowUint64 method.
		ShadowUint64 []ShadowTypesMockShadowUint64Calls
		// ShadowUint8 holds details about calls to the ShadowUint8 method.
		ShadowUint8 []ShadowTypesMockShadowUint8Calls
		// ShadowUintptr holds details about calls to the ShadowUintptr method.
		ShadowUintptr []ShadowTypesMockShadowUintptrCalls
	}
	lockShadowBool       sync.RWMutex
	lockShadowByte       sync.RWMutex
	lockShadowComplex128 sync.RWMutex
	lockShadowComplex64  sync.RWMutex
	lockShadowFloat32    sync.RWMutex
	lockShadowFloat64    sync.RWMutex
	lockShadowInt        sync.RWMutex
	lockShadowInt16      sync.RWMutex
	lockShadowInt32      sync.RWMutex
	lockShadowInt64      sync.RWMutex
	lockShadowInt8       sync.RWMutex
	lockShadowRune       sync.RWMutex
	lockShadowString     sync.RWMutex
	lockShadowUint       sync.RWMutex
	lockShadowUint16     sync.RWMutex
	lockShadowUint32     sync.RWMutex
	lockShadowUint64     sync.RWMutex
	lockShadowUint8      sync.RWMutex
	lockShadowUintptr    sync.RWMutex
}

// ShadowTypesMockShadowBoolCalls holds details about calls to the ShadowBool method.
type ShadowTypesMockShadowBoolCalls struct {
	// B is the b argument value.
	B bool
	// BoolMoqParam is the boolMoqParam argument value.
	BoolMoqParam types.Bool
}

// ShadowTypesMockShadowByteCalls holds details about calls to the ShadowByte method.
type ShadowTypesMockShadowByteCalls struct {
	// V is the v argument value.
	V byte
	// ByteMoqParam is the byteMoqParam argument value.
	ByteMoqParam types.Byte
}

// ShadowTypesMockShadowComplex128Calls holds details about calls to the ShadowComplex128 method.
type ShadowTypesMockShadowComplex128Calls struct {
	// V is the v argument value.
	V complex128
	// Complex128MoqParam is the complex128MoqParam argument value.
	Complex128MoqParam types.Complex128
}

// ShadowTypesMockShadowComplex64Calls holds details about calls to the ShadowComplex64 method.
type ShadowTypesMockShadowComplex64Calls struct {
	// V is the v argument value.
	V complex64
	// Complex64MoqParam is the complex64MoqParam argument value.
	Complex64MoqParam types.Complex64
}

// ShadowTypesMockShadowFloat32Calls holds details about calls to the ShadowFloat32 method.
type ShadowTypesMockShadowFloat32Calls struct {
	// F is the f argument value.
	F float32
	// Float32MoqParam is the float32MoqParam argument value.
	Float32MoqParam types.Float32
}

// ShadowTypesMockShadowFloat64Calls holds details about calls to the ShadowFloat64 method.
type ShadowTypesMockShadowFloat64Calls struct {
	// F is the f argument value.
	F float64
	// Float64MoqParam is the float64MoqParam argument value.
	Float64MoqParam types.Float64
}

// ShadowTypesMockShadowIntCalls holds details about calls to the ShadowInt method.
type ShadowTypesMockShadowIntCalls struct {
	// N is the n argument value.
	N int
	// IntMoqParam is the intMoqParam argument value.
	IntMoqParam types.Int
}

// ShadowTypesMockShadowInt16Calls holds details about calls to the ShadowInt16 method.
type ShadowTypesMockShadowInt16Calls struct {
	// N is the n argument value.
	N int16
	// Int16MoqParam is the int16MoqParam argument value.
	Int16MoqParam types.Int16
}

// ShadowTypesMockShadowInt32Calls holds details about calls to the ShadowInt32 method.
type ShadowTypesMockShadowInt32Calls struct {
	// N is the n argument value.
	N int32
	// Int32MoqParam is the int32MoqParam argument value.
	Int32MoqParam types.Int32
}

// ShadowTypesMockShadowInt64Calls holds details about calls to the ShadowInt64 method.
type ShadowTypesMockShadowInt64Calls struct {
	// N is the n argument value.
	N int64
	// Int64MoqParam is the int64MoqParam argument value.
	Int64MoqParam types.Int64
}

// ShadowTypesMockShadowInt8Calls holds details about calls to the ShadowInt8 method.
type ShadowTypesMockShadowInt8Calls struct {
	// N is the n argument value.
	N int8
	// Int8MoqParam is the int8MoqParam argument value.
	Int8MoqParam types.Int8
}

// ShadowTypesMockShadowRuneCalls holds details about calls to the ShadowRune method.
type ShadowTypesMockShadowRuneCalls struct {
	// N is the n argument value.
	N rune
	// RuneMoqParam is the runeMoqParam argument value.
	RuneMoqParam types.Rune
}

// ShadowTypesMockShadowStringCalls holds details about calls to the ShadowString method.
type ShadowTypesMockShadowStringCalls struct {
	// S is the s argument value.
	S string
	// StringMoqParam is the stringMoqParam argument value.
	StringMoqParam types.String
}

// ShadowTypesMockShadowUintCalls holds details about calls to the ShadowUint method.
type ShadowTypesMockShadowUintCalls struct {
	// V is the v argument value.
	V uint
	// UintMoqParam is the uintMoqParam argument value.
	UintMoqParam types.Uint
}

// ShadowTypesMockShadowUint16Calls holds details about calls to the ShadowUint16 method.
type ShadowTypesMockShadowUint16Calls struct {
	// V is the v argument value.
	V uint16
	// Uint16MoqParam is the uint16MoqParam argument value.
	Uint16MoqParam types.Uint16
}

// ShadowTypesMockShadowUint32Calls holds details about calls to the ShadowUint32 method.
type ShadowTypesMockShadowUint32Calls struct {
	// V is the v argument value.
	V uint32
	// Uint32MoqParam is the uint32MoqParam argument value.
	Uint32MoqParam types.Uint32
}

// ShadowTypesMockShadowUint64Calls holds details about calls to the ShadowUint64 method.
type ShadowTypesMockShadowUint64Calls struct {
	// V is the v argument value.
	V uint64
	// Uint64MoqParam is the uint64MoqParam argument value.
	Uint64MoqParam types.Uint64
}

// ShadowTypesMockShadowUint8Calls holds details about calls to the ShadowUint8 method.
type ShadowTypesMockShadowUint8Calls struct {
	// V is the v argument value.
	V uint8
	// Uint8MoqParam is the uint8MoqParam argument value.
	Uint8MoqParam types.Uint8
}

// ShadowTypesMockShadowUintptrCalls holds details about calls to the ShadowUintptr method.
type ShadowTypesMockShadowUintptrCalls struct {
	// V is the v argument value.
	V uintptr
	// UintptrMoqParam is the uintptrMoqParam argument value.
	UintptrMoqParam types.Uintptr
}

// ShadowBool calls ShadowBoolFunc.
func (mock *ShadowTypesMock) ShadowBool(b bool, boolMoqParam types.Bool) {
	if mock.ShadowBoolFunc == nil {
		panic("ShadowTypesMock.ShadowBoolFunc: method is nil but ShadowTypes.ShadowBool was just called")
	}
	callInfo := ShadowTypesMockShadowBoolCalls{
		B:            b,
		BoolMoqParam: boolMoqParam,
	}
	mock.lockShadowBool.Lock()
	mock.calls.ShadowBool = append(mock.calls.ShadowBool, callInfo)
	mock.lockShadowBool.Unlock()
	mock.ShadowBoolFunc(b, boolMoqParam)
}

// ShadowBoolCalls gets all the calls that were made to ShadowBool.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowBoolCalls())
func (mock *ShadowTypesMock) ShadowBoolCalls() []ShadowTypesMockShadowBoolCalls {
	var calls []ShadowTypesMockShadowBoolCalls
	mock.lockShadowBool.RLock()
	calls = mock.calls.ShadowBool
	mock.lockShadowBool.RUnlock()
	return calls
}

// ShadowByte calls ShadowByteFunc.
func (mock *ShadowTypesMock) ShadowByte(v byte, byteMoqParam types.Byte) {
	if mock.ShadowByteFunc == nil {
		panic("ShadowTypesMock.ShadowByteFunc: method is nil but ShadowTypes.ShadowByte was just called")
	}
	callInfo := ShadowTypesMockShadowByteCalls{
		V:            v,
		ByteMoqParam: byteMoqParam,
	}
	mock.lockShadowByte.Lock()
	mock.calls.ShadowByte = append(mock.calls.ShadowByte, callInfo)
	mock.lockShadowByte.Unlock()
	mock.ShadowByteFunc(v, byteMoqParam)
}

// ShadowByteCalls gets all the calls that were made to ShadowByte.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowByteCalls())
func (mock *ShadowTypesMock) ShadowByteCalls() []ShadowTypesMockShadowByteCalls {
	var calls []ShadowTypesMockShadowByteCalls
	mock.lockShadowByte.RLock()
	calls = mock.calls.ShadowByte
	mock.lockShadowByte.RUnlock()
	return calls
}

// ShadowComplex128 calls ShadowComplex128Func.
func (mock *ShadowTypesMock) ShadowComplex128(v complex128, complex128MoqParam types.Complex128) {
	if mock.ShadowComplex128Func == nil {
		panic("ShadowTypesMock.ShadowComplex128Func: method is nil but ShadowTypes.ShadowComplex128 was just called")
	}
	callInfo := ShadowTypesMockShadowComplex128Calls{
		V:                  v,
		Complex128MoqParam: complex128MoqParam,
	}
	mock.lockShadowComplex128.Lock()
	mock.calls.ShadowComplex128 = append(mock.calls.ShadowComplex128, callInfo)
	mock.lockShadowComplex128.Unlock()
	mock.ShadowComplex128Func(v, complex128MoqParam)
}

// ShadowComplex128Calls gets all the calls that were made to ShadowComplex128.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowComplex128Calls())
func (mock *ShadowTypesMock) ShadowComplex128Calls() []ShadowTypesMockShadowComplex128Calls {
	var calls []ShadowTypesMockShadowComplex128Calls
	mock.lockShadowComplex128.RLock()
	calls = mock.calls.ShadowComplex128
	mock.lockShadowComplex128.RUnlock()
	return calls
}

// ShadowComplex64 calls ShadowComplex64Func.
func (mock *ShadowTypesMock) ShadowComplex64(v complex64, complex64MoqParam types.Complex64) {
	if mock.ShadowComplex64Func == nil {
		panic("ShadowTypesMock.ShadowComplex64Func: method is nil but ShadowTypes.ShadowComplex64 was just called")
	}
	callInfo := ShadowTypesMockShadowComplex64Calls{
		V:                 v,
		Complex64MoqParam: complex64MoqParam,
	}
	mock.lockShadowComplex64.Lock()
	mock.calls.ShadowComplex64 = append(mock.calls.ShadowComplex64, callInfo)
	mock.lockShadowComplex64.Unlock()
	mock.ShadowComplex64Func(v, complex64MoqParam)
}

// ShadowComplex64Calls gets all the calls that were made to ShadowComplex64.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowComplex64Calls())
func (mock *ShadowTypesMock) ShadowComplex64Calls() []ShadowTypesMockShadowComplex64Calls {
	var calls []ShadowTypesMockShadowComplex64Calls
	mock.lockShadowComplex64.RLock()
	calls = mock.calls.ShadowComplex64
	mock.lockShadowComplex64.RUnlock()
	return calls
}

// ShadowFloat32 calls ShadowFloat32Func.
func (mock *ShadowTypesMock) ShadowFloat32(f float32, float32MoqParam types.Float32) {
	if mock.ShadowFloat32Func == nil {
		panic("ShadowTypesMock.ShadowFloat32Func: method is nil but ShadowTypes.ShadowFloat32 was just called")
	}
	callInfo := ShadowTypesMockShadowFloat32Calls{
		F:               f,
		Float32MoqParam: float32MoqParam,
	}
	mock.lockShadowFloat32.Lock()
	mock.calls.ShadowFloat32 = append(mock.calls.ShadowFloat32, callInfo)
	mock.lockShadowFloat32.Unlock()
	mock.ShadowFloat32Func(f, float32MoqParam)
}

// ShadowFloat32Calls gets all the calls that were made to ShadowFloat32.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowFloat32Calls())
func (mock *ShadowTypesMock) ShadowFloat32Calls() []ShadowTypesMockShadowFloat32Calls {
	var calls []ShadowTypesMockShadowFloat32Calls
	mock.lockShadowFloat32.RLock()
	calls = mock.calls.ShadowFloat32
	mock.lockShadowFloat32.RUnlock()
	return calls
}

// ShadowFloat64 calls ShadowFloat64Func.
func (mock *ShadowTypesMock) ShadowFloat64(f float64, float64MoqParam types.Float64) {
	if mock.ShadowFloat64Func == nil {
		panic("ShadowTypesMock.ShadowFloat64Func: method is nil but ShadowTypes.ShadowFloat64 was just called")
	}
	callInfo := ShadowTypesMockShadowFloat64Calls{
		F:               f,
		Float64MoqParam: float64MoqParam,
	}
	mock.lockShadowFloat64.Lock()
	mock.calls.ShadowFloat64 = append(mock.calls.ShadowFloat64, callInfo)
	mock.lockShadowFloat64.Unlock()
	mock.ShadowFloat64Func(f, float64MoqParam)
}

// ShadowFloat64Calls gets all the calls that were made to ShadowFloat64.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowFloat64Calls())
func (mock *ShadowTypesMock) ShadowFloat64Calls() []ShadowTypesMockShadowFloat64Calls {
	var calls []ShadowTypesMockShadowFloat64Calls
	mock.lockShadowFloat64.RLock()
	calls = mock.calls.ShadowFloat64
	mock.lockShadowFloat64.RUnlock()
	return calls
}

// ShadowInt calls ShadowIntFunc.
func (mock *ShadowTypesMock) ShadowInt(n int, intMoqParam types.Int) {
	if mock.ShadowIntFunc == nil {
		panic("ShadowTypesMock.ShadowIntFunc: method is nil but ShadowTypes.ShadowInt was just called")
	}
	callInfo := ShadowTypesMockShadowIntCalls{
		N:           n,
		IntMoqParam: intMoqParam,
	}
	mock.lockShadowInt.Lock()
	mock.calls.ShadowInt = append(mock.calls.ShadowInt, callInfo)
	mock.lockShadowInt.Unlock()
	mock.ShadowIntFunc(n, intMoqParam)
}

// ShadowIntCalls gets all the calls that were made to ShadowInt.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowIntCalls())
func (mock *ShadowTypesMock) ShadowIntCalls() []ShadowTypesMockShadowIntCalls {
	var calls []ShadowTypesMockShadowIntCalls
	mock.lockShadowInt.RLock()
	calls = mock.calls.ShadowInt
	mock.lockShadowInt.RUnlock()
	return calls
}

// ShadowInt16 calls ShadowInt16Func.
func (mock *ShadowTypesMock) ShadowInt16(n int16, int16MoqParam types.Int16) {
	if mock.ShadowInt16Func == nil {
		panic("ShadowTypesMock.ShadowInt16Func: method is nil but ShadowTypes.ShadowInt16 was just called")
	}
	callInfo := ShadowTypesMockShadowInt16Calls{
		N:             n,
		Int16MoqParam: int16MoqParam,
	}
	mock.lockShadowInt16.Lock()
	mock.calls.ShadowInt16 = append(mock.calls.ShadowInt16, callInfo)
	mock.lockShadowInt16.Unlock()
	mock.ShadowInt16Func(n, int16MoqParam)
}

// ShadowInt16Calls gets all the calls that were made to ShadowInt16.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowInt16Calls())
func (mock *ShadowTypesMock) ShadowInt16Calls() []ShadowTypesMockShadowInt16Calls {
	var calls []ShadowTypesMockShadowInt16Calls
	mock.lockShadowInt16.RLock()
	calls = mock.calls.ShadowInt16
	mock.lockShadowInt16.RUnlock()
	return calls
}

// ShadowInt32 calls ShadowInt32Func.
func (mock *ShadowTypesMock) ShadowInt32(n int32, int32MoqParam types.Int32) {
	if mock.ShadowInt32Func == nil {
		panic("ShadowTypesMock.ShadowInt32Func: method is nil but ShadowTypes.ShadowInt32 was just called")
	}
	callInfo := ShadowTypesMockShadowInt32Calls{
		N:             n,
		Int32MoqParam: int32MoqParam,
	}
	mock.lockShadowInt32.Lock()
	mock.calls.ShadowInt32 = append(mock.calls.ShadowInt32, callInfo)
	mock.lockShadowInt32.Unlock()
	mock.ShadowInt32Func(n, int32MoqParam)
}

// ShadowInt32Calls gets all the calls that were made to ShadowInt32.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowInt32Calls())
func (mock *ShadowTypesMock) ShadowInt32Calls() []ShadowTypesMockShadowInt32Calls {
	var calls []ShadowTypesMockShadowInt32Calls
	mock.lockShadowInt32.RLock()
	calls = mock.calls.ShadowInt32
	mock.lockShadowInt32.RUnlock()
	return calls
}

// ShadowInt64 calls ShadowInt64Func.
func (mock *ShadowTypesMock) ShadowInt64(n int64, int64MoqParam types.Int64) {
	if mock.ShadowInt64Func == nil {
		panic("ShadowTypesMock.ShadowInt64Func: method is nil but ShadowTypes.ShadowInt64 was just called")
	}
	callInfo := ShadowTypesMockShadowInt64Calls{
		N:             n,
		Int64MoqParam: int64MoqParam,
	}
	mock.lockShadowInt64.Lock()
	mock.calls.ShadowInt64 = append(mock.calls.ShadowInt64, callInfo)
	mock.lockShadowInt64.Unlock()
	mock.ShadowInt64Func(n, int64MoqParam)
}

// ShadowInt64Calls gets all the calls that were made to ShadowInt64.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowInt64Calls())
func (mock *ShadowTypesMock) ShadowInt64Calls() []ShadowTypesMockShadowInt64Calls {
	var calls []ShadowTypesMockShadowInt64Calls
	mock.lockShadowInt64.RLock()
	calls = mock.calls.ShadowInt64
	mock.lockShadowInt64.RUnlock()
	return calls
}

// ShadowInt8 calls ShadowInt8Func.
func (mock *ShadowTypesMock) ShadowInt8(n int8, int8MoqParam types.Int8) {
	if mock.ShadowInt8Func == nil {
		panic("ShadowTypesMock.ShadowInt8Func: method is nil but ShadowTypes.ShadowInt8 was just called")
	}
	callInfo := ShadowTypesMockShadowInt8Calls{
		N:            n,
		Int8MoqParam: int8MoqParam,
	}
	mock.lockShadowInt8.Lock()
	mock.calls.ShadowInt8 = append(mock.calls.ShadowInt8, callInfo)
	mock.lockShadowInt8.Unlock()
	mock.ShadowInt8Func(n, int8MoqParam)
}

// ShadowInt8Calls gets all the calls that were made to ShadowInt8.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowInt8Calls())
func (mock *ShadowTypesMock) ShadowInt8Calls() []ShadowTypesMockShadowInt8Calls {
	var calls []ShadowTypesMockShadowInt8Calls
	mock.lockShadowInt8.RLock()
	calls = mock.calls.ShadowInt8
	mock.lockShadowInt8.RUnlock()
	return calls
}

// ShadowRune calls ShadowRuneFunc.
func (mock *ShadowTypesMock) ShadowRune(n rune, runeMoqParam types.Rune) {
	if mock.ShadowRuneFunc == nil {
		panic("ShadowTypesMock.ShadowRuneFunc: method is nil but ShadowTypes.ShadowRune was just called")
	}
	callInfo := ShadowTypesMockShadowRuneCalls{
		N:            n,
		RuneMoqParam: runeMoqParam,
	}
	mock.lockShadowRune.Lock()
	mock.calls.ShadowRune = append(mock.calls.ShadowRune, callInfo)
	mock.lockShadowRune.Unlock()
	mock.ShadowRuneFunc(n, runeMoqParam)
}

// ShadowRuneCalls gets all the calls that were made to ShadowRune.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowRuneCalls())
func (mock *ShadowTypesMock) ShadowRuneCalls() []ShadowTypesMockShadowRuneCalls {
	var calls []ShadowTypesMockShadowRuneCalls
	mock.lockShadowRune.RLock()
	calls = mock.calls.ShadowRune
	mock.lockShadowRune.RUnlock()
	return calls
}

// ShadowString calls ShadowStringFunc.
func (mock *ShadowTypesMock) ShadowString(s string, stringMoqParam types.String) {
	if mock.ShadowStringFunc == nil {
		panic("ShadowTypesMock.ShadowStringFunc: method is nil but ShadowTypes.ShadowString was just called")
	}
	callInfo := ShadowTypesMockShadowStringCalls{
		S:              s,
		StringMoqParam: stringMoqParam,
	}
	mock.lockShadowString.Lock()
	mock.calls.ShadowString = append(mock.calls.ShadowString, callInfo)
	mock.lockShadowString.Unlock()
	mock.ShadowStringFunc(s, stringMoqParam)
}

// ShadowStringCalls gets all the calls that were made to ShadowString.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowStringCalls())
func (mock *ShadowTypesMock) ShadowStringCalls() []ShadowTypesMockShadowStringCalls {
	var calls []ShadowTypesMockShadowStringCalls
	mock.lockShadowString.RLock()
	calls = mock.calls.ShadowString
	mock.lockShadowString.RUnlock()
	return calls
}

// ShadowUint calls ShadowUintFunc.
func (mock *ShadowTypesMock) ShadowUint(v uint, uintMoqParam types.Uint) {
	if mock.ShadowUintFunc == nil {
		panic("ShadowTypesMock.ShadowUintFunc: method is nil but ShadowTypes.ShadowUint was just called")
	}
	callInfo := ShadowTypesMockShadowUintCalls{
		V:            v,
		UintMoqParam: uintMoqParam,
	}
	mock.lockShadowUint.Lock()
	mock.calls.ShadowUint = append(mock.calls.ShadowUint, callInfo)
	mock.lockShadowUint.Unlock()
	mock.ShadowUintFunc(v, uintMoqParam)
}

// ShadowUintCalls gets all the calls that were made to ShadowUint.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowUintCalls())
func (mock *ShadowTypesMock) ShadowUintCalls() []ShadowTypesMockShadowUintCalls {
	var calls []ShadowTypesMockShadowUintCalls
	mock.lockShadowUint.RLock()
	calls = mock.calls.ShadowUint
	mock.lockShadowUint.RUnlock()
	return calls
}

// ShadowUint16 calls ShadowUint16Func.
func (mock *ShadowTypesMock) ShadowUint16(v uint16, uint16MoqParam types.Uint16) {
	if mock.ShadowUint16Func == nil {
		panic("ShadowTypesMock.ShadowUint16Func: method is nil but ShadowTypes.ShadowUint16 was just called")
	}
	callInfo := ShadowTypesMockShadowUint16Calls{
		V:              v,
		Uint16MoqParam: uint16MoqParam,
	}
	mock.lockShadowUint16.Lock()
	mock.calls.ShadowUint16 = append(mock.calls.ShadowUint16, callInfo)
	mock.lockShadowUint16.Unlock()
	mock.ShadowUint16Func(v, uint16MoqParam)
}

// ShadowUint16Calls gets all the calls that were made to ShadowUint16.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowUint16Calls())
func (mock *ShadowTypesMock) ShadowUint16Calls() []ShadowTypesMockShadowUint16Calls {
	var calls []ShadowTypesMockShadowUint16Calls
	mock.lockShadowUint16.RLock()
	calls = mock.calls.ShadowUint16
	mock.lockShadowUint16.RUnlock()
	return calls
}

// ShadowUint32 calls ShadowUint32Func.
func (mock *ShadowTypesMock) ShadowUint32(v uint32, uint32MoqParam types.Uint32) {
	if mock.ShadowUint32Func == nil {
		panic("ShadowTypesMock.ShadowUint32Func: method is nil but ShadowTypes.ShadowUint32 was just called")
	}
	callInfo := ShadowTypesMockShadowUint32Calls{
		V:              v,
		Uint32MoqParam: uint32MoqParam,
	}
	mock.lockShadowUint32.Lock()
	mock.calls.ShadowUint32 = append(mock.calls.ShadowUint32, callInfo)
	mock.lockShadowUint32.Unlock()
	mock.ShadowUint32Func(v, uint32MoqParam)
}

// ShadowUint32Calls gets all the calls that were made to ShadowUint32.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowUint32Calls())
func (mock *ShadowTypesMock) ShadowUint32Calls() []ShadowTypesMockShadowUint32Calls {
	var calls []ShadowTypesMockShadowUint32Calls
	mock.lockShadowUint32.RLock()
	calls = mock.calls.ShadowUint32
	mock.lockShadowUint32.RUnlock()
	return calls
}

// ShadowUint64 calls ShadowUint64Func.
func (mock *ShadowTypesMock) ShadowUint64(v uint64, uint64MoqParam types.Uint64) {
	if mock.ShadowUint64Func == nil {
		panic("ShadowTypesMock.ShadowUint64Func: method is nil but ShadowTypes.ShadowUint64 was just called")
	}
	callInfo := ShadowTypesMockShadowUint64Calls{
		V:              v,
		Uint64MoqParam: uint64MoqParam,
	}
	mock.lockShadowUint64.Lock()
	mock.calls.ShadowUint64 = append(mock.calls.ShadowUint64, callInfo)
	mock.lockShadowUint64.Unlock()
	mock.ShadowUint64Func(v, uint64MoqParam)
}

// ShadowUint64Calls gets all the calls that were made to ShadowUint64.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowUint64Calls())
func (mock *ShadowTypesMock) ShadowUint64Calls() []ShadowTypesMockShadowUint64Calls {
	var calls []ShadowTypesMockShadowUint64Calls
	mock.lockShadowUint64.RLock()
	calls = mock.calls.ShadowUint64
	mock.lockShadowUint64.RUnlock()
	return calls
}

// ShadowUint8 calls ShadowUint8Func.
func (mock *ShadowTypesMock) ShadowUint8(v uint8, uint8MoqParam types.Uint8) {
	if mock.ShadowUint8Func == nil {
		panic("ShadowTypesMock.ShadowUint8Func: method is nil but ShadowTypes.ShadowUint8 was just called")
	}
	callInfo := ShadowTypesMockShadowUint8Calls{
		V:             v,
		Uint8MoqParam: uint8MoqParam,
	}
	mock.lockShadowUint8.Lock()
	mock.calls.ShadowUint8 = append(mock.calls.ShadowUint8, callInfo)
	mock.lockShadowUint8.Unlock()
	mock.ShadowUint8Func(v, uint8MoqParam)
}

// ShadowUint8Calls gets all the calls that were made to ShadowUint8.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowUint8Calls())
func (mock *ShadowTypesMock) ShadowUint8Calls() []ShadowTypesMockShadowUint8Calls {
	var calls []ShadowTypesMockShadowUint8Calls
	mock.lockShadowUint8.RLock()
	calls = mock.calls.ShadowUint8
	mock.lockShadowUint8.RUnlock()
	return calls
}

// ShadowUintptr calls ShadowUintptrFunc.
func (mock *ShadowTypesMock) ShadowUintptr(v uintptr, uintptrMoqParam types.Uintptr) {
	if mock.ShadowUintptrFunc == nil {
		panic("ShadowTypesMock.ShadowUintptrFunc: method is nil but ShadowTypes.ShadowUintptr was just called")
	}
	callInfo := ShadowTypesMockShadowUintptrCalls{
		V:               v,
		UintptrMoqParam: uintptrMoqParam,
	}
	mock.lockShadowUintptr.Lock()
	mock.calls.ShadowUintptr = append(mock.calls.ShadowUintptr, callInfo)
	mock.lockShadowUintptr.Unlock()
	mock.ShadowUintptrFunc(v, uintptrMoqParam)
}

// ShadowUintptrCalls gets all the calls that were made to ShadowUintptr.
// Check the length with:
//
//	len(mockedShadowTypes.ShadowUintptrCalls())
func (mock *ShadowTypesMock) ShadowUintptrCalls() []ShadowTypesMockShadowUintptrCalls {
	var calls []ShadowTypesMockShadowUintptrCalls
	mock.lockShadowUintptr.RLock()
	calls = mock.calls.ShadowUintptr
	mock.lockShadowUintptr.RUnlock()
	return calls
}
