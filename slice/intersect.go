package slice

// 功能：源切片 和 目的切片取交集,只支持 comparable 类型
// 参数：源切片src，目标切片dst
// 返回值：交集切片
func IntersectSet[T comparable](src, dst []T) []T {
	srcMap := toMap(src)
	var ret = make([]T, 0, len(src))

	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			// 取到交集
			ret = append(ret, v)
		}
	}

	return deduplicate(ret)
}

// 功能：源切片 和 目的切片取交集
// 参数：源切片src，目标切片dst
// 返回值：交集切片,已去重
func IntersectSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {

	var ret = make([]T, 0, len(src))
	for _, s := range src {
		for _, d := range dst {
			if equal(s, d) {
				ret = append(ret, s)
				break
			}
		}
	}

	return deduplicateFunc(ret, equal)
}
