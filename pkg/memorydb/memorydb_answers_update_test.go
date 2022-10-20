package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestAnswerUpdateSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{ID: "01"})

	err := col.Update(ctx, "01", entity.Answers{ID: "01"})

	assert.NoError(t, err, "[TestAnswerUpdateSuccess] Error should nil")
}

func TestAnswerUpdateErrNotFound(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{ID: "01"})

	err := col.Update(ctx, "02", entity.Answers{ID: "02"})

	assert.Error(t, err, "[TestAnswerUpdateErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrNotFound, err, "[TestAnswerUpdateErrNotFound] should Error not found")
}
