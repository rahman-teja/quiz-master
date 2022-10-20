package memorydb

import "fmt"

var (
	ErrNotFound       error = fmt.Errorf("err: Data not found")
	ErrDuplicate      error = fmt.Errorf("err: Data duplicate")
	ErrDecodeData     error = fmt.Errorf("err: Something wrong with decode")
	ErrInternalServer error = fmt.Errorf("err: Unkown error")
)
