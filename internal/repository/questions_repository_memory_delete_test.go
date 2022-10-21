package repository_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/repository"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/rahman-teja/quiz-master/pkg/memorydb/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQuestionRepositoryDeleteSuccess(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Delete", mock.Anything, mock.AnythingOfType("string")).
		Return(nil)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	err := repo.Delete(context.Background(), "1")

	assert.NoError(t, err, "[TestQuestionRepositoryDeleteSuccess] Should not error")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestQuestionRepositoryDeleteErrNotFound(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Delete", mock.Anything, mock.AnythingOfType("string")).
		Return(memorydb.ErrNotFound)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	err := repo.Delete(context.Background(), "1")

	assert.Error(t, err, "[TestQuestionRepositoryDeleteErrNotFound] Should error")
	assert.Equal(t, "[err_not_found] Question no 1 is not found", err.Error(), "[TestQuestionRepositoryDeleteErrNotFound] Should error not found")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestQuestionRepositoryDeleteErrInternalServer(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Delete", mock.Anything, mock.AnythingOfType("string")).
		Return(memorydb.ErrInternalServer)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	err := repo.Delete(context.Background(), "1")

	assert.Error(t, err, "[TestQuestionRepositoryDeleteErrInternalServer] Should error")
	assert.Equal(t, "[err_internal_server] something went wrong on question", err.Error(), "[TestQuestionRepositoryDeleteErrInternalServer] Should error not found")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}
