// Code generated by mockery v2.37.1. DO NOT EDIT.

package mockpbdns

import (
	context "context"

	pbdns "github.com/hashicorp/consul/proto-public/pbdns"
	mock "github.com/stretchr/testify/mock"
)

// DNSServiceServer is an autogenerated mock type for the DNSServiceServer type
type DNSServiceServer struct {
	mock.Mock
}

type DNSServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *DNSServiceServer) EXPECT() *DNSServiceServer_Expecter {
	return &DNSServiceServer_Expecter{mock: &_m.Mock}
}

// Query provides a mock function with given fields: _a0, _a1
func (_m *DNSServiceServer) Query(_a0 context.Context, _a1 *pbdns.QueryRequest) (*pbdns.QueryResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pbdns.QueryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pbdns.QueryRequest) (*pbdns.QueryResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pbdns.QueryRequest) *pbdns.QueryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pbdns.QueryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pbdns.QueryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DNSServiceServer_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type DNSServiceServer_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *pbdns.QueryRequest
func (_e *DNSServiceServer_Expecter) Query(_a0 interface{}, _a1 interface{}) *DNSServiceServer_Query_Call {
	return &DNSServiceServer_Query_Call{Call: _e.mock.On("Query", _a0, _a1)}
}

func (_c *DNSServiceServer_Query_Call) Run(run func(_a0 context.Context, _a1 *pbdns.QueryRequest)) *DNSServiceServer_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pbdns.QueryRequest))
	})
	return _c
}

func (_c *DNSServiceServer_Query_Call) Return(_a0 *pbdns.QueryResponse, _a1 error) *DNSServiceServer_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DNSServiceServer_Query_Call) RunAndReturn(run func(context.Context, *pbdns.QueryRequest) (*pbdns.QueryResponse, error)) *DNSServiceServer_Query_Call {
	_c.Call.Return(run)
	return _c
}

// NewDNSServiceServer creates a new instance of DNSServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDNSServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *DNSServiceServer {
	mock := &DNSServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
