package slice

// 功能：找到源切片在目的切片中不存在的元素 和 找到目的切片在源切片中不存在的元素，结果会去重，返回结果的顺序是不确定的,只支持 comparable 类型
// 参数：源切片src，目标切片dst
// 返回值：对称差集切片 (src-dst) U (dst-src)
func SymmetricDiffSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap(src), toMap(dst)

	for dstKey, _ := range dstMap {
		if _, exist := srcMap[dstKey]; exist {
			// 两个切片都有的元素，从俩map中删除，剩下的俩map是互相不存在的元素
			delete(dstMap, dstKey)
			delete(srcMap, dstKey)
		}
	}

	for k, v := range dstMap {
		srcMap[k] = v
	}

	var res = make([]T, 0, len(srcMap))
	for k, _ := range srcMap {
		res = append(res, k)
	}

	return res
}

// 功能：找到源切片在目的切片中不存在的元素 和 找到目的切片在源切片中不存在的元素，结果会去重，返回结果的顺序是不确定的
// 参数：源切片src，目标切片dst
// 返回值：对称差集切片 (src-dst) U (dst-src)
func SymmetricDiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	// 找交集
	interSection := IntersectSetFunc(src, dst, equal)

	// 容量一定 <= len(src) + len(dst) - len(interSection) * 2
	var res = make([]T, 0, len(src)+len(dst)-len(interSection)*2)

	// 遍历 src dst ，不包含交集元素的就是 对称差集元素
	for _, v := range src {
		if !ContainsFunc(interSection, v, equal) {
			res = append(res, v)
		}
	}
	for _, v := range dst {
		if !ContainsFunc(interSection, v, equal) {
			res = append(res, v)
		}
	}

	return deduplicateFunc(res, equal)
}
