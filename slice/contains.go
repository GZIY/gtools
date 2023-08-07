package slice

// 功能：判断 src 里是否存在 dst,T必须是可比较类型
// 参数：源切片src，目标值 dst
// 返回值：bool
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc(src, dst, func(src, dst T) bool {
		return src == dst
	})
}

// 功能：判断 src 里是否存在 dst,T 为any类型
// 参数：源切片src，目标值 dst，自定义相等方法
// 返回值：bool
func ContainsFunc[T any](src []T, dst T, equal equalFunc[T]) bool {
	for _, v := range src {
		if equal(v, dst) {
			return true
		}
	}

	return false
}

// 功能：判断 src 里是否存在 dst 中的任何一个元素,T 为可比较类型
// 参数：源切片src，目标切片 dst
// 返回值：bool
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap(src)
	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			return true
		}
	}

	return false
}

// 功能：判断 src 里是否存在 dst 中的任何一个元素,T 为任意类型
// 参数：源切片src，目标切片 dst， 自定义比较方法 equal
// 返回值：bool
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valDst, valSrc) {
				return true
			}
		}
	}

	return false
}

// 功能：判断 src 里是否存在 dst 中的所有元素,T 为可比较类型
// 参数：源切片src，目标切片 dst
// 返回值：bool
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap(src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}

	return true
}

// 功能：判断 src 里是否存在 dst 中的所有元素,T 为任意类型
// 参数：源切片src，目标切片 dst， 自定义比较方法 equal
// 返回值：bool
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		if !ContainsFunc(src, valDst, equal) {
			return false
		}
	}

	return true
}
