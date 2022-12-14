// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import memorydb "github.com/rahman-teja/quiz-master/pkg/memorydb"
import mock "github.com/stretchr/testify/mock"

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Collection provides a mock function with given fields: colname
func (_m *Database) Collection(colname string) memorydb.Collection {
	ret := _m.Called(colname)

	var r0 memorydb.Collection
	if rf, ok := ret.Get(0).(func(string) memorydb.Collection); ok {
		r0 = rf(colname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memorydb.Collection)
		}
	}

	return r0
}
