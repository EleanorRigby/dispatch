// Code generated by mockery v1.0.0

package mocks

import image "github.com/vmware/dispatch/pkg/image-manager/gen/client/image"
import mock "github.com/stretchr/testify/mock"
import runtime "github.com/go-openapi/runtime"

// ImageManager is an autogenerated mock type for the ImageManager type
type ImageManager struct {
	mock.Mock
}

// GetImageByName provides a mock function with given fields: _a0, _a1
func (_m *ImageManager) GetImageByName(_a0 *image.GetImageByNameParams, _a1 runtime.ClientAuthInfoWriter) (*image.GetImageByNameOK, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *image.GetImageByNameOK
	if rf, ok := ret.Get(0).(func(*image.GetImageByNameParams, runtime.ClientAuthInfoWriter) *image.GetImageByNameOK); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*image.GetImageByNameOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*image.GetImageByNameParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
