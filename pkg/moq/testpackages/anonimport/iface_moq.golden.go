// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package anonimport

import (
	"context"
	"sync"
)

// Ensure, that ExampleMock does implement Example.
// If this is not the case, regenerate this file with moq.
var _ Example = &ExampleMock{}

// ExampleMock is a mock implementation of Example.
//
// 	func TestSomethingThatUsesExample(t *testing.T) {
//
// 		// make and configure a mocked Example
// 		mockedExample := &ExampleMock{
// 			CtxFunc: func(ctx context.Context)  {
// 				panic("mock out the Ctx method")
// 			},
// 		}
//
// 		// use mockedExample in code that requires Example
// 		// and then make assertions.
//
// 	}
type ExampleMock struct {
	// CtxFunc mocks the Ctx method.
	CtxFunc func(ctx context.Context)

	// calls tracks calls to the methods.
	calls struct {
		// Ctx holds details about calls to the Ctx method.
		Ctx []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockCtx sync.RWMutex
}

// Ctx calls CtxFunc.
func (mock *ExampleMock) Ctx(ctx context.Context) {
	if mock.CtxFunc == nil {
		panic("ExampleMock.CtxFunc: method is nil but Example.Ctx was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCtx.Lock()
	mock.calls.Ctx = append(mock.calls.Ctx, callInfo)
	mock.lockCtx.Unlock()
	mock.CtxFunc(ctx)
}

// CtxCalls gets all the calls that were made to Ctx.
// Check the length with:
//     len(mockedExample.CtxCalls())
func (mock *ExampleMock) CtxCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCtx.RLock()
	calls = mock.calls.Ctx
	mock.lockCtx.RUnlock()
	return calls
}
