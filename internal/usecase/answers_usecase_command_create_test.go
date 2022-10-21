package usecase_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/repository/mocks"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/rteja-library3/rapperror"
)

func TestAnswerCreateSuccess(t *testing.T) {
	repo := new(mocks.AnswersRepository)
	qstRepo := new(mocks.QuestionsRepository)

	qstRepo.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
		Return(entity.Questions{
			ID:        "1",
			Questions: "1 + 1?",
			Answers:   []string{"2", "Two", "two"},
			Point:     1,
		}, nil)

	repo.On("Create", mock.Anything, mock.AnythingOfType("entity.Answers")).
		Return(nil)

	prop := usecase.AnswersUsecaseProperty{
		Repository:         repo,
		QuestionRepository: qstRepo,
	}

	command := usecase.NewAnswersUsecaseCommand(prop)

	payload := model.Answers{
		QuestionId: "1",
		Answer:     "2",
	}

	res, _, err := command.Create(context.Background(), payload)

	assert.NoError(t, err, "[TestAnswerCreateSuccess] Should not error")
	assert.NotEmpty(t, res, "[TestAnswerCreateSuccess] Should not empty")
	assert.Equal(t, true, res.IsCorrect, "[TestAnswerCreateSuccess] Should not empty")

	repo.AssertExpectations(t)
	qstRepo.AssertExpectations(t)
}

func TestAnswerCreateErrGetOne(t *testing.T) {
	repo := new(mocks.AnswersRepository)
	qstRepo := new(mocks.QuestionsRepository)

	qstRepo.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
		Return(entity.Questions{}, rapperror.ErrNotFound(
			"",
			"Answer not found",
			"",
			nil,
		))

	prop := usecase.AnswersUsecaseProperty{
		Repository:         repo,
		QuestionRepository: qstRepo,
	}

	command := usecase.NewAnswersUsecaseCommand(prop)

	payload := model.Answers{
		QuestionId: "1",
		Answer:     "2",
	}

	res, _, err := command.Create(context.Background(), payload)

	assert.Error(t, err, "[TestAnswerCreateErrGetOne] Should error")
	assert.Empty(t, res, "[TestAnswerCreateErrGetOne] Should empty")

	repo.AssertExpectations(t)
	qstRepo.AssertExpectations(t)
}

func TestAnswerCreateErrCreate(t *testing.T) {
	repo := new(mocks.AnswersRepository)
	qstRepo := new(mocks.QuestionsRepository)

	qstRepo.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
		Return(entity.Questions{
			ID:        "1",
			Questions: "1 + 1?",
			Answers:   []string{"2", "Two", "two"},
			Point:     1,
		}, nil)

	repo.On("Create", mock.Anything, mock.AnythingOfType("entity.Answers")).
		Return(rapperror.ErrInternalServerError(
			"",
			"something went wrong on answer",
			"",
			nil,
		))

	prop := usecase.AnswersUsecaseProperty{
		Repository:         repo,
		QuestionRepository: qstRepo,
	}

	command := usecase.NewAnswersUsecaseCommand(prop)

	payload := model.Answers{
		QuestionId: "1",
		Answer:     "2",
	}

	res, _, err := command.Create(context.Background(), payload)

	assert.Error(t, err, "[TestAnswerCreateErrCreate] Should error")
	assert.Empty(t, res, "[TestAnswerCreateErrCreate] Should empty")

	repo.AssertExpectations(t)
	qstRepo.AssertExpectations(t)
}
