package usecase_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/repository/mocks"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/rteja-library3/rapperror"
)

func TestQuestionsGetSuccess(t *testing.T) {
	repo := new(mocks.QuestionsRepository)

	repoRes := []entity.Questions{
		{
			ID:        "1",
			Questions: "1 + 1?",
			Answers:   []string{"2", "Two", "two"},
			Point:     1,
		},
	}

	repo.On("Get", mock.Anything).
		Return(repoRes, nil)

	property := usecase.QuestionsUsecaseProperty{
		Repository: repo,
	}

	query := usecase.NewQuestionsUsecaseQuery(property)

	res, _, err := query.Get(context.Background())

	assert.NoError(t, err, "[TestQuestionsGetSuccess] Should not error")
	assert.NotEmpty(t, res, "[TestQuestionsGetSuccess] Should not empty")
	assert.Len(t, res, 1, "[TestQuestionsGetSuccess] Should have 1 data")

	repo.AssertExpectations(t)
}

func TestQuestionsGetError(t *testing.T) {
	repo := new(mocks.QuestionsRepository)

	repoRes := []entity.Questions{}

	repo.On("Get", mock.Anything).
		Return(repoRes, rapperror.ErrNotFound(
			"",
			"Question not found",
			"",
			nil,
		))

	property := usecase.QuestionsUsecaseProperty{
		Repository: repo,
	}

	query := usecase.NewQuestionsUsecaseQuery(property)

	res, _, err := query.Get(context.Background())

	assert.Error(t, err, "[TestQuestionsGetError] Should error")
	assert.Empty(t, res, "[TestQuestionsGetError] Should empty")

	repo.AssertExpectations(t)
}
