package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestQuestionUpdateSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{ID: "01"})

	err := col.Update(ctx, "01", entity.Questions{ID: "01"})

	assert.NoError(t, err, "[TestQuestionUpdateSuccess] Error should nil")
}

func TestQuestionUpdateErrNotFound(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{ID: "01"})

	err := col.Update(ctx, "02", entity.Questions{ID: "02"})

	assert.Error(t, err, "[TestQuestionUpdateErrNotFound] should Error")
	assert.Equal(t, memorydb.ErrNotFound, err, "[TestQuestionUpdateErrNotFound] should Error not found")
}
