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

func TestQuestionGetOneCLIHandlerValidate(t *testing.T) {
	var err error

	cmd := new(mocks.QuestionsQueryUsecase)

	err = clihandler.QuestionGetOneCLIHandler{cmd}.Validate()
	assert.Error(t, err, "Should error")

	err = clihandler.QuestionGetOneCLIHandler{cmd}.Validate("1")
	assert.NoError(t, err, "Should no error")

	cmd.AssertExpectations(t)
}

func TestQuestionGetOneCLIHandlerDescription(t *testing.T) {
	cmd := new(mocks.QuestionsQueryUsecase)

	desc := clihandler.QuestionGetOneCLIHandler{cmd}.Description()

	assert.Equal(t, "Shows a question", desc)

	cmd.AssertExpectations(t)
}

func TestQuestionGetOneCLIHandlerExample(t *testing.T) {
	cmd := new(mocks.QuestionsQueryUsecase)

	exp := clihandler.QuestionGetOneCLIHandler{cmd}.Example()

	assert.Equal(t, "question <no>", exp)

	cmd.AssertExpectations(t)
}

func TestQuestionGetOneCLIHandlerHandler(t *testing.T) {
	res := entity.Questions{
		ID:        "1",
		Questions: "Tes",
		Answers:   []string{"Tes"},
		Point:     1,
	}

	cmd := new(mocks.QuestionsQueryUsecase)
	cmd.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
		Return(res, nil, nil)

	clihandler.QuestionGetOneCLIHandler{cmd}.Handler("1")

	cmd.AssertExpectations(t)
}

func TestQuestionGetOneCLIHandlerHandlerRapperror(t *testing.T) {
	res := entity.Questions{}

	cmd := new(mocks.QuestionsQueryUsecase)
	cmd.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
		Return(res, nil, rapperror.ErrNotFound(
			"",
			"Question not found",
			"",
			nil,
		))

	clihandler.QuestionGetOneCLIHandler{cmd}.Handler("1")

	cmd.AssertExpectations(t)
}

func TestQuestionGetOneCLIHandlerHandlerError(t *testing.T) {
	res := entity.Questions{}

	cmd := new(mocks.QuestionsQueryUsecase)
	cmd.On("GetOne", mock.Anything, mock.AnythingOfType("string")).
		Return(res, nil, errors.New("err: Something went wrong"))

	clihandler.QuestionGetOneCLIHandler{cmd}.Handler("1")

	cmd.AssertExpectations(t)
}
