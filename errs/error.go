package errs

import "fmt"

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(lenth int, idx int) error {
	return fmt.Errorf("gtools : 下标超出范围，长度 %d，下标 %d", lenth, idx)
}
