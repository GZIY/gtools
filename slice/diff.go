package slice

// 功能：找到源切片在目的切片中不存在的元素，结果会去重，返回结果的顺序是不确定的,只支持 comparable 类型
// 参数：源切片src，目标切片dst
// 返回值：差集切片 即 src - dst
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap(src)
	for _, v := range dst {
		delete(srcMap, v)
	}

	var res = make([]T, 0, len(srcMap))
	for k, _ := range srcMap {
		res = append(res, k)
	}

	return res
}

// 功能：找到源切片在目的切片中不存在的元素，结果会去重，返回结果的顺序是不确定的
// 参数：源切片src，目标切片dst
// 返回值：差集切片
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var res = make([]T, 0, len(src))
	for _, v := range src {
		if !ContainsFunc(dst, v, equal) {
			res = append(res, v)
		}
	}

	// 切片去重
	return deduplicateFunc(res, equal)
}
