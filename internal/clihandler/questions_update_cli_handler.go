package clihandler

import (
	"context"
	"errors"
	"fmt"

	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"gitlab.com/rteja-library3/rapperror"
)

type QuestionUpdateCLIHandler struct {
	Usecase usecase.QuestionsCommandUsecase
}

func (c QuestionUpdateCLIHandler) Validate(params ...string) error {
	if len(params) != 3 {
		return errors.New("err: Invalid command " + c.Example())
	}

	return nil
}

func (c QuestionUpdateCLIHandler) Handler(params ...string) {
	ctx := context.Background()

	res, _, err := c.Usecase.Update(ctx, params[0], model.Questions{
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

	fmt.Printf("Question no %s updated: \nQ: %s \nA: %s \n", res.ID, res.Questions, res.Answers[0])
}

func (c QuestionUpdateCLIHandler) Description() string {
	return "Updates a question"
}

func (c QuestionUpdateCLIHandler) Example() string {
	return "update_question <no> <question> <answer>"
}
