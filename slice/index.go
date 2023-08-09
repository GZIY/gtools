package slice

// 功能：找到 和 dst 相等的 第一个元素的下标,没找到返回 -1
// 参数：源切片src，目标值dst
// 返回值：和 dst 相等的 第一个元素的下标
func Index[T comparable](src []T, dst T) int {
	return IndexFunc(src, dst, func(src, dst T) bool {
		return src == dst
	})
}

// 功能：找到 和 dst 相等的 第一个元素的下标,没找到返回 -1
// 参数：源切片src，目标值dst
// 返回值：和 dst 相等的 第一个元素的下标
func IndexFunc[T any](src []T, dst T, equal equalFunc[T]) int {
	for i, v := range src {
		if equal(v, dst) {
			return i
		}
	}

	return -1
}

// 功能：找到 和 dst 相等的 最后一个元素的下标,没找到返回 -1
// 参数：源切片src，目标值dst
// 返回值：和 dst 相等的 第一个元素的下标
func IndexLast[T comparable](src []T, dst T) int {
	return IndexLastFunc(src, dst, func(src, dst T) bool {
		return src == dst
	})
}

// 功能：找到 和 dst 相等的 最后一个元素的下标,没找到返回 -1
// 参数：源切片src，目标值dst
// 返回值：和 dst 相等的 第一个元素的下标
func IndexLastFunc[T any](src []T, dst T, equal equalFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if equal(src[i], dst) {
			return i
		}
	}

	return -1
}

// 功能：找到 和 dst 相等的 所有元素的下标,没找到返回 -1
// 参数：源切片src，目标值dst
// 返回值：和 dst 相等的 第一个元素的下标
func IndexAll[T comparable](src []T, dst T) []int {
	return IndexAllFunc(src, dst, func(src, dst T) bool {
		return src == dst
	})
}

// 功能：找到 和 dst 相等的 所有元素的下标,没找到返回 -1
// 参数：源切片src，目标值dst
// 返回值：和 dst 相等的 第一个元素的下标
func IndexAllFunc[T any](src []T, dst T, equal equalFunc[T]) []int {
	var indexes = make([]int, 0, len(src))

	for i, v := range src {
		if equal(v, dst) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}
