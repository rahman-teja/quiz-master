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

func TestAnswerRepositoryCreateSuccess(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Create", mock.Anything, mock.AnythingOfType("entity.Answers")).
		Return(nil)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewAnswersRepositoryMemory(repository.AnswersRepositoryMemoryProp{
		DB: dbMock,
	})

	params := entity.Answers{
		Answer: "Tes",
		Questions: entity.Questions{
			ID:        "1",
			Questions: "Tes",
			Answers:   []string{"Tes"},
			Point:     1,
		},
		IsCorrect: true,
		Point:     1,
	}

	err := repo.Create(context.Background(), params)

	assert.NoError(t, err, "[TestAnswerRepositoryCreateSuccess] Should not error")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}

func TestAnswerRepositoryCreateErrDuplicate(t *testing.T) {
	collMock := new(mocks.Collection)
	collMock.On("Create", mock.Anything, mock.AnythingOfType("entity.Answers")).
		Return(memorydb.ErrDuplicate)

	dbMock := new(mocks.Database)
	dbMock.On("Collection", mock.AnythingOfType("string")).
		Return(collMock)

	repo := repository.NewAnswersRepositoryMemory(repository.AnswersRepositoryMemoryProp{
		DB: dbMock,
	})

	params := entity.Answers{
		Answer: "Tes",
		Questions: entity.Questions{
			ID:        "1",
			Questions: "Tes",
			Answers:   []string{"Tes"},
			Point:     1,
		},
		IsCorrect: true,
		Point:     1,
	}

	err := repo.Create(context.Background(), params)

	assert.Error(t, err, "[TestAnswerRepositoryCreateErrDuplicate] Should error")
	assert.Equal(t, "[err_duplicate] Answer already created", err.Error(), "[TestAnswerRepositoryCreateErrDuplicate] Should error duplicate")

	dbMock.AssertExpectations(t)
	collMock.AssertExpectations(t)
}
