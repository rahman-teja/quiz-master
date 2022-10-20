package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestAnswerFindOneSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{ID: "01"})

	res := col.FindOne(ctx, "01")

	assert.NotNil(t, res, "[TestAnswerFindOneSuccess] Result should not nil")
	assert.NoError(t, res.Err(), "[TestAnswerFindOneSuccess] Error should nil")
}

func TestAnswerFindOneErrNotFound(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{ID: "01"})

	res := col.FindOne(ctx, "02")

	assert.NotNil(t, res, "[TestAnswerFindOneErrNotFound] Result should not nil")
	assert.Error(t, res.Err(), "[TestAnswerFindOneErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrNotFound, res.Err(), "[TestAnswerFindOneErrNotFound] should Error not found")
}
