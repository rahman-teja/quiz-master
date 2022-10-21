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

func TestQuestionCreateCLIHandlerValidate(t *testing.T) {
	var err error

	cmd := new(mocks.QuestionsCommandUsecase)

	err = clihandler.QuestionCreateCLIHandler{cmd}.Validate()
	assert.Error(t, err, "Should error")

	err = clihandler.QuestionCreateCLIHandler{cmd}.Validate("1", "2", "3")
	assert.NoError(t, err, "Should no error")

	cmd.AssertExpectations(t)
}

func TestQuestionCreateCLIHandlerDescription(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)

	desc := clihandler.QuestionCreateCLIHandler{cmd}.Description()

	assert.Equal(t, "Creates a question", desc)

	cmd.AssertExpectations(t)
}

func TestQuestionCreateCLIHandlerExample(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)

	exp := clihandler.QuestionCreateCLIHandler{cmd}.Example()

	assert.Equal(t, "create_question <no> <question> <answer>", exp)

	cmd.AssertExpectations(t)
}

func TestQuestionCreateCLIHandlerHandler(t *testing.T) {
	res := entity.Questions{
		ID:        "1",
		Questions: "Tes",
		Answers:   []string{"Tes"},
		Point:     1,
	}

	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Create", mock.Anything, mock.AnythingOfType("model.Questions")).
		Return(res, nil, nil)

	clihandler.QuestionCreateCLIHandler{cmd}.Handler("1", "2", "3")

	cmd.AssertExpectations(t)
}

func TestQuestionCreateCLIHandlerHandlerRapperror(t *testing.T) {
	res := entity.Questions{}

	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Create", mock.Anything, mock.AnythingOfType("model.Questions")).
		Return(res, nil, rapperror.ErrNotFound(
			"",
			"Question not found",
			"",
			nil,
		))

	clihandler.QuestionCreateCLIHandler{cmd}.Handler("1", "2", "3")

	cmd.AssertExpectations(t)
}

func TestQuestionCreateCLIHandlerHandlerError(t *testing.T) {
	res := entity.Questions{}

	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Create", mock.Anything, mock.AnythingOfType("model.Questions")).
		Return(res, nil, errors.New("err: Something went wrong"))

	clihandler.QuestionCreateCLIHandler{cmd}.Handler("1", "2", "3")

	cmd.AssertExpectations(t)
}
