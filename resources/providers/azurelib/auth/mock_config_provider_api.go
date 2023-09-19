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

// Code generated by mockery v2.24.0. DO NOT EDIT.

package auth

import mock "github.com/stretchr/testify/mock"

// MockConfigProviderAPI is an autogenerated mock type for the ConfigProviderAPI type
type MockConfigProviderAPI struct {
	mock.Mock
}

type MockConfigProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfigProviderAPI) EXPECT() *MockConfigProviderAPI_Expecter {
	return &MockConfigProviderAPI_Expecter{mock: &_m.Mock}
}

// GetAzureClientConfig provides a mock function with given fields:
func (_m *MockConfigProviderAPI) GetAzureClientConfig() (*AzureFactoryConfig, error) {
	ret := _m.Called()

	var r0 *AzureFactoryConfig
	var r1 error
	if rf, ok := ret.Get(0).(func() (*AzureFactoryConfig, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *AzureFactoryConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AzureFactoryConfig)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConfigProviderAPI_GetAzureClientConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAzureClientConfig'
type MockConfigProviderAPI_GetAzureClientConfig_Call struct {
	*mock.Call
}

// GetAzureClientConfig is a helper method to define mock.On call
func (_e *MockConfigProviderAPI_Expecter) GetAzureClientConfig() *MockConfigProviderAPI_GetAzureClientConfig_Call {
	return &MockConfigProviderAPI_GetAzureClientConfig_Call{Call: _e.mock.On("GetAzureClientConfig")}
}

func (_c *MockConfigProviderAPI_GetAzureClientConfig_Call) Run(run func()) *MockConfigProviderAPI_GetAzureClientConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfigProviderAPI_GetAzureClientConfig_Call) Return(_a0 *AzureFactoryConfig, _a1 error) *MockConfigProviderAPI_GetAzureClientConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConfigProviderAPI_GetAzureClientConfig_Call) RunAndReturn(run func() (*AzureFactoryConfig, error)) *MockConfigProviderAPI_GetAzureClientConfig_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockConfigProviderAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockConfigProviderAPI creates a new instance of MockConfigProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockConfigProviderAPI(t mockConstructorTestingTNewMockConfigProviderAPI) *MockConfigProviderAPI {
	mock := &MockConfigProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}