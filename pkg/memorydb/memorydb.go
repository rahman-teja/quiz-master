package memorydb

import (
	"context"
	"sync"
)

type Cursor interface {
	Next(ctx context.Context) bool
	Decode(val interface{}) error
	Close(ctx context.Context) error
}

type SingleResult interface {
	Decode(dt interface{}) error
	Err() error
}

type Collection interface {
	FindOne(ctx context.Context, id string) SingleResult
	Find(ctx context.Context) (Cursor, error)
	Create(ctx context.Context, data interface{}) error
	Update(ctx context.Context, id string, data interface{}) error
	Delete(ctx context.Context, id string) error
}

type Database interface {
	Collection(colname string) Collection
}

type MemoryDB struct {
	mtx         *sync.Mutex
	collections map[string]Collection
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		mtx: new(sync.Mutex),
		collections: map[string]Collection{
			"questions": NewMemoryDBQuestions(),
			"answers":   NewMemoryDBAnswers(),
		},
	}
}

func (m *MemoryDB) Collection(colname string) Collection {
	return m.collections[colname]
}
