package clihandler

import (
	"context"
	"errors"
	"fmt"

	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"gitlab.com/rteja-library3/rapperror"
)

type AnswersCreateCLIHandler struct {
	Usecase usecase.AnswersCommandUsecase
}

func (c AnswersCreateCLIHandler) Validate(params ...string) error {
	if len(params) != 2 {
		return errors.New("err: Invalid command " + c.Example())
	}

	return nil
}

func (c AnswersCreateCLIHandler) Handler(params ...string) {
	ctx := context.Background()

	res, _, err := c.Usecase.Create(ctx, model.Answers{
		QuestionId: params[0],
		Answer:     params[1],
	})
	if err != nil {
		rapp, ok := err.(*rapperror.AppError)
		if ok {
			fmt.Println(rapp.Message)
			return
		}

		fmt.Println(err.Error())
		return
	}

	if !res.IsCorrect {
		fmt.Println("Incorrect!")
		return
	}

	fmt.Println("Correct!")
}

func (c AnswersCreateCLIHandler) Description() string {
	return "Answer a question"
}

func (c AnswersCreateCLIHandler) Example() string {
	return "answer_question <no_question> <answer>"
}
