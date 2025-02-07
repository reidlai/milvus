// Code generated by mockery v2.46.0. DO NOT EDIT.

package mock_client

import (
	context "context"

	messagespb "github.com/milvus-io/milvus/pkg/proto/messagespb"
	message "github.com/milvus-io/milvus/pkg/streaming/util/message"

	mock "github.com/stretchr/testify/mock"

	types "github.com/milvus-io/milvus/pkg/streaming/util/types"
)

// MockBroadcastService is an autogenerated mock type for the BroadcastService type
type MockBroadcastService struct {
	mock.Mock
}

type MockBroadcastService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBroadcastService) EXPECT() *MockBroadcastService_Expecter {
	return &MockBroadcastService_Expecter{mock: &_m.Mock}
}

// Ack provides a mock function with given fields: ctx, req
func (_m *MockBroadcastService) Ack(ctx context.Context, req types.BroadcastAckRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Ack")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.BroadcastAckRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBroadcastService_Ack_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ack'
type MockBroadcastService_Ack_Call struct {
	*mock.Call
}

// Ack is a helper method to define mock.On call
//   - ctx context.Context
//   - req types.BroadcastAckRequest
func (_e *MockBroadcastService_Expecter) Ack(ctx interface{}, req interface{}) *MockBroadcastService_Ack_Call {
	return &MockBroadcastService_Ack_Call{Call: _e.mock.On("Ack", ctx, req)}
}

func (_c *MockBroadcastService_Ack_Call) Run(run func(ctx context.Context, req types.BroadcastAckRequest)) *MockBroadcastService_Ack_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.BroadcastAckRequest))
	})
	return _c
}

func (_c *MockBroadcastService_Ack_Call) Return(_a0 error) *MockBroadcastService_Ack_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBroadcastService_Ack_Call) RunAndReturn(run func(context.Context, types.BroadcastAckRequest) error) *MockBroadcastService_Ack_Call {
	_c.Call.Return(run)
	return _c
}

// BlockUntilEvent provides a mock function with given fields: ctx, ev
func (_m *MockBroadcastService) BlockUntilEvent(ctx context.Context, ev *messagespb.BroadcastEvent) error {
	ret := _m.Called(ctx, ev)

	if len(ret) == 0 {
		panic("no return value specified for BlockUntilEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *messagespb.BroadcastEvent) error); ok {
		r0 = rf(ctx, ev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBroadcastService_BlockUntilEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockUntilEvent'
type MockBroadcastService_BlockUntilEvent_Call struct {
	*mock.Call
}

// BlockUntilEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - ev *messagespb.BroadcastEvent
func (_e *MockBroadcastService_Expecter) BlockUntilEvent(ctx interface{}, ev interface{}) *MockBroadcastService_BlockUntilEvent_Call {
	return &MockBroadcastService_BlockUntilEvent_Call{Call: _e.mock.On("BlockUntilEvent", ctx, ev)}
}

func (_c *MockBroadcastService_BlockUntilEvent_Call) Run(run func(ctx context.Context, ev *messagespb.BroadcastEvent)) *MockBroadcastService_BlockUntilEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*messagespb.BroadcastEvent))
	})
	return _c
}

func (_c *MockBroadcastService_BlockUntilEvent_Call) Return(_a0 error) *MockBroadcastService_BlockUntilEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBroadcastService_BlockUntilEvent_Call) RunAndReturn(run func(context.Context, *messagespb.BroadcastEvent) error) *MockBroadcastService_BlockUntilEvent_Call {
	_c.Call.Return(run)
	return _c
}

// Broadcast provides a mock function with given fields: ctx, msg
func (_m *MockBroadcastService) Broadcast(ctx context.Context, msg message.BroadcastMutableMessage) (*types.BroadcastAppendResult, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for Broadcast")
	}

	var r0 *types.BroadcastAppendResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, message.BroadcastMutableMessage) (*types.BroadcastAppendResult, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, message.BroadcastMutableMessage) *types.BroadcastAppendResult); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BroadcastAppendResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, message.BroadcastMutableMessage) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroadcastService_Broadcast_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Broadcast'
type MockBroadcastService_Broadcast_Call struct {
	*mock.Call
}

// Broadcast is a helper method to define mock.On call
//   - ctx context.Context
//   - msg message.BroadcastMutableMessage
func (_e *MockBroadcastService_Expecter) Broadcast(ctx interface{}, msg interface{}) *MockBroadcastService_Broadcast_Call {
	return &MockBroadcastService_Broadcast_Call{Call: _e.mock.On("Broadcast", ctx, msg)}
}

func (_c *MockBroadcastService_Broadcast_Call) Run(run func(ctx context.Context, msg message.BroadcastMutableMessage)) *MockBroadcastService_Broadcast_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(message.BroadcastMutableMessage))
	})
	return _c
}

func (_c *MockBroadcastService_Broadcast_Call) Return(_a0 *types.BroadcastAppendResult, _a1 error) *MockBroadcastService_Broadcast_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroadcastService_Broadcast_Call) RunAndReturn(run func(context.Context, message.BroadcastMutableMessage) (*types.BroadcastAppendResult, error)) *MockBroadcastService_Broadcast_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockBroadcastService) Close() {
	_m.Called()
}

// MockBroadcastService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockBroadcastService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockBroadcastService_Expecter) Close() *MockBroadcastService_Close_Call {
	return &MockBroadcastService_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockBroadcastService_Close_Call) Run(run func()) *MockBroadcastService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBroadcastService_Close_Call) Return() *MockBroadcastService_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBroadcastService_Close_Call) RunAndReturn(run func()) *MockBroadcastService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBroadcastService creates a new instance of MockBroadcastService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBroadcastService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBroadcastService {
	mock := &MockBroadcastService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
