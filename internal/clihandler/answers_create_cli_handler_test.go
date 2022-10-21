package clihandler_test

import (
	"errors"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/clihandler"
	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/rteja-library3/rapperror"
)

func TestAnswersCreateCLIHandlerValidate(t *testing.T) {
	var err error

	cmd := new(mocks.AnswersCommandUsecase)

	err = clihandler.AnswersCreateCLIHandler{cmd}.Validate()
	assert.Error(t, err, "Should error")

	err = clihandler.AnswersCreateCLIHandler{cmd}.Validate("1", "2")
	assert.NoError(t, err, "Should no error")

	cmd.AssertExpectations(t)
}

func TestAnswersCreateCLIHandlerDescription(t *testing.T) {
	cmd := new(mocks.AnswersCommandUsecase)

	desc := clihandler.AnswersCreateCLIHandler{cmd}.Description()

	assert.Equal(t, "Answer a question", desc)

	cmd.AssertExpectations(t)
}

func TestAnswersCreateCLIHandlerExample(t *testing.T) {
	cmd := new(mocks.AnswersCommandUsecase)

	exp := clihandler.AnswersCreateCLIHandler{cmd}.Example()

	assert.Equal(t, "answer_question <no_question> <answer>", exp)

	cmd.AssertExpectations(t)
}

func TestAnswersCreateCLIHandlerHandler(t *testing.T) {
	res := entity.Answers{
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

	cmd := new(mocks.AnswersCommandUsecase)
	cmd.On("Create", mock.Anything, mock.AnythingOfType("model.Answers")).
		Return(res, nil, nil)

	clihandler.AnswersCreateCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}

func TestAnswersCreateCLIHandlerHandlerIncorrect(t *testing.T) {
	res := entity.Answers{
		Answer: "Tes",
		Questions: entity.Questions{
			ID:        "1",
			Questions: "Tes",
			Answers:   []string{"Tes1"},
			Point:     1,
		},
		IsCorrect: false,
		Point:     0,
	}

	cmd := new(mocks.AnswersCommandUsecase)
	cmd.On("Create", mock.Anything, mock.AnythingOfType("model.Answers")).
		Return(res, nil, nil)

	clihandler.AnswersCreateCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}

func TestAnswersCreateCLIHandlerHandlerRapperror(t *testing.T) {
	res := entity.Answers{}

	cmd := new(mocks.AnswersCommandUsecase)
	cmd.On("Create", mock.Anything, mock.AnythingOfType("model.Answers")).
		Return(res, nil, rapperror.ErrNotFound(
			"",
			"Answer not found",
			"",
			nil,
		))

	clihandler.AnswersCreateCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}

func TestAnswersCreateCLIHandlerHandlerError(t *testing.T) {
	res := entity.Answers{}

	cmd := new(mocks.AnswersCommandUsecase)
	cmd.On("Create", mock.Anything, mock.AnythingOfType("model.Answers")).
		Return(res, nil, errors.New("err: Something went wrong"))

	clihandler.AnswersCreateCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}
