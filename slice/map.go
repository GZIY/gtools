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

// 功能：切片去重，不保证顺序不变
// 参数：可比较类型的 切片
// 返回值：去重后的切片
func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	var newData = make([]T, 0, len(data))
	for i, v := range data {
		// 后面的元素 都 跟当前元素不重复，则认为当前元素是不重复的元素
		if !ContainsFunc(data[i+1:], v, equal) {
			newData = append(newData, v)
		}
	}

	return newData
}

// 功能：切片去重，不保证顺序不变
// 参数：可比较类型的 切片
// 返回值：去重后的切片
func deduplicate[T comparable](data []T) []T {
	dataMap := toMap(data)
	var newData = make([]T, 0, len(dataMap))
	for k, _ := range dataMap {
		newData = append(newData, k)
	}

	return newData
}
