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

func TestQuestionGetCLIHandlerValidate(t *testing.T) {
	var err error

	cmd := new(mocks.QuestionsQueryUsecase)

	err = clihandler.QuestionGetCLIHandler{cmd}.Validate()
	assert.NoError(t, err, "Should no error")

	err = clihandler.QuestionGetCLIHandler{cmd}.Validate("1")
	assert.NoError(t, err, "Should no error")

	cmd.AssertExpectations(t)
}

func TestQuestionGetCLIHandlerDescription(t *testing.T) {
	cmd := new(mocks.QuestionsQueryUsecase)

	desc := clihandler.QuestionGetCLIHandler{cmd}.Description()

	assert.Equal(t, "Shows question list", desc)

	cmd.AssertExpectations(t)
}

func TestQuestionGetCLIHandlerExample(t *testing.T) {
	cmd := new(mocks.QuestionsQueryUsecase)

	exp := clihandler.QuestionGetCLIHandler{cmd}.Example()

	assert.Equal(t, "questions", exp)

	cmd.AssertExpectations(t)
}

func TestQuestionGetCLIHandlerHandler(t *testing.T) {
	res := []entity.Questions{
		{
			ID:        "1",
			Questions: "Tes",
			Answers:   []string{"Tes"},
			Point:     1,
		},
	}

	cmd := new(mocks.QuestionsQueryUsecase)
	cmd.On("Get", mock.Anything).
		Return(res, nil, nil)

	clihandler.QuestionGetCLIHandler{cmd}.Handler()

	cmd.AssertExpectations(t)
}

func TestQuestionGetCLIHandlerHandlerRapperror(t *testing.T) {
	res := []entity.Questions{}

	cmd := new(mocks.QuestionsQueryUsecase)
	cmd.On("Get", mock.Anything).
		Return(res, nil, rapperror.ErrNotFound(
			"",
			"Question not found",
			"",
			nil,
		))

	clihandler.QuestionGetCLIHandler{cmd}.Handler()

	cmd.AssertExpectations(t)
}

func TestQuestionGetCLIHandlerHandlerError(t *testing.T) {
	res := []entity.Questions{}

	cmd := new(mocks.QuestionsQueryUsecase)
	cmd.On("Get", mock.Anything).
		Return(res, nil, errors.New("err: Something went wrong"))

	clihandler.QuestionGetCLIHandler{cmd}.Handler()

	cmd.AssertExpectations(t)
}
