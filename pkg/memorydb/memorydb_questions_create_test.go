package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestQuestionCreateSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	err := col.Create(ctx, entity.Questions{ID: "01"})

	assert.NoError(t, err, "[TestQuestionCreateSuccess] Error should nil")
}

func TestQuestionCreateErrDuplicate(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{ID: "01"})

	err := col.Create(ctx, entity.Questions{ID: "01"})

	assert.Error(t, err, "[TestQuestionCreateErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrDuplicate, err, "[TestQuestionCreateErrNotFound] should Error duplicate")
}
