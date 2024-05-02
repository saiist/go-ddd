// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	users "go-ddd/src/domain/models/users"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: user
func (_m *IUserRepository) Delete(user *users.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*users.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *IUserRepository) FindAll() ([]*users.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*users.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*users.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*users.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *IUserRepository) FindById(id *users.UserId) (*users.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*users.UserId) (*users.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(*users.UserId) *users.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*users.UserId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *IUserRepository) FindByName(name *users.UserName) (*users.User, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for FindByName")
	}

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*users.UserName) (*users.User, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(*users.UserName) *users.User); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*users.UserName) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: user
func (_m *IUserRepository) Save(user *users.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*users.User) error); ok {
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
