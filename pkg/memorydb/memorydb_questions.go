package memorydb

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/rahman-teja/quiz-master/internal/entity"
)

type MemoryDBQuestions struct {
	mtx   *sync.Mutex
	datas map[string][]byte
}

func NewMemoryDBQuestions() *MemoryDBQuestions {
	return &MemoryDBQuestions{
		mtx:   new(sync.Mutex),
		datas: make(map[string][]byte),
	}
}

func (m *MemoryDBQuestions) Find(ctx context.Context) (Cursor, error) {
	datas := make([][]byte, len(m.datas))
	r := 0

	for _, data := range m.datas {
		datas[r] = data
		r++
	}

	return NewCursorJson(datas), nil
}

func (m *MemoryDBQuestions) FindOne(ctx context.Context, id string) SingleResult {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	data, ok := m.datas[id]
	if !ok {
		return NewSingleResultJson(nil, ErrNotFound)
	}

	return NewSingleResultJson(data, nil)
}

func (m *MemoryDBQuestions) Create(ctx context.Context, data interface{}) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	qst, _ := data.(entity.Questions)

	// check duplicate
	_, ok := m.datas[qst.ID]
	if ok {
		return ErrDuplicate
	}

	m.datas[qst.ID], _ = json.Marshal(data)
	return nil
}

func (m *MemoryDBQuestions) Update(ctx context.Context, id string, data interface{}) error {
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

func (m *MemoryDBQuestions) Delete(ctx context.Context, id string) error {
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
