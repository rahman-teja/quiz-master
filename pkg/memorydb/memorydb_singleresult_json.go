package memorydb

import (
	"encoding/json"
)

type SingleResultJson struct {
	datas []byte
	err   error
}

func NewSingleResultJson(datas []byte, err error) *SingleResultJson {
	return &SingleResultJson{
		datas: datas,
		err:   err,
	}
}

func (s SingleResultJson) Decode(dt interface{}) error {
	if s.err != nil {
		return s.err
	}

	err := json.Unmarshal(s.datas, dt)
	if err != nil {
		s.err = ErrDecodeData
	}

	return s.err
}

func (s SingleResultJson) Err() error {
	return s.err
}
