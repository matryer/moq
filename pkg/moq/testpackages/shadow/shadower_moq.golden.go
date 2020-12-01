// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package shadow

import (
	"io"
	"sync"
)

// Ensure, that ShadowerMock does implement Shadower.
// If this is not the case, regenerate this file with moq.
var _ Shadower = &ShadowerMock{}

// ShadowerMock is a mock implementation of Shadower.
//
//     func TestSomethingThatUsesShadower(t *testing.T) {
//
//         // make and configure a mocked Shadower
//         mockedShadower := &ShadowerMock{
//             ShadowFunc: func(ioMoqParam io.Reader)  {
// 	               panic("mock out the Shadow method")
//             },
//             ShadowTwoFunc: func(r io.Reader, ioMoqParam interface{})  {
// 	               panic("mock out the ShadowTwo method")
//             },
//         }
//
//         // use mockedShadower in code that requires Shadower
//         // and then make assertions.
//
//     }
type ShadowerMock struct {
	// ShadowFunc mocks the Shadow method.
	ShadowFunc func(ioMoqParam io.Reader)

	// ShadowTwoFunc mocks the ShadowTwo method.
	ShadowTwoFunc func(r io.Reader, ioMoqParam interface{})

	// calls tracks calls to the methods.
	calls struct {
		// Shadow holds details about calls to the Shadow method.
		Shadow []struct {
			// IoMoqParam is the ioMoqParam argument value.
			IoMoqParam io.Reader
		}
		// ShadowTwo holds details about calls to the ShadowTwo method.
		ShadowTwo []struct {
			// R is the r argument value.
			R io.Reader
			// IoMoqParam is the ioMoqParam argument value.
			IoMoqParam interface{}
		}
	}
	lockShadow    sync.RWMutex
	lockShadowTwo sync.RWMutex
}

// Shadow calls ShadowFunc.
func (mock *ShadowerMock) Shadow(ioMoqParam io.Reader) {
	if mock.ShadowFunc == nil {
		panic("ShadowerMock.ShadowFunc: method is nil but Shadower.Shadow was just called")
	}
	callInfo := struct {
		IoMoqParam io.Reader
	}{
		IoMoqParam: ioMoqParam,
	}
	mock.lockShadow.Lock()
	mock.calls.Shadow = append(mock.calls.Shadow, callInfo)
	mock.lockShadow.Unlock()
	mock.ShadowFunc(ioMoqParam)
}

// ShadowCalls gets all the calls that were made to Shadow.
// Check the length with:
//     len(mockedShadower.ShadowCalls())
func (mock *ShadowerMock) ShadowCalls() []struct {
	IoMoqParam io.Reader
} {
	var calls []struct {
		IoMoqParam io.Reader
	}
	mock.lockShadow.RLock()
	calls = mock.calls.Shadow
	mock.lockShadow.RUnlock()
	return calls
}

// ShadowTwo calls ShadowTwoFunc.
func (mock *ShadowerMock) ShadowTwo(r io.Reader, ioMoqParam interface{}) {
	if mock.ShadowTwoFunc == nil {
		panic("ShadowerMock.ShadowTwoFunc: method is nil but Shadower.ShadowTwo was just called")
	}
	callInfo := struct {
		R          io.Reader
		IoMoqParam interface{}
	}{
		R:          r,
		IoMoqParam: ioMoqParam,
	}
	mock.lockShadowTwo.Lock()
	mock.calls.ShadowTwo = append(mock.calls.ShadowTwo, callInfo)
	mock.lockShadowTwo.Unlock()
	mock.ShadowTwoFunc(r, ioMoqParam)
}

// ShadowTwoCalls gets all the calls that were made to ShadowTwo.
// Check the length with:
//     len(mockedShadower.ShadowTwoCalls())
func (mock *ShadowerMock) ShadowTwoCalls() []struct {
	R          io.Reader
	IoMoqParam interface{}
} {
	var calls []struct {
		R          io.Reader
		IoMoqParam interface{}
	}
	mock.lockShadowTwo.RLock()
	calls = mock.calls.ShadowTwo
	mock.lockShadowTwo.RUnlock()
	return calls
}
