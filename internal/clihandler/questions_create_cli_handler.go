package clihandler

import (
	"context"
	"errors"
	"fmt"

	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"gitlab.com/rteja-library3/rapperror"
)

type QuestionCreateCLIHandler struct {
	Usecase usecase.QuestionsCommandUsecase
}

func (c QuestionCreateCLIHandler) Validate(params ...string) error {
	if len(params) != 3 {
		return errors.New("err: Invalid command " + c.Example())
	}

	return nil
}

func (c QuestionCreateCLIHandler) Handler(params ...string) {
	ctx := context.Background()

	res, _, err := c.Usecase.Create(ctx, model.Questions{
		ID:        params[0],
		Questions: params[1],
		Answer:    params[2],
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

	fmt.Printf("Question no %s created: \nQ: %s \nA: %s \n", res.ID, res.Questions, res.Answers[0])
}

func (c QuestionCreateCLIHandler) Description() string {
	return "Creates a question"
}

func (c QuestionCreateCLIHandler) Example() string {
	return "create_question <no> <question> <answer>"
}
