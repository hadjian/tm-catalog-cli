// Code generated by mockery v2.39.1. DO NOT EDIT.

package http

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/web-of-things-open-source/tm-catalog-cli/internal/model"
)

// MockHandlerService is an autogenerated mock type for the HandlerService type
type MockHandlerService struct {
	mock.Mock
}

// CheckHealth provides a mock function with given fields: ctx
func (_m *MockHandlerService) CheckHealth(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckHealth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckHealthLive provides a mock function with given fields: ctx
func (_m *MockHandlerService) CheckHealthLive(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckHealthLive")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckHealthReady provides a mock function with given fields: ctx
func (_m *MockHandlerService) CheckHealthReady(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckHealthReady")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckHealthStartup provides a mock function with given fields: ctx
func (_m *MockHandlerService) CheckHealthStartup(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckHealthStartup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchThingModel provides a mock function with given fields: ctx, tmID
func (_m *MockHandlerService) FetchThingModel(ctx context.Context, tmID string) ([]byte, error) {
	ret := _m.Called(ctx, tmID)

	if len(ret) == 0 {
		panic("no return value specified for FetchThingModel")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, tmID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, tmID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tmID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindInventoryEntry provides a mock function with given fields: ctx, name
func (_m *MockHandlerService) FindInventoryEntry(ctx context.Context, name string) (*model.FoundEntry, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for FindInventoryEntry")
	}

	var r0 *model.FoundEntry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.FoundEntry, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.FoundEntry); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FoundEntry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAuthors provides a mock function with given fields: ctx, search
func (_m *MockHandlerService) ListAuthors(ctx context.Context, search *model.SearchParams) ([]string, error) {
	ret := _m.Called(ctx, search)

	if len(ret) == 0 {
		panic("no return value specified for ListAuthors")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) ([]string, error)); ok {
		return rf(ctx, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) []string); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.SearchParams) error); ok {
		r1 = rf(ctx, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInventory provides a mock function with given fields: ctx, search
func (_m *MockHandlerService) ListInventory(ctx context.Context, search *model.SearchParams) (*model.SearchResult, error) {
	ret := _m.Called(ctx, search)

	if len(ret) == 0 {
		panic("no return value specified for ListInventory")
	}

	var r0 *model.SearchResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) (*model.SearchResult, error)); ok {
		return rf(ctx, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) *model.SearchResult); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SearchResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.SearchParams) error); ok {
		r1 = rf(ctx, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManufacturers provides a mock function with given fields: ctx, search
func (_m *MockHandlerService) ListManufacturers(ctx context.Context, search *model.SearchParams) ([]string, error) {
	ret := _m.Called(ctx, search)

	if len(ret) == 0 {
		panic("no return value specified for ListManufacturers")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) ([]string, error)); ok {
		return rf(ctx, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) []string); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.SearchParams) error); ok {
		r1 = rf(ctx, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMpns provides a mock function with given fields: ctx, search
func (_m *MockHandlerService) ListMpns(ctx context.Context, search *model.SearchParams) ([]string, error) {
	ret := _m.Called(ctx, search)

	if len(ret) == 0 {
		panic("no return value specified for ListMpns")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) ([]string, error)); ok {
		return rf(ctx, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) []string); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.SearchParams) error); ok {
		r1 = rf(ctx, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PushThingModel provides a mock function with given fields: ctx, file
func (_m *MockHandlerService) PushThingModel(ctx context.Context, file []byte) (string, error) {
	ret := _m.Called(ctx, file)

	if len(ret) == 0 {
		panic("no return value specified for PushThingModel")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) (string, error)); ok {
		return rf(ctx, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) string); ok {
		r0 = rf(ctx, file)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockHandlerService creates a new instance of MockHandlerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandlerService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandlerService {
	mock := &MockHandlerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
