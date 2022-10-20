package memorydb_test

import (
	"testing"

	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestNewMemoryDB(t *testing.T) {
	db := memorydb.NewMemoryDB()

	assert.NotNil(t, db, "[TestNewMemoryDB] db should not nil")
}

func TestCollectionMemoryDB(t *testing.T) {
	db := memorydb.NewMemoryDB()
	col := db.Collection("answers")

	assert.NotNil(t, col, "[TestCollectionMemoryDB] Collection should not nil")
}

func TestCollectionMemoryDBIsNil(t *testing.T) {
	db := memorydb.NewMemoryDB()
	col := db.Collection("answer")

	assert.Nil(t, col, "[TestCollectionMemoryDBIsNil] Collection should nil")
}
