package memorydb_test

import (
	"context"
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestCursorDecodeSuccess(t *testing.T) {
	ctx := context.Background()
	csr := memorydb.NewCursorJson([][]byte{
		[]byte(`{"id": 1}`),
	})

	csr.Next(ctx)

	err := csr.Decode(&map[string]int{})

	assert.NoError(t, err, "[TestCursorDecodeSuccess] Error should nil")
}

func TestCursorDecodeErrInvalidCursor(t *testing.T) {
	csr := memorydb.NewCursorJson([][]byte{
		[]byte(`{"id": 1}`),
	})

	err := csr.Decode(&map[string]int{})

	assert.Error(t, err, "[TestCursorDecodeErrInvalidCursor] should error")
	assert.Equal(t, memorydb.ErrInvalidCursor, err, "[TestCursorDecodeErrInvalidCursor] error should invalid cursor")
}

func TestCursorDecodeErrInvalidMarshal(t *testing.T) {
	ctx := context.Background()
	csr := memorydb.NewCursorJson([][]byte{
		[]byte(`{"id": 1}`),
	})

	csr.Next(ctx)

	err := csr.Decode(&map[string]string{})

	assert.Error(t, err, "[TestCursorDecodeErrInvalidMarshal] should error")
	assert.Equal(t, memorydb.ErrDecodeData, err, "[TestCursorDecodeErrInvalidMarshal] error should decode data")
}
