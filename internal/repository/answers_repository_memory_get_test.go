package repository_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/repository"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/rahman-teja/quiz-master/pkg/memorydb/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAnswerRepositoryGetSuccess(t *testing.T) {
	csrMock := new(mocks.Cursor)

	csrMock.On("Next", mock.Anything).
		Return(true).
		Once()

	csrMock.On("Next", mock.Anything).
		Return(false).
		Once()

	csrMock.On("Decode", mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			ret := args[0].(*entity.Answers)
			ret.ID = "65cc1f87-142d-4839-b65d-7a247688ef6f"
		})

	csrMock.On("Close", mock.Anything).
		Return(nil)

	collMock := new(mocks.Collection)
	collMock.On("Find", mock.Anything).
		Return(csrMock, nil)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewAnswersRepositoryMemory(repository.AnswersRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.NoError(t, err, "[TestAnswerRepositoryGetSuccess] Should not error")
	assert.Len(t, answers, 1, "[TestAnswerRepositoryGetSuccess] Answer should 1 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}

func TestAnswerRepositoryGetErrInvalidCursor(t *testing.T) {
	csrMock := new(mocks.Cursor)

	collMock := new(mocks.Collection)
	collMock.On("Find", mock.Anything).
		Return(csrMock, memorydb.ErrNotFound)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewAnswersRepositoryMemory(repository.AnswersRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.Error(t, err, "[TestAnswerRepositoryGetErrInvalidCursor] Should error")
	assert.Equal(t, "[err_not_found] Answer not found", err.Error(), "[TestAnswerRepositoryGetErrInvalidCursor] Should error not found")
	assert.Len(t, answers, 0, "[TestAnswerRepositoryGetErrInvalidCursor] Answer should 0 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}

func TestAnswerRepositoryGetErrNotFound(t *testing.T) {
	csrMock := new(mocks.Cursor)

	csrMock.On("Next", mock.Anything).
		Return(false).
		Once()

	csrMock.On("Close", mock.Anything).
		Return(nil)

	collMock := new(mocks.Collection)
	collMock.On("Find", mock.Anything).
		Return(csrMock, nil)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewAnswersRepositoryMemory(repository.AnswersRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.Error(t, err, "[TestAnswerRepositoryGetErrNotFound] Should error")
	assert.Equal(t, "[err_not_found] Answer not found", err.Error(), "[TestAnswerRepositoryGetErrNotFound] Should error not found")
	assert.Len(t, answers, 0, "[TestAnswerRepositoryGetErrNotFound] Answer should 0 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}

func TestAnswerRepositoryGetErrDecode(t *testing.T) {
	csrMock := new(mocks.Cursor)

	csrMock.On("Next", mock.Anything).
		Return(true).
		Once()

	csrMock.On("Decode", mock.Anything).
		Return(memorydb.ErrInvalidCursor)

	csrMock.On("Close", mock.Anything).
		Return(nil)

	collMock := new(mocks.Collection)
	collMock.On("Find", mock.Anything).
		Return(csrMock, nil)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewAnswersRepositoryMemory(repository.AnswersRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.Error(t, err, "[TestAnswerRepositoryGetErrDecode] Should error")
	assert.Equal(t, "[err_internal_server] something went wrong on answer", err.Error(), "[TestAnswerRepositoryGetErrDecode] Should error internal server")
	assert.Len(t, answers, 0, "[TestAnswerRepositoryGetErrDecode] Answer should 0 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}
