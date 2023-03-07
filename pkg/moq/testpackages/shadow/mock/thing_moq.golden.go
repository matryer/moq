// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	shadowhttp "github.com/matryer/moq/pkg/moq/testpackages/shadow/http"
	nethttp "net/http"
	"sync"
)

// Ensure, that ThingMock does implement shadowhttp.Thing.
// If this is not the case, regenerate this file with moq.
var _ shadowhttp.Thing = &ThingMock{}

// ThingMock is a mock implementation of shadowhttp.Thing.
//
//	func TestSomethingThatUsesThing(t *testing.T) {
//
//		// make and configure a mocked shadowhttp.Thing
//		mockedThing := &ThingMock{
//			BlahFunc: func(w nethttp.ResponseWriter, r *nethttp.Request)  {
//				panic("mock out the Blah method")
//			},
//		}
//
//		// use mockedThing in code that requires shadowhttp.Thing
//		// and then make assertions.
//
//	}
type ThingMock struct {
	// BlahFunc mocks the Blah method.
	BlahFunc func(w nethttp.ResponseWriter, r *nethttp.Request)

	// calls tracks calls to the methods.
	calls struct {
		// Blah holds details about calls to the Blah method.
		Blah []struct {
			// W is the w argument value.
			W nethttp.ResponseWriter
			// R is the r argument value.
			R *nethttp.Request
		}
	}
	lockBlah sync.RWMutex
}

// Blah calls BlahFunc.
func (mock *ThingMock) Blah(w nethttp.ResponseWriter, r *nethttp.Request) {
	if mock.BlahFunc == nil {
		panic("ThingMock.BlahFunc: method is nil but Thing.Blah was just called")
	}
	callInfo := struct {
		W nethttp.ResponseWriter
		R *nethttp.Request
	}{
		W: w,
		R: r,
	}
	mock.lockBlah.Lock()
	mock.calls.Blah = append(mock.calls.Blah, callInfo)
	mock.lockBlah.Unlock()
	mock.BlahFunc(w, r)
}

// BlahCalls gets all the calls that were made to Blah.
// Check the length with:
//
//	len(mockedThing.BlahCalls())
func (mock *ThingMock) BlahCalls() []struct {
	W nethttp.ResponseWriter
	R *nethttp.Request
} {
	var calls []struct {
		W nethttp.ResponseWriter
		R *nethttp.Request
	}
	mock.lockBlah.RLock()
	calls = mock.calls.Blah
	mock.lockBlah.RUnlock()
	return calls
}
