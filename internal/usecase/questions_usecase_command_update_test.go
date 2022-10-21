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

func TestQuestionsUpdateSuccess(t *testing.T) {
	repo := new(mocks.QuestionsRepository)
	repo.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("entity.Questions")).
		Return(nil)

	property := usecase.QuestionsUsecaseProperty{
		Repository: repo,
	}

	command := usecase.NewQuestionsUsecaseCommand(property)

	payload := model.Questions{
		ID:        "1",
		Questions: "1 + 1?",
		Answer:    "2",
	}

	res, _, err := command.Update(context.Background(), "id", payload)

	assert.NoError(t, err, "[TestQuestionsUpdateSuccess] Should not error")
	assert.NotEmpty(t, res, "[TestQuestionsUpdateSuccess] Should not empty")
	assert.Len(t, res.Answers, 3, "[TestQuestionsUpdateSuccess] Should have 3 answer")

	repo.AssertExpectations(t)
}

func TestQuestionsUpdateError(t *testing.T) {
	repo := new(mocks.QuestionsRepository)
	repo.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("entity.Questions")).
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

	res, _, err := command.Update(context.Background(), "id", payload)

	assert.Error(t, err, "[TestQuestionsUpdateError] Should error")
	assert.Empty(t, res, "[TestQuestionsUpdateError] Should empty")

	repo.AssertExpectations(t)
}
