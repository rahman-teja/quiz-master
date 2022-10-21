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

func TestQuestionUpdateCLIHandlerValidate(t *testing.T) {
	var err error

	cmd := new(mocks.QuestionsCommandUsecase)

	err = clihandler.QuestionUpdateCLIHandler{cmd}.Validate()
	assert.Error(t, err, "Should error")

	err = clihandler.QuestionUpdateCLIHandler{cmd}.Validate("1", "2", "3")
	assert.NoError(t, err, "Should no error")

	cmd.AssertExpectations(t)
}

func TestQuestionUpdateCLIHandlerDescription(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)

	desc := clihandler.QuestionUpdateCLIHandler{cmd}.Description()

	assert.Equal(t, "Updates a question", desc)

	cmd.AssertExpectations(t)
}

func TestQuestionUpdateCLIHandlerExample(t *testing.T) {
	cmd := new(mocks.QuestionsCommandUsecase)

	exp := clihandler.QuestionUpdateCLIHandler{cmd}.Example()

	assert.Equal(t, "update_question <no> <question> <answer>", exp)

	cmd.AssertExpectations(t)
}

func TestQuestionUpdateCLIHandlerHandler(t *testing.T) {
	res := entity.Questions{
		ID:        "1",
		Questions: "Tes",
		Answers:   []string{"Tes"},
		Point:     1,
	}

	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("model.Questions")).
		Return(res, nil, nil)

	clihandler.QuestionUpdateCLIHandler{cmd}.Handler("1", "2", "3")

	cmd.AssertExpectations(t)
}

func TestQuestionUpdateCLIHandlerHandlerRapperror(t *testing.T) {
	res := entity.Questions{}

	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("model.Questions")).
		Return(res, nil, rapperror.ErrNotFound(
			"",
			"Question not found",
			"",
			nil,
		))

	clihandler.QuestionUpdateCLIHandler{cmd}.Handler("1", "2", "3")

	cmd.AssertExpectations(t)
}

func TestQuestionUpdateCLIHandlerHandlerError(t *testing.T) {
	res := entity.Questions{}

	cmd := new(mocks.QuestionsCommandUsecase)
	cmd.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("model.Questions")).
		Return(res, nil, errors.New("err: Something went wrong"))

	clihandler.QuestionUpdateCLIHandler{cmd}.Handler("1", "2", "3")

	cmd.AssertExpectations(t)
}
