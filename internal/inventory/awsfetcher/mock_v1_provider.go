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

// Code generated by mockery v2.37.1. DO NOT EDIT.

package awsfetcher

import (
	context "context"

	awslib "github.com/elastic/cloudbeat/internal/resources/providers/awslib"

	mock "github.com/stretchr/testify/mock"
)

// mockV1Provider is an autogenerated mock type for the v1Provider type
type mockV1Provider struct {
	mock.Mock
}

type mockV1Provider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockV1Provider) EXPECT() *mockV1Provider_Expecter {
	return &mockV1Provider_Expecter{mock: &_m.Mock}
}

// DescribeAllLoadBalancers provides a mock function with given fields: _a0
func (_m *mockV1Provider) DescribeAllLoadBalancers(_a0 context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(_a0)

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockV1Provider_DescribeAllLoadBalancers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeAllLoadBalancers'
type mockV1Provider_DescribeAllLoadBalancers_Call struct {
	*mock.Call
}

// DescribeAllLoadBalancers is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *mockV1Provider_Expecter) DescribeAllLoadBalancers(_a0 interface{}) *mockV1Provider_DescribeAllLoadBalancers_Call {
	return &mockV1Provider_DescribeAllLoadBalancers_Call{Call: _e.mock.On("DescribeAllLoadBalancers", _a0)}
}

func (_c *mockV1Provider_DescribeAllLoadBalancers_Call) Run(run func(_a0 context.Context)) *mockV1Provider_DescribeAllLoadBalancers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockV1Provider_DescribeAllLoadBalancers_Call) Return(_a0 []awslib.AwsResource, _a1 error) *mockV1Provider_DescribeAllLoadBalancers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockV1Provider_DescribeAllLoadBalancers_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *mockV1Provider_DescribeAllLoadBalancers_Call {
	_c.Call.Return(run)
	return _c
}

// newMockV1Provider creates a new instance of mockV1Provider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockV1Provider(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockV1Provider {
	mock := &mockV1Provider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
