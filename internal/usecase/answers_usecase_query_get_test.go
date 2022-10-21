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

func TestAnswerGetSuccess(t *testing.T) {
	repoRes := []entity.Answers{
		{
			ID: "1",
		},
	}

	repo := new(mocks.AnswersRepository)
	repo.On("Get", mock.Anything).
		Return(repoRes, nil)

	qstRepo := new(mocks.QuestionsRepository)

	prop := usecase.AnswersUsecaseProperty{
		Repository:         repo,
		QuestionRepository: qstRepo,
	}

	query := usecase.NewAnswersUsecaseQuery(prop)

	res, _, err := query.Get(context.Background())

	assert.NoError(t, err, "[TestAnswerGetSuccess] Should not error")
	assert.NotEmpty(t, res, "[TestAnswerGetSuccess] Should not empty")
	assert.Len(t, res, 1, "[TestAnswerGetSuccess] Should return \"1\" data")

	repo.AssertExpectations(t)
	qstRepo.AssertExpectations(t)
}

func TestAnswerGetError(t *testing.T) {
	repo := new(mocks.AnswersRepository)
	repo.On("Get", mock.Anything).
		Return(nil, rapperror.ErrNotFound(
			"",
			"Answer not found",
			"",
			nil,
		))

	qstRepo := new(mocks.QuestionsRepository)

	prop := usecase.AnswersUsecaseProperty{
		Repository:         repo,
		QuestionRepository: qstRepo,
	}

	query := usecase.NewAnswersUsecaseQuery(prop)

	res, _, err := query.Get(context.Background())

	assert.Error(t, err, "[TestAnswerGetError] Should error")
	assert.Empty(t, res, "[TestAnswerGetError] Should empty")
	assert.Len(t, res, 0, "[TestAnswerGetError] Should return \"0\" data")

	repo.AssertExpectations(t)
	qstRepo.AssertExpectations(t)
}
