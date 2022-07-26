// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.14.0. DO NOT EDIT.

package awslib

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockIamRolePermissionGetter is an autogenerated mock type for the IamRolePermissionGetter type
type MockIamRolePermissionGetter struct {
	mock.Mock
}

type MockIamRolePermissionGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIamRolePermissionGetter) EXPECT() *MockIamRolePermissionGetter_Expecter {
	return &MockIamRolePermissionGetter_Expecter{mock: &_m.Mock}
}

// GetIAMRolePermissions provides a mock function with given fields: ctx, roleName
func (_m *MockIamRolePermissionGetter) GetIAMRolePermissions(ctx context.Context, roleName string) ([]RolePolicyInfo, error) {
	ret := _m.Called(ctx, roleName)

	var r0 []RolePolicyInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) []RolePolicyInfo); ok {
		r0 = rf(ctx, roleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RolePolicyInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamRolePermissionGetter_GetIAMRolePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIAMRolePermissions'
type MockIamRolePermissionGetter_GetIAMRolePermissions_Call struct {
	*mock.Call
}

// GetIAMRolePermissions is a helper method to define mock.On call
//  - ctx context.Context
//  - roleName string
func (_e *MockIamRolePermissionGetter_Expecter) GetIAMRolePermissions(ctx interface{}, roleName interface{}) *MockIamRolePermissionGetter_GetIAMRolePermissions_Call {
	return &MockIamRolePermissionGetter_GetIAMRolePermissions_Call{Call: _e.mock.On("GetIAMRolePermissions", ctx, roleName)}
}

func (_c *MockIamRolePermissionGetter_GetIAMRolePermissions_Call) Run(run func(ctx context.Context, roleName string)) *MockIamRolePermissionGetter_GetIAMRolePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockIamRolePermissionGetter_GetIAMRolePermissions_Call) Return(_a0 []RolePolicyInfo, _a1 error) *MockIamRolePermissionGetter_GetIAMRolePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockIamRolePermissionGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIamRolePermissionGetter creates a new instance of MockIamRolePermissionGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIamRolePermissionGetter(t mockConstructorTestingTNewMockIamRolePermissionGetter) *MockIamRolePermissionGetter {
	mock := &MockIamRolePermissionGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
