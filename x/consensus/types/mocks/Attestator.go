// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	types "github.com/palomachain/paloma/x/consensus/types"
)

// Attestator is an autogenerated mock type for the Attestator type
type Attestator struct {
	mock.Mock
}

// BytesToSign provides a mock function with given fields:
func (_m *Attestator) BytesToSign() types.BytesToSignFunc {
	ret := _m.Called()

	var r0 types.BytesToSignFunc
	if rf, ok := ret.Get(0).(func() types.BytesToSignFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.BytesToSignFunc)
		}
	}

	return r0
}

// ChainInfo provides a mock function with given fields:
func (_m *Attestator) ChainInfo() (string, string) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// ConsensusQueue provides a mock function with given fields:
func (_m *Attestator) ConsensusQueue() types.ConsensusQueueType {
	ret := _m.Called()

	var r0 types.ConsensusQueueType
	if rf, ok := ret.Get(0).(func() types.ConsensusQueueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.ConsensusQueueType)
	}

	return r0
}

// ProcessAllEvidence provides a mock function with given fields: ctx, task, evidence
func (_m *Attestator) ProcessAllEvidence(ctx context.Context, task types.AttestTask, evidence []types.Evidence) (types.AttestResult, error) {
	ret := _m.Called(ctx, task, evidence)

	var r0 types.AttestResult
	if rf, ok := ret.Get(0).(func(context.Context, types.AttestTask, []types.Evidence) types.AttestResult); ok {
		r0 = rf(ctx, task, evidence)
	} else {
		r0 = ret.Get(0).(types.AttestResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.AttestTask, []types.Evidence) error); ok {
		r1 = rf(ctx, task, evidence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Type provides a mock function with given fields:
func (_m *Attestator) Type() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// ValidateEvidence provides a mock function with given fields: ctx, task, evidence
func (_m *Attestator) ValidateEvidence(ctx context.Context, task types.AttestTask, evidence types.Evidence) error {
	ret := _m.Called(ctx, task, evidence)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AttestTask, types.Evidence) error); ok {
		r0 = rf(ctx, task, evidence)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifySignature provides a mock function with given fields:
func (_m *Attestator) VerifySignature() types.VerifySignatureFunc {
	ret := _m.Called()

	var r0 types.VerifySignatureFunc
	if rf, ok := ret.Get(0).(func() types.VerifySignatureFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.VerifySignatureFunc)
		}
	}

	return r0
}

// NewAttestator creates a new instance of Attestator. It also registers a cleanup function to assert the mocks expectations.
func NewAttestator(t testing.TB) *Attestator {
	mock := &Attestator{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
