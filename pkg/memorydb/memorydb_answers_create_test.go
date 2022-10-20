package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestAnswerCreateSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	err := col.Create(ctx, entity.Answers{ID: "01"})

	assert.NoError(t, err, "[TestAnswerCreateSuccess] Error should nil")
}

func TestAnswerCreateErrDuplicate(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{ID: "01"})

	err := col.Create(ctx, entity.Answers{ID: "01"})

	assert.Error(t, err, "[TestAnswerCreateErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrDuplicate, err, "[TestAnswerCreateErrNotFound] should Error duplicate")
}
