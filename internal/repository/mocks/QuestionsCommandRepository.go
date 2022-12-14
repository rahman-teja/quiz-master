// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "github.com/rahman-teja/quiz-master/internal/entity"
import mock "github.com/stretchr/testify/mock"

// QuestionsCommandRepository is an autogenerated mock type for the QuestionsCommandRepository type
type QuestionsCommandRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, questions
func (_m *QuestionsCommandRepository) Create(ctx context.Context, questions entity.Questions) error {
	ret := _m.Called(ctx, questions)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Questions) error); ok {
		r0 = rf(ctx, questions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *QuestionsCommandRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, id, questions
func (_m *QuestionsCommandRepository) Update(ctx context.Context, id string, questions entity.Questions) error {
	ret := _m.Called(ctx, id, questions)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.Questions) error); ok {
		r0 = rf(ctx, id, questions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
