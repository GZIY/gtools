package slice

// 功能：切片转map
// 参数：可比较类型的 切片
// 返回值：map
func toMap[T comparable](src []T) map[T]struct{} {
	var m = make(map[T]struct{}, len(src))
	for _, v := range src {
		// 使用空结构体,减少内存消耗
		m[v] = struct{}{}
	}

	return m
}
