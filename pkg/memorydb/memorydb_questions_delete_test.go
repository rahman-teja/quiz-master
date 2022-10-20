package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestQuestionDeleteSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{ID: "01"})

	err := col.Delete(ctx, "01")

	assert.NoError(t, err, "[TestQuestionDeleteSuccess] Error should nil")
}

func TestQuestionDeleteErrNotFound(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{ID: "01"})

	err := col.Delete(ctx, "02")

	assert.Error(t, err, "[TestQuestionDeleteErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrNotFound, err, "[TestQuestionDeleteErrNotFound] should Error not found")
}
