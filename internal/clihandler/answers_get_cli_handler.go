package clihandler

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/rahman-teja/quiz-master/internal/usecase"
	"gitlab.com/rteja-library3/rapperror"
)

type AnswersGetCLIHandler struct {
	Usecase usecase.AnswersQueryUsecase
}

func (c AnswersGetCLIHandler) Validate(params ...string) error {
	return nil
}

func (c AnswersGetCLIHandler) Handler(params ...string) {
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
	buf.WriteString("No | Question | Correct Answer | Answer | Is Correct | Point")
	buf.WriteString("\n")

	correctCount := 0
	inCorrectCount := 0
	for rr, result := range results {
		buf.WriteString(strconv.Itoa(rr + 1))
		buf.WriteString(" | ")
		buf.WriteString(result.Questions.Questions)
		buf.WriteString(" | ")
		buf.WriteString(result.Questions.Answers[0])
		buf.WriteString(" | ")
		buf.WriteString(result.Answer)
		buf.WriteString(" | ")

		if result.IsCorrect {
			buf.WriteString("Correct")
			correctCount++
		} else {
			buf.WriteString("Incorrect")
			inCorrectCount++
		}

		buf.WriteString(" | ")
		buf.WriteString(strconv.Itoa(int(result.Point)))
		buf.WriteString("\n")
	}

	buf.WriteString("Correct")
	buf.WriteString(": ")
	buf.WriteString(strconv.Itoa(correctCount))
	buf.WriteString("\n")

	buf.WriteString("Incorrect")
	buf.WriteString(": ")
	buf.WriteString(strconv.Itoa(inCorrectCount))
	buf.WriteString("\n")

	buf.WriteString("Score")
	buf.WriteString(": ")
	buf.WriteString(strconv.Itoa(int(float64(correctCount) / float64(correctCount+inCorrectCount) * 100)))
	buf.WriteString(" of 100")
	buf.WriteString("\n")

	fmt.Println(buf.String())
}

func (c AnswersGetCLIHandler) Description() string {
	return "Shows answer list"
}

func (c AnswersGetCLIHandler) Example() string {
	return "answers"
}
