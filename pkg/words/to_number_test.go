package words_test

import (
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/words"
	"github.com/stretchr/testify/assert"
)

func TestToNumberSuccess(t *testing.T) {
	res, err := words.ToNumber("minus one hundred thousand")

	assert.Equal(t, -100_000, res, `ToNumber should be 1`)
	assert.NoError(t, err, "ToNumber should not error")
}

func TestToNumberErrEmptyWord(t *testing.T) {
	res, err := words.ToNumber("")

	assert.Equal(t, 0, res, `ToNumber should be 0`)
	assert.Error(t, err, "ToNumber should error")
}

func TestToNumberErrInvalidWord(t *testing.T) {
	res, err := words.ToNumber("minus rahman")

	assert.Equal(t, 0, res, `ToNumber should be 0`)
	assert.Error(t, err, "ToNumber should error")
}
