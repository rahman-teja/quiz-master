package cli_test

import (
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/cli"
	"github.com/stretchr/testify/assert"
)

func TestIsQuote(t *testing.T) {
	assert.True(t, cli.IsQuote('"'), "[TestIsQuote] Should \"true\"")
	assert.False(t, cli.IsQuote('A'), "[TestIsQuote] Should \"false\"")
}

func TestBuildCommand(t *testing.T) {
	cmds, err := cli.BuildCommand("c \"nama saya\" teja")

	assert.Len(t, cmds, 3, "[TestBuildCommand] Command should 3 length")
	assert.NoError(t, err, "[TestBuildCommand] Should not error")
}

func TestBuildCommandErr(t *testing.T) {
	cmds, err := cli.BuildCommand("")

	assert.Len(t, cmds, 0, "[TestBuildCommandErr] Command should 0 length")
	assert.Error(t, err, "[TestBuildCommandErr] Should error")
}
