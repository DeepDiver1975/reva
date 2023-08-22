// Copyright 2018-2022 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	collaborationv1beta1 "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"

	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	mock "github.com/stretchr/testify/mock"

	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"

	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// GetReceivedShare provides a mock function with given fields: ctx, ref
func (_m *Manager) GetReceivedShare(ctx context.Context, ref *collaborationv1beta1.ShareReference) (*collaborationv1beta1.ReceivedShare, error) {
	ret := _m.Called(ctx, ref)

	var r0 *collaborationv1beta1.ReceivedShare
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ShareReference) (*collaborationv1beta1.ReceivedShare, error)); ok {
		return rf(ctx, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ShareReference) *collaborationv1beta1.ReceivedShare); ok {
		r0 = rf(ctx, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.ReceivedShare)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.ShareReference) error); ok {
		r1 = rf(ctx, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShare provides a mock function with given fields: ctx, ref
func (_m *Manager) GetShare(ctx context.Context, ref *collaborationv1beta1.ShareReference) (*collaborationv1beta1.Share, error) {
	ret := _m.Called(ctx, ref)

	var r0 *collaborationv1beta1.Share
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ShareReference) (*collaborationv1beta1.Share, error)); ok {
		return rf(ctx, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ShareReference) *collaborationv1beta1.Share); ok {
		r0 = rf(ctx, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.Share)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.ShareReference) error); ok {
		r1 = rf(ctx, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReceivedShares provides a mock function with given fields: ctx, filters
func (_m *Manager) ListReceivedShares(ctx context.Context, filters []*collaborationv1beta1.Filter) ([]*collaborationv1beta1.ReceivedShare, error) {
	ret := _m.Called(ctx, filters)

	var r0 []*collaborationv1beta1.ReceivedShare
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*collaborationv1beta1.Filter) ([]*collaborationv1beta1.ReceivedShare, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*collaborationv1beta1.Filter) []*collaborationv1beta1.ReceivedShare); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*collaborationv1beta1.ReceivedShare)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*collaborationv1beta1.Filter) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListShares provides a mock function with given fields: ctx, filters
func (_m *Manager) ListShares(ctx context.Context, filters []*collaborationv1beta1.Filter) ([]*collaborationv1beta1.Share, error) {
	ret := _m.Called(ctx, filters)

	var r0 []*collaborationv1beta1.Share
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*collaborationv1beta1.Filter) ([]*collaborationv1beta1.Share, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*collaborationv1beta1.Filter) []*collaborationv1beta1.Share); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*collaborationv1beta1.Share)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*collaborationv1beta1.Filter) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Share provides a mock function with given fields: ctx, md, g
func (_m *Manager) Share(ctx context.Context, md *providerv1beta1.ResourceInfo, g *collaborationv1beta1.ShareGrant) (*collaborationv1beta1.Share, error) {
	ret := _m.Called(ctx, md, g)

	var r0 *collaborationv1beta1.Share
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *providerv1beta1.ResourceInfo, *collaborationv1beta1.ShareGrant) (*collaborationv1beta1.Share, error)); ok {
		return rf(ctx, md, g)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *providerv1beta1.ResourceInfo, *collaborationv1beta1.ShareGrant) *collaborationv1beta1.Share); ok {
		r0 = rf(ctx, md, g)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.Share)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *providerv1beta1.ResourceInfo, *collaborationv1beta1.ShareGrant) error); ok {
		r1 = rf(ctx, md, g)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unshare provides a mock function with given fields: ctx, ref
func (_m *Manager) Unshare(ctx context.Context, ref *collaborationv1beta1.ShareReference) error {
	ret := _m.Called(ctx, ref)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ShareReference) error); ok {
		r0 = rf(ctx, ref)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateReceivedShare provides a mock function with given fields: ctx, _a1, fieldMask, forUser
func (_m *Manager) UpdateReceivedShare(ctx context.Context, _a1 *collaborationv1beta1.ReceivedShare, fieldMask *fieldmaskpb.FieldMask, forUser *userv1beta1.UserId) (*collaborationv1beta1.ReceivedShare, error) {
	ret := _m.Called(ctx, _a1, fieldMask, forUser)

	var r0 *collaborationv1beta1.ReceivedShare
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ReceivedShare, *fieldmaskpb.FieldMask, *userv1beta1.UserId) (*collaborationv1beta1.ReceivedShare, error)); ok {
		return rf(ctx, _a1, fieldMask, forUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ReceivedShare, *fieldmaskpb.FieldMask, *userv1beta1.UserId) *collaborationv1beta1.ReceivedShare); ok {
		r0 = rf(ctx, _a1, fieldMask, forUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.ReceivedShare)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.ReceivedShare, *fieldmaskpb.FieldMask, *userv1beta1.UserId) error); ok {
		r1 = rf(ctx, _a1, fieldMask, forUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateShare provides a mock function with given fields: ctx, ref, p, updated, fieldMask
func (_m *Manager) UpdateShare(ctx context.Context, ref *collaborationv1beta1.ShareReference, p *collaborationv1beta1.SharePermissions, updated *collaborationv1beta1.Share, fieldMask *fieldmaskpb.FieldMask) (*collaborationv1beta1.Share, error) {
	ret := _m.Called(ctx, ref, p, updated, fieldMask)

	var r0 *collaborationv1beta1.Share
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ShareReference, *collaborationv1beta1.SharePermissions, *collaborationv1beta1.Share, *fieldmaskpb.FieldMask) (*collaborationv1beta1.Share, error)); ok {
		return rf(ctx, ref, p, updated, fieldMask)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ShareReference, *collaborationv1beta1.SharePermissions, *collaborationv1beta1.Share, *fieldmaskpb.FieldMask) *collaborationv1beta1.Share); ok {
		r0 = rf(ctx, ref, p, updated, fieldMask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.Share)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.ShareReference, *collaborationv1beta1.SharePermissions, *collaborationv1beta1.Share, *fieldmaskpb.FieldMask) error); ok {
		r1 = rf(ctx, ref, p, updated, fieldMask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t mockConstructorTestingTNewManager) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
