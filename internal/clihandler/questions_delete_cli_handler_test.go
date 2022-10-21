package clihandler_test

import (
	"errors"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/clihandler"
	"github.com/rahman-teja/quiz-master/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/rteja-library3/rapperror"
)

func TestQuestionDeleteCLIHandlerValidate(t *testing.T) {
	var err error

	cmd := new(mocks.QuestionsCommandUsecase)

	err = clihandler.QuestionDeleteCLIHandler{cmd}.Validate()
	assert.Error(t, err, "Should error")

	err = clihandler.QuestionDeleteCLIHandler{cmd}.Validate("1")
	assert.NoError(t, err, "Should no error")

	cmd.AssertExpectations(t)
}

func TestQuestionDeleteCLIHandlerDescription(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)

	desc := clihandler.QuestionDeleteCLIHandler{cmd}.Description()

	assert.Equal(t, "Deletes a question", desc)

	cmd.AssertExpectations(t)
}

func TestQuestionDeleteCLIHandlerExample(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)

	exp := clihandler.QuestionDeleteCLIHandler{cmd}.Example()

	assert.Equal(t, "delete_question <no>", exp)

	cmd.AssertExpectations(t)
}

func TestQuestionDeleteCLIHandlerHandler(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Delete", mock.Anything, mock.AnythingOfType("string")).
		Return(nil)

	clihandler.QuestionDeleteCLIHandler{cmd}.Handler("1")

	cmd.AssertExpectations(t)
}

func TestQuestionDeleteCLIHandlerHandlerRapperror(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Delete", mock.Anything, mock.AnythingOfType("string")).
		Return(rapperror.ErrNotFound(
			"",
			"Question not found",
			"",
			nil,
		))

	clihandler.QuestionDeleteCLIHandler{cmd}.Handler("1")

	cmd.AssertExpectations(t)
}

func TestQuestionDeleteCLIHandlerHandlerError(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Delete", mock.Anything, mock.AnythingOfType("string")).
		Return(errors.New("err: Something went wrong"))

	clihandler.QuestionDeleteCLIHandler{cmd}.Handler("1")

	cmd.AssertExpectations(t)
}
