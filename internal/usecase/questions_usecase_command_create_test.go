package usecase_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/repository/mocks"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/rteja-library3/rapperror"
)

func TestQuestionsCreateSuccess(t *testing.T) {
	repo := new(mocks.QuestionsRepository)
	repo.On("Create", mock.Anything, mock.AnythingOfType("entity.Questions")).
		Return(nil)

	property := usecase.QuestionsUsecaseProperty{
		Repository: repo,
	}

	command := usecase.NewQuestionsUsecaseCommand(property)

	payload := model.Questions{
		ID:        "1",
		Questions: "West, North, East?",
		Answer:    "South",
	}

	res, _, err := command.Create(context.Background(), payload)

	assert.NoError(t, err, "[TestQuestionsCreateSuccess] Should not error")
	assert.NotEmpty(t, res, "[TestQuestionsCreateSuccess] Should not empty")
	assert.Len(t, res.Answers, 1, "[TestQuestionsCreateSuccess] Should have 1 answer")

	repo.AssertExpectations(t)
}

func TestQuestionsCreateError(t *testing.T) {
	repo := new(mocks.QuestionsRepository)
	repo.On("Create", mock.Anything, mock.AnythingOfType("entity.Questions")).
		Return(rapperror.ErrNotFound(
			"",
			"Question not found",
			"",
			nil,
		))

	property := usecase.QuestionsUsecaseProperty{
		Repository: repo,
	}

	command := usecase.NewQuestionsUsecaseCommand(property)

	payload := model.Questions{
		ID:        "1",
		Questions: "1 + 1?",
		Answer:    "two",
	}

	res, _, err := command.Create(context.Background(), payload)

	assert.Error(t, err, "[TestQuestionsCreateError] Should error")
	assert.Empty(t, res, "[TestQuestionsCreateError] Should empty")

	repo.AssertExpectations(t)
}
