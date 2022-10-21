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

func TestQuestionRepositoryGetSuccess(t *testing.T) {
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
			ret := args[0].(*entity.Questions)
			ret.ID = "1"
		})

	csrMock.On("Close", mock.Anything).
		Return(nil)

	collMock := new(mocks.Collection)
	collMock.On("Find", mock.Anything).
		Return(csrMock, nil)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.NoError(t, err, "[TestQuestionRepositoryGetSuccess] Should not error")
	assert.Len(t, answers, 1, "[TestQuestionRepositoryGetSuccess] Question should 1 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}

func TestQuestionRepositoryGetErrInvalidCursor(t *testing.T) {
	csrMock := new(mocks.Cursor)

	collMock := new(mocks.Collection)
	collMock.On("Find", mock.Anything).
		Return(csrMock, memorydb.ErrNotFound)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.Error(t, err, "[TestQuestionRepositoryGetErrInvalidCursor] Should error")
	assert.Len(t, answers, 0, "[TestQuestionRepositoryGetErrInvalidCursor] Question should 0 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}

func TestQuestionRepositoryGetErrNotFound(t *testing.T) {
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

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.Error(t, err, "[TestQuestionRepositoryGetErrNotFound] Should error")
	assert.Len(t, answers, 0, "[TestQuestionRepositoryGetErrNotFound] Question should 0 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}

func TestQuestionRepositoryGetErrDecode(t *testing.T) {
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

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	answers, err := repo.Get(context.Background())

	assert.Error(t, err, "[TestQuestionRepositoryGetErrDecode] Should error")
	assert.Len(t, answers, 0, "[TestQuestionRepositoryGetErrDecode] Question should 0 data")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	csrMock.AssertExpectations(t)
}
