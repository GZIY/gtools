package slice

import "github.com/GZIY/gtools/errs"

// 功能：删除元素指定位置的下标
// 参数：源切片src，索引idx
// 返回值：删除下标之后的切片，删除下标对应的元素，错误信息
func Delete[T any](src []T, idx int) ([]T, T, error) {
	length := len(src)
	if idx < 0 || idx >= length {
		// 索引异常返回T类型的零值
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, idx)
	}

	// 从删除位置遍历切片，将后续元素往前覆盖一个位置
	res := src[idx]
	for i := idx; i < len(src)-1; i++ {
		src[idx] = src[idx+1]
	}

	return src[:len(src)-1], res, nil
}
