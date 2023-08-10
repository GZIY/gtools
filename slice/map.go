package slice

// 功能：切片 对 每一个元素都 进行一个 操作，转为另一个切片
// 参数：源切片src， 每个元素的操作方法m
// 返回值：转换完的 切片 dst
func Map[Src, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src), len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}

	return dst
}

// 功能：切片 对 每一个元素都 进行一个 操作，转为另一个切片 ,支持对某些元素的过滤功能
// 参数：源切片src， 每个元素的操作方法m, 如果m方法返回false，表示要过滤掉改元素
// 返回值：转换完的 切片 dst
func FilterMap[Src, Dst any](src []Src, m func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for i, s := range src {
		dst, ok := m(i, s)
		if ok {
			// ok 则不过滤该元素
			res = append(res, dst)
		}
	}

	return res
}

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
