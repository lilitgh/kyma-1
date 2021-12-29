// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	rafter "github.com/kyma-project/kyma/components/console-backend-service/internal/domain/rafter"
	v1beta1 "github.com/kyma-project/rafter/pkg/apis/rafter/v1beta1"
	mock "github.com/stretchr/testify/mock"
)

// fileSvc is an autogenerated mock type for the fileSvc type
type fileSvc struct {
	mock.Mock
}

// Extract provides a mock function with given fields: statusRef
func (_m *fileSvc) Extract(statusRef *v1beta1.AssetStatusRef) ([]*rafter.File, error) {
	ret := _m.Called(statusRef)

	var r0 []*rafter.File
	if rf, ok := ret.Get(0).(func(*v1beta1.AssetStatusRef) []*rafter.File); ok {
		r0 = rf(statusRef)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*rafter.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.AssetStatusRef) error); ok {
		r1 = rf(statusRef)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterByExtensionsAndExtract provides a mock function with given fields: statusRef, filterExtensions
func (_m *fileSvc) FilterByExtensionsAndExtract(statusRef *v1beta1.AssetStatusRef, filterExtensions []string) ([]*rafter.File, error) {
	ret := _m.Called(statusRef, filterExtensions)

	var r0 []*rafter.File
	if rf, ok := ret.Get(0).(func(*v1beta1.AssetStatusRef, []string) []*rafter.File); ok {
		r0 = rf(statusRef, filterExtensions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*rafter.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.AssetStatusRef, []string) error); ok {
		r1 = rf(statusRef, filterExtensions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}