package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestQuestionFindOneSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{ID: "01"})

	res := col.FindOne(ctx, "01")

	assert.NotNil(t, res, "[TestQuestionFindOneSuccess] Result should not nil")
	assert.NoError(t, res.Err(), "[TestQuestionFindOneSuccess] Error should nil")
}

func TestQuestionFindOneErrNotFound(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{ID: "01"})

	res := col.FindOne(ctx, "02")

	assert.NotNil(t, res, "[TestQuestionFindOneErrNotFound] Result should not nil")
	assert.Error(t, res.Err(), "[TestQuestionFindOneErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrNotFound, res.Err(), "[TestQuestionFindOneErrNotFound] should Error not found")
}
