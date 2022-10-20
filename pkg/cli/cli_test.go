package cli_test

import (
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/cli"
	"github.com/stretchr/testify/assert"
)

func TestNewCli(t *testing.T) {
	assert.NotNil(t, cli.NewCli(), "[TestNewCli] Should not nil")
}
