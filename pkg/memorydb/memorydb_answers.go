package memorydb

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/rahman-teja/quiz-master/internal/entity"
)

type MemoryDBAnswers struct {
	mtx   *sync.Mutex
	datas map[string][]byte
}

func NewMemoryDBAnswers() *MemoryDBAnswers {
	return &MemoryDBAnswers{
		mtx:   new(sync.Mutex),
		datas: make(map[string][]byte),
	}
}

func (m *MemoryDBAnswers) Find(ctx context.Context) (Cursor, error) {
	datas := make([][]byte, len(m.datas))
	r := 0

	for _, data := range m.datas {
		datas[r] = data
		r++
	}

	return NewCursorJson(datas), nil
}

func (m *MemoryDBAnswers) FindOne(ctx context.Context, id string) SingleResult {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	data, ok := m.datas[id]
	if !ok {
		return NewSingleResultJson(nil, ErrNotFound)
	}

	return NewSingleResultJson(data, nil)
}

func (m *MemoryDBAnswers) Create(ctx context.Context, data interface{}) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	ans, _ := data.(entity.Answers)

	// check duplicate
	_, ok := m.datas[ans.ID]
	if ok {
		return ErrDuplicate
	}

	m.datas[ans.ID], _ = json.Marshal(data)
	return nil
}

func (m *MemoryDBAnswers) Update(ctx context.Context, id string, data interface{}) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	// check not found
	_, ok := m.datas[id]
	if !ok {
		return ErrNotFound
	}

	m.datas[id], _ = json.Marshal(data)
	return nil
}

func (m *MemoryDBAnswers) Delete(ctx context.Context, id string) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	// check not found
	_, ok := m.datas[id]
	if !ok {
		return ErrNotFound
	}

	delete(m.datas, id)
	return nil
}
