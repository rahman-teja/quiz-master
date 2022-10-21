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

func TestQuestionRepositoryGetOneSuccess(t *testing.T) {
	srMock := new(mocks.SingleResult)

	srMock.On("Decode", mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			ret := args[0].(*entity.Questions)
			ret.ID = "1"
		})

	collMock := new(mocks.Collection)
	collMock.On("FindOne", mock.Anything, mock.AnythingOfType("string")).
		Return(srMock)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	res, err := repo.GetOne(context.Background(), "1")

	assert.NoError(t, err, "[TestQuestionRepositoryGetOneSuccess] Should not error")
	assert.NotEmpty(t, res, "[TestQuestionRepositoryGetOneSuccess] Result should not empty")
	assert.Equal(t, "1", res.ID, "[TestQuestionRepositoryGetOneSuccess] ID should \"1\"")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	srMock.AssertExpectations(t)
}

func TestQuestionRepositoryGetOneErrDecode(t *testing.T) {
	srMock := new(mocks.SingleResult)

	srMock.On("Decode", mock.Anything).
		Return(memorydb.ErrInternalServer)

	collMock := new(mocks.Collection)
	collMock.On("FindOne", mock.Anything, mock.AnythingOfType("string")).
		Return(srMock)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	res, err := repo.GetOne(context.Background(), "1")

	assert.Error(t, err, "[TestQuestionRepositoryGetOneErrDecode] Should error")
	assert.Empty(t, res, "[TestQuestionRepositoryGetOneErrDecode] Result should empty")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	srMock.AssertExpectations(t)
}

func TestQuestionRepositoryGetOneErrNotFound(t *testing.T) {
	srMock := new(mocks.SingleResult)

	srMock.On("Decode", mock.Anything).
		Return(memorydb.ErrNotFound)

	collMock := new(mocks.Collection)
	collMock.On("FindOne", mock.Anything, mock.AnythingOfType("string")).
		Return(srMock)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	res, err := repo.GetOne(context.Background(), "1")

	assert.Error(t, err, "[TestQuestionRepositoryGetOneErrDecode] Should error")
	assert.Empty(t, res, "[TestQuestionRepositoryGetOneErrDecode] Result should empty")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
	srMock.AssertExpectations(t)
}
