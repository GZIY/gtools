package errs

import "fmt"

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(lenth int, idx int) error {
	return fmt.Errorf("gtools : 下标超出范围，长度 %d，下标 %d", lenth, idx)
}

// NewErrEmptySlice 创建一个代表空切片异常的错误
func NewErrEmptySlice() error {
	return fmt.Errorf("gtools : 空切片异常")
}
