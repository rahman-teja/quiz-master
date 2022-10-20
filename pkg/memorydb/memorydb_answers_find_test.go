package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestAnswerFindSuccess(t *testing.T) {
	ctx := context.Background()

	col := memorydb.NewMemoryDBAnswers()

	col.Create(ctx, entity.Answers{})
	csr, err := col.Find(ctx)

	assert.NotNil(t, csr, "[TestAnswerFindSuccess] Cursor should not nil")
	assert.NoError(t, err, "[TestAnswerFindSuccess] should not error")
}
