package clihandler

import (
	"context"
	"errors"
	"fmt"

	"github.com/rahman-teja/quiz-master/internal/usecase"
	"gitlab.com/rteja-library3/rapperror"
)

type QuestionDeleteCLIHandler struct {
	Usecase usecase.QuestionsCommandUsecase
}

func (c QuestionDeleteCLIHandler) Validate(params ...string) error {
	if len(params) != 1 {
		return errors.New("err: Invalid command " + c.Example())
	}

	return nil
}

func (c QuestionDeleteCLIHandler) Handler(params ...string) {
	ctx := context.Background()

	err := c.Usecase.Delete(ctx, params[0])
	if err != nil {
		rapp, ok := err.(*rapperror.AppError)
		if ok {
			fmt.Println(rapp.Message)
			return
		}

		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %s deleted!\n", params[0])
}

func (c QuestionDeleteCLIHandler) Description() string {
	return "Deletes a question"
}

func (c QuestionDeleteCLIHandler) Example() string {
	return "delete_question <no>"
}
