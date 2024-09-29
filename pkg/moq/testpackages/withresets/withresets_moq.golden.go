// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package withresets

import (
	"context"
	"sync"
)

// Ensure, that ResetStoreMock does implement ResetStore.
// If this is not the case, regenerate this file with moq.
var _ ResetStore = &ResetStoreMock{}

// ResetStoreMock is a mock implementation of ResetStore.
//
//	func TestSomethingThatUsesResetStore(t *testing.T) {
//
//		// make and configure a mocked ResetStore
//		mockedResetStore := &ResetStoreMock{
//			ClearCacheFunc: func(id string)  {
//				panic("mock out the ClearCache method")
//			},
//			CreateFunc: func(ctx context.Context, person *Reset, confirm bool) error {
//				panic("mock out the Create method")
//			},
//			GetFunc: func(ctx context.Context, id string) (*Reset, error) {
//				panic("mock out the Get method")
//			},
//		}
//
//		// use mockedResetStore in code that requires ResetStore
//		// and then make assertions.
//
//	}

type ResetStoreMock struct {
	// ClearCacheFunc mocks the ClearCache method.
	ClearCacheFunc func(id string)

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, person *Reset, confirm bool) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id string) (*Reset, error)

	// calls tracks calls to the methods.
	calls struct {
		// ClearCache holds details about calls to the ClearCache method.
		ClearCache []ResetStoreMockClearCacheCalls
		// Create holds details about calls to the Create method.
		Create []ResetStoreMockCreateCalls
		// Get holds details about calls to the Get method.
		Get []ResetStoreMockGetCalls
	}
	lockClearCache sync.RWMutex
	lockCreate     sync.RWMutex
	lockGet        sync.RWMutex
}

// ResetStoreMockClearCacheCalls holds details about calls to the ClearCache method.
type ResetStoreMockClearCacheCalls struct {
	// ID is the id argument value.
	ID string
}

// ResetStoreMockCreateCalls holds details about calls to the Create method.
type ResetStoreMockCreateCalls struct {
	// Ctx is the ctx argument value.
	Ctx context.Context
	// Person is the person argument value.
	Person *Reset
	// Confirm is the confirm argument value.
	Confirm bool
}

// ResetStoreMockGetCalls holds details about calls to the Get method.
type ResetStoreMockGetCalls struct {
	// Ctx is the ctx argument value.
	Ctx context.Context
	// ID is the id argument value.
	ID string
}

// ClearCache calls ClearCacheFunc.
func (mock *ResetStoreMock) ClearCache(id string) {
	if mock.ClearCacheFunc == nil {
		panic("ResetStoreMock.ClearCacheFunc: method is nil but ResetStore.ClearCache was just called")
	}
	callInfo := ResetStoreMockClearCacheCalls{
		ID: id,
	}
	mock.lockClearCache.Lock()
	mock.calls.ClearCache = append(mock.calls.ClearCache, callInfo)
	mock.lockClearCache.Unlock()
	mock.ClearCacheFunc(id)
}

// ClearCacheCalls gets all the calls that were made to ClearCache.
// Check the length with:
//
//	len(mockedResetStore.ClearCacheCalls())
func (mock *ResetStoreMock) ClearCacheCalls() []ResetStoreMockClearCacheCalls {
	var calls []ResetStoreMockClearCacheCalls
	mock.lockClearCache.RLock()
	calls = mock.calls.ClearCache
	mock.lockClearCache.RUnlock()
	return calls
}

// ResetClearCacheCalls reset all the calls that were made to ClearCache.
func (mock *ResetStoreMock) ResetClearCacheCalls() {
	mock.lockClearCache.Lock()
	mock.calls.ClearCache = nil
	mock.lockClearCache.Unlock()
}

// Create calls CreateFunc.
func (mock *ResetStoreMock) Create(ctx context.Context, person *Reset, confirm bool) error {
	if mock.CreateFunc == nil {
		panic("ResetStoreMock.CreateFunc: method is nil but ResetStore.Create was just called")
	}
	callInfo := ResetStoreMockCreateCalls{
		Ctx:     ctx,
		Person:  person,
		Confirm: confirm,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, person, confirm)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedResetStore.CreateCalls())
func (mock *ResetStoreMock) CreateCalls() []ResetStoreMockCreateCalls {
	var calls []ResetStoreMockCreateCalls
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// ResetCreateCalls reset all the calls that were made to Create.
func (mock *ResetStoreMock) ResetCreateCalls() {
	mock.lockCreate.Lock()
	mock.calls.Create = nil
	mock.lockCreate.Unlock()
}

// Get calls GetFunc.
func (mock *ResetStoreMock) Get(ctx context.Context, id string) (*Reset, error) {
	if mock.GetFunc == nil {
		panic("ResetStoreMock.GetFunc: method is nil but ResetStore.Get was just called")
	}
	callInfo := ResetStoreMockGetCalls{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedResetStore.GetCalls())
func (mock *ResetStoreMock) GetCalls() []ResetStoreMockGetCalls {
	var calls []ResetStoreMockGetCalls
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// ResetGetCalls reset all the calls that were made to Get.
func (mock *ResetStoreMock) ResetGetCalls() {
	mock.lockGet.Lock()
	mock.calls.Get = nil
	mock.lockGet.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *ResetStoreMock) ResetCalls() {
	mock.lockClearCache.Lock()
	mock.calls.ClearCache = nil
	mock.lockClearCache.Unlock()

	mock.lockCreate.Lock()
	mock.calls.Create = nil
	mock.lockCreate.Unlock()

	mock.lockGet.Lock()
	mock.calls.Get = nil
	mock.lockGet.Unlock()
}
