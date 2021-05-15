// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/ivantedja/xmarvel/entity"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Search provides a mock function with given fields: ctx, filter
func (_m *Usecase) Search(ctx context.Context, filter map[string]string) (*entity.CharacterCollection, error) {
	ret := _m.Called(ctx, filter)

	var r0 *entity.CharacterCollection
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) *entity.CharacterCollection); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.CharacterCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Show provides a mock function with given fields: ctx, ID
func (_m *Usecase) Show(ctx context.Context, ID int) (*entity.Character, error) {
	ret := _m.Called(ctx, ID)

	var r0 *entity.Character
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Character); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Character)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
