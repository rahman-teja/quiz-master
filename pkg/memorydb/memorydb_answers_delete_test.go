package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestAnswerDeleteSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{ID: "01"})

	err := col.Delete(ctx, "01")

	assert.NoError(t, err, "[TestAnswerDeleteSuccess] Error should nil")
}

func TestAnswerDeleteErrNotFound(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{ID: "01"})

	err := col.Delete(ctx, "02")

	assert.Error(t, err, "[TestAnswerDeleteErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrNotFound, err, "[TestAnswerDeleteErrNotFound] should Error not found")
}
