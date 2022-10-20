package cli_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/cli"
	"github.com/rahman-teja/quiz-master/pkg/cli/mocks"
	"github.com/stretchr/testify/mock"
)

func TestStart(t *testing.T) {
	cli := cli.NewCliWithReader(strings.NewReader("\nhelp\nanswers\nquestions\nqq\nexit"))

	ansMocks := new(mocks.Handler)
	ansMocks.On("Example").Return("answers")
	ansMocks.On("Description").Return("Shows answer list")
	ansMocks.On("Validate", mock.Anything).Return(nil)
	ansMocks.On("Handler", mock.Anything)

	qstMocks := new(mocks.Handler)
	qstMocks.On("Example").Return("questions")
	qstMocks.On("Description").Return("Shows questions list")
	qstMocks.On("Validate", mock.Anything).Return(errors.New("err"))
	qstMocks.On("Handler", mock.Anything)

	cli.Register("help", nil)
	cli.Register("answers", ansMocks)
	cli.Register("questions", qstMocks)

	cli.Start()

	ansMocks.AssertExpectations(t)
}
