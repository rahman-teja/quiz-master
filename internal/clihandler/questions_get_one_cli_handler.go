package clihandler

import (
	"context"
	"errors"
	"fmt"

	"github.com/rahman-teja/quiz-master/internal/usecase"
	"gitlab.com/rteja-library3/rapperror"
)

type QuestionGetOneCLIHandler struct {
	Usecase usecase.QuestionsQueryUsecase
}

func (c QuestionGetOneCLIHandler) Validate(params ...string) error {
	if len(params) != 1 {
		return errors.New("err: Invalid command " + c.Example())
	}

	return nil
}

func (c QuestionGetOneCLIHandler) Handler(params ...string) {
	ctx := context.Background()

	result, _, err := c.Usecase.GetOne(ctx, params[0])
	if err != nil {
		rapp, ok := err.(*rapperror.AppError)
		if ok {
			fmt.Println(rapp.Message)
			return
		}

		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %s: \n Q: %s \n A: %s \n", result.ID, result.Questions, result.Answers[0])
}

func (c QuestionGetOneCLIHandler) Description() string {
	return "Shows a question"
}

func (c QuestionGetOneCLIHandler) Example() string {
	return "question <no>"
}
