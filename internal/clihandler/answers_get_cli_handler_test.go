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

func TestAnswersGetCLIHandlerValidate(t *testing.T) {
	var err error

	cmd := new(mocks.AnswersQueryUsecase)

	err = clihandler.AnswersGetCLIHandler{cmd}.Validate()
	assert.NoError(t, err, "Should no error")

	err = clihandler.AnswersGetCLIHandler{cmd}.Validate("1", "2")
	assert.NoError(t, err, "Should no error")

	cmd.AssertExpectations(t)
}

func TestAnswersGetCLIHandlerDescription(t *testing.T) {
	cmd := new(mocks.AnswersQueryUsecase)

	desc := clihandler.AnswersGetCLIHandler{cmd}.Description()

	assert.Equal(t, "Shows answer list", desc)

	cmd.AssertExpectations(t)
}

func TestAnswersGetCLIHandlerExample(t *testing.T) {
	cmd := new(mocks.AnswersQueryUsecase)

	exp := clihandler.AnswersGetCLIHandler{cmd}.Example()

	assert.Equal(t, "answers", exp)

	cmd.AssertExpectations(t)
}

func TestAnswersGetCLIHandlerHandler(t *testing.T) {
	res := []entity.Answers{
		{
			Answer: "Tes",
			Questions: entity.Questions{
				ID:        "1",
				Questions: "Tes",
				Answers:   []string{"Tes"},
				Point:     1,
			},
			IsCorrect: true,
			Point:     1,
		},
	}

	cmd := new(mocks.AnswersQueryUsecase)
	cmd.On("Get", mock.Anything).
		Return(res, nil, nil)

	clihandler.AnswersGetCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}

func TestAnswersGetCLIHandlerHandlerIncorrect(t *testing.T) {
	res := []entity.Answers{
		{
			Answer: "Tes",
			Questions: entity.Questions{
				ID:        "1",
				Questions: "Tes",
				Answers:   []string{"Tes"},
				Point:     1,
			},
			IsCorrect: false,
			Point:     0,
		},
	}

	cmd := new(mocks.AnswersQueryUsecase)
	cmd.On("Get", mock.Anything).
		Return(res, nil, nil)

	clihandler.AnswersGetCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}

func TestAnswersGetCLIHandlerHandlerRapperror(t *testing.T) {
	res := []entity.Answers{}

	cmd := new(mocks.AnswersQueryUsecase)
	cmd.On("Get", mock.Anything).
		Return(res, nil, rapperror.ErrNotFound(
			"",
			"Answer not found",
			"",
			nil,
		))

	clihandler.AnswersGetCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}

func TestAnswersGetCLIHandlerHandlerError(t *testing.T) {
	res := []entity.Answers{}

	cmd := new(mocks.AnswersQueryUsecase)
	cmd.On("Get", mock.Anything).
		Return(res, nil, errors.New("err: Something went wrong"))

	clihandler.AnswersGetCLIHandler{cmd}.Handler("1", "2")

	cmd.AssertExpectations(t)
}
