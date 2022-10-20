package clihandler

import (
	"context"
	"fmt"
	"strings"

	"github.com/rahman-teja/quiz-master/internal/usecase"
	"gitlab.com/rteja-library3/rapperror"
)

type QuestionGetCLIHandler struct {
	Usecase usecase.QuestionsQueryUsecase
}

func (c QuestionGetCLIHandler) Validate(params ...string) error {
	return nil
}

func (c QuestionGetCLIHandler) Handler(params ...string) {
	ctx := context.Background()

	results, _, err := c.Usecase.Get(ctx)
	if err != nil {
		rapp, ok := err.(*rapperror.AppError)
		if ok {
			fmt.Println(rapp.Message)
			return
		}

		fmt.Println(err.Error())
		return
	}

	buf := new(strings.Builder)
	buf.WriteString("No | Question | Answer")
	buf.WriteString("\n")

	for _, result := range results {
		buf.WriteString(result.ID)
		buf.WriteString(" | ")
		buf.WriteString(result.Questions)
		buf.WriteString(" | ")
		buf.WriteString(result.Answers[0])
		buf.WriteString("\n")
	}

	fmt.Println(buf.String())
}

func (c QuestionGetCLIHandler) Description() string {
	return "Shows question list"
}

func (c QuestionGetCLIHandler) Example() string {
	return "questions"
}
