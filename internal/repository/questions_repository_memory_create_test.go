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

func TestQuestionRepositoryCreateSuccess(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Create", mock.Anything, mock.AnythingOfType("entity.Questions")).
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

	err := repo.Create(context.Background(), params)

	assert.NoError(t, err, "[TestQuestionRepositoryCreateSuccess] Should not error")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestQuestionRepositoryCreateErrDuplicate(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Create", mock.Anything, mock.AnythingOfType("entity.Questions")).
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

	err := repo.Create(context.Background(), params)

	assert.Error(t, err, "[TestQuestionRepositoryCreateErrDuplicate] Should error")
	assert.Equal(t, "[err_duplicate] Question no 1 already existed!", err.Error(), "[TestQuestionRepositoryCreateErrDuplicate] Should error duplicate")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestQuestionRepositoryCreateErrInternalServer(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Create", mock.Anything, mock.AnythingOfType("entity.Questions")).
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

	err := repo.Create(context.Background(), params)

	assert.Error(t, err, "[TestQuestionRepositoryCreateErrInternalServer] Should error")
	assert.Equal(t, "[err_internal_server] something went wrong on question", err.Error(), "[TestQuestionRepositoryCreateErrInternalServer] Should error duplicate")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}
