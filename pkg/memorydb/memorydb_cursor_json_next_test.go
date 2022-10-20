package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/memorydb"
)

func TestCursorNext(t *testing.T) {
	ctx := context.Background()
	csr := memorydb.NewCursorJson([][]byte{
		{},
	})

	csr.Next(ctx)
	csr.Next(ctx)
}
