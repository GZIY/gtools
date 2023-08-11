package slice

// 功能：两个切片取并集,只支持comparable
// 参数：源切片src，目的切片dst
// 返回值：取并集后的切片，已去重，返回顺序不定
func UnionSet[T comparable](src, dst []T) []T {
	srcMap := toMap(src)
	dstMap := toMap(dst)

	for k, _ := range srcMap {
		dstMap[k] = struct{}{}
	}

	var res = make([]T, 0, len(dstMap))
	for k, _ := range dstMap {
		res = append(res, k)
	}

	return res
}

// 功能：两个切片取并集
// 参数：源切片src，目的切片dst, 自定义比较方法
// 返回值：取并集后的切片，已去重，返回顺序不定
func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var res = make([]T, 0, len(src)+len(dst))
	res = append(res, src...)
	res = append(res, dst...)

	return deduplicateFunc(res, equal)
}
