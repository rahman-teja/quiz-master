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

func TestQuestionsGetOneSuccess(t *testing.T) {
	repo := new(mocks.QuestionsRepository)

	repoRes := entity.Questions{
		ID:        "1",
		Questions: "1 + 1?",
		Answers:   []string{"2", "Two", "two"},
		Point:     1,
	}

	repo.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
		Return(repoRes, nil)

	property := usecase.QuestionsUsecaseProperty{
		Repository: repo,
	}

	query := usecase.NewQuestionsUsecaseQuery(property)

	res, _, err := query.GetOne(context.Background(), "id")

	assert.NoError(t, err, "[TestQuestionsGetOneSuccess] Should not error")
	assert.NotEmpty(t, res, "[TestQuestionsGetOneSuccess] Should not empty")
	assert.Equal(t, "1", res.ID, "[TestQuestionsGetOneSuccess] ID should \"1\"")

	repo.AssertExpectations(t)
}

func TestQuestionsGetOneError(t *testing.T) {
	repo := new(mocks.QuestionsRepository)

	repoRes := entity.Questions{}

	repo.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
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

	res, _, err := query.GetOne(context.Background(), "id")

	assert.Error(t, err, "[TestQuestionsGetOneError] Should error")
	assert.Empty(t, res, "[TestQuestionsGetOneError] Should empty")

	repo.AssertExpectations(t)
}
