package usecase_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/repository/mocks"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/rteja-library3/rapperror"
)

func TestQuestionsDeleteSuccess(t *testing.T) {
	repo := new(mocks.QuestionsRepository)
	repo.On("Delete", mock.Anything, mock.AnythingOfType("string")).
		Return(nil)

	property := usecase.QuestionsUsecaseProperty{
		Repository: repo,
	}

	command := usecase.NewQuestionsUsecaseCommand(property)

	err := command.Delete(context.Background(), "id")

	assert.NoError(t, err, "[TestQuestionsDeleteSuccess] Should not error")

	repo.AssertExpectations(t)
}

func TestQuestionsDeleteErr(t *testing.T) {
	repo := new(mocks.QuestionsRepository)
	repo.On("Delete", mock.Anything, mock.AnythingOfType("string")).
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

	err := command.Delete(context.Background(), "id")

	assert.Error(t, err, "[TestQuestionsDeleteErr] Should error")

	repo.AssertExpectations(t)
}
