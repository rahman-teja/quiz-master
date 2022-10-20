package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestCursorClose(t *testing.T) {
	ctx := context.Background()
	csr := memorydb.NewCursorJson([][]byte{
		{},
	})

	err := csr.Close(ctx)

	assert.NoError(t, err, "[TestCursorClose] Error should nil")
}
