// Code generated by mockery v2.6.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// InsertOne provides a mock function with given fields: _a0, document
func (_m *Repository) InsertOne(_a0 context.Context, document interface{}) error {
	ret := _m.Called(_a0, document)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(_a0, document)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
