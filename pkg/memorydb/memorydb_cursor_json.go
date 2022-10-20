package memorydb

import (
	"context"
	"encoding/json"
	"errors"
)

var (
	ErrInvalidCursor error = errors.New("err: Invalid cursor")
)

type CursorJson struct {
	datas  [][]byte
	cursor int
}

func NewCursorJson(datas [][]byte) *CursorJson {
	return &CursorJson{
		datas:  datas,
		cursor: -1,
	}
}

func (s *CursorJson) Next(ctx context.Context) bool {
	if s.cursor >= len(s.datas)-1 || len(s.datas) == 0 {
		return false
	}

	s.cursor++

	return true
}

func (s *CursorJson) Close(ctx context.Context) error {
	s.cursor = -1
	s.datas = nil
	return nil
}

func (s *CursorJson) Decode(dt interface{}) error {
	if s.cursor < 0 {
		return ErrInvalidCursor
	}

	err := json.Unmarshal(s.datas[s.cursor], dt)
	if err != nil {
		return ErrDecodeData
	}

	return nil
}
