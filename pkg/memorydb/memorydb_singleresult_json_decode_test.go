package memorydb_test

import (
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestSingleResultDecodeSuccess(t *testing.T) {
	csr := memorydb.NewSingleResultJson([]byte(`{"id": 1}`), nil)

	err := csr.Decode(&map[string]int{})

	assert.NoError(t, err, "[TestSingleResultDecodeSuccess] Error should nil")
}

func TestSingleResultDecodeErrInvalidMarshal(t *testing.T) {
	csr := memorydb.NewSingleResultJson([]byte(`{"id": 1}`), nil)

	err := csr.Decode(&map[string]string{})

	assert.Error(t, err, "[TestSingleResultDecodeErrInvalidMarshal] should error")
	assert.Equal(t, memorydb.ErrDecodeData, err, "[TestSingleResultDecodeErrInvalidMarshal] error should decode data")
}

func TestSingleResultDecodeErr(t *testing.T) {
	csr := memorydb.NewSingleResultJson(nil, memorydb.ErrNotFound)

	err := csr.Decode(&map[string]int{})

	assert.Error(t, err, "[TestSingleResultDecodeErr] Should not error")
	assert.Equal(t, memorydb.ErrNotFound, err, "[TestSingleResultDecodeErrInvalidMarshal] error should not found")
}
