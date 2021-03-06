// Code generated by mockery v1.0.0

// CLOSE THIS FILE AS QUICKLY AS POSSIBLE

package mocks

import context "context"
import functions "github.com/vmware/dispatch/pkg/functions"
import mock "github.com/stretchr/testify/mock"

// ImageBuilder is an autogenerated mock type for the ImageBuilder type
type ImageBuilder struct {
	mock.Mock
}

// BuildImage provides a mock function with given fields: ctx, f, code
func (_m *ImageBuilder) BuildImage(ctx context.Context, f *functions.Function, code []byte) (string, error) {
	ret := _m.Called(ctx, f, code)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *functions.Function, []byte) string); ok {
		r0 = rf(ctx, f, code)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *functions.Function, []byte) error); ok {
		r1 = rf(ctx, f, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveImage provides a mock function with given fields: ctx, f
func (_m *ImageBuilder) RemoveImage(ctx context.Context, f *functions.Function) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *functions.Function) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
