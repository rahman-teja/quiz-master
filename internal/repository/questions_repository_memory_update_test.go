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

func TestQuestionRepositoryUpdateSuccess(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("entity.Questions")).
		Return(nil)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	params := entity.Questions{
		ID:        "1",
		Questions: "Tes",
		Answers:   []string{"Tes"},
		Point:     1,
	}

	err := repo.Update(context.Background(), "1", params)

	assert.NoError(t, err, "[TestQuestionRepositoryUpdateSuccess] Should not error")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestQuestionRepositoryUpdateErrDuplicate(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("entity.Questions")).
		Return(memorydb.ErrDuplicate)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	params := entity.Questions{
		ID:        "1",
		Questions: "Tes",
		Answers:   []string{"Tes"},
		Point:     1,
	}

	err := repo.Update(context.Background(), "1", params)

	assert.Error(t, err, "[TestQuestionRepositoryUpdateErrDuplicate] Should error")
	assert.Equal(t, "[err_duplicate] Question no 1 already existed!", err.Error(), "[TestQuestionRepositoryUpdateErrDuplicate] Should error duplicate")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestQuestionRepositoryUpdateErrNotFound(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("entity.Questions")).
		Return(memorydb.ErrNotFound)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	params := entity.Questions{
		ID:        "1",
		Questions: "Tes",
		Answers:   []string{"Tes"},
		Point:     1,
	}

	err := repo.Update(context.Background(), "1", params)

	assert.Error(t, err, "[TestQuestionRepositoryUpdateErrNotFound] Should error")
	assert.Equal(t, "[err_not_found] Question no 1 is not found", err.Error(), "[TestQuestionRepositoryUpdateErrNotFound] Should error not found")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestQuestionRepositoryUpdateErrInternalServer(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("entity.Questions")).
		Return(memorydb.ErrInternalServer)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
		DB: dbMock,
	})

	params := entity.Questions{
		ID:        "1",
		Questions: "Tes",
		Answers:   []string{"Tes"},
		Point:     1,
	}

	err := repo.Update(context.Background(), "1", params)

	assert.Error(t, err, "[TestQuestionRepositoryUpdateErrInternalServer] Should error")
	assert.Equal(t, "[err_internal_server] something went wrong on question", err.Error(), "[TestQuestionRepositoryUpdateErrInternalServer] Should error not found")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}
