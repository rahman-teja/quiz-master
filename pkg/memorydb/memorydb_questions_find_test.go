package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestQuestionFindSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBQuestions()

	col.Create(ctx, entity.Questions{})
	csr, err := col.Find(ctx)

	assert.NotNil(t, csr, "[TestQuestionFindSuccess] Cursor should not nil")
	assert.NoError(t, err, "[TestQuestionFindSuccess] should not error")
}
