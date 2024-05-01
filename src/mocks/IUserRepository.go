// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	entity "go-ddd/src/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: user
func (_m *IUserRepository) Delete(user *entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: name
func (_m *IUserRepository) Find(name *entity.UserName) (*entity.User, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.UserName) (*entity.User, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(*entity.UserName) *entity.User); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.UserName) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: user
func (_m *IUserRepository) Save(user *entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
