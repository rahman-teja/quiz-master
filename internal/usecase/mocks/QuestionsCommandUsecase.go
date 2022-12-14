// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "github.com/rahman-teja/quiz-master/internal/entity"
import mock "github.com/stretchr/testify/mock"
import model "github.com/rahman-teja/quiz-master/internal/model"

// QuestionsCommandUsecase is an autogenerated mock type for the QuestionsCommandUsecase type
type QuestionsCommandUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, payload
func (_m *QuestionsCommandUsecase) Create(ctx context.Context, payload model.Questions) (entity.Questions, interface{}, error) {
	ret := _m.Called(ctx, payload)

	var r0 entity.Questions
	if rf, ok := ret.Get(0).(func(context.Context, model.Questions) entity.Questions); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(entity.Questions)
	}

	var r1 interface{}
	if rf, ok := ret.Get(1).(func(context.Context, model.Questions) interface{}); ok {
		r1 = rf(ctx, payload)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.Questions) error); ok {
		r2 = rf(ctx, payload)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Delete provides a mock function with given fields: ctx, id
func (_m *QuestionsCommandUsecase) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, id, payload
func (_m *QuestionsCommandUsecase) Update(ctx context.Context, id string, payload model.Questions) (entity.Questions, interface{}, error) {
	ret := _m.Called(ctx, id, payload)

	var r0 entity.Questions
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Questions) entity.Questions); ok {
		r0 = rf(ctx, id, payload)
	} else {
		r0 = ret.Get(0).(entity.Questions)
	}

	var r1 interface{}
	if rf, ok := ret.Get(1).(func(context.Context, string, model.Questions) interface{}); ok {
		r1 = rf(ctx, id, payload)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, model.Questions) error); ok {
		r2 = rf(ctx, id, payload)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
