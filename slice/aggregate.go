package slice

import (
	"github.com/GZIY/gtools"
	"github.com/GZIY/gtools/errs"
)

// 聚合函数
// Max返回最大值
func Max[T gtools.RealNumber](s []T) (T, error) {
	var res T
	if len(s) == 0 {
		// 空切片返回T类型的零值
		return res, errs.NewErrEmptySlice()
	}

	res = s[0]

	for i := 1; i < len(s); i++ {
		if s[i] > res {
			res = s[i]
		}
	}
	return res, nil
}

// Min返回最小值
func Min[T gtools.RealNumber](s []T) (T, error) {
	var res T
	if len(s) == 0 {
		// 空切片返回T类型的零值
		return res, errs.NewErrEmptySlice()
	}

	res = s[0]

	for i := 1; i < len(s); i++ {
		if s[i] < res {
			res = s[i]
		}
	}
	return res, nil
}

// sum 求和
func Sum[T gtools.Number](s []T) T {
	var res T
	for _, v := range s {
		res += v
	}

	return res
}
