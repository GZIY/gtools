package slice

// 功能：对一个切片进行缩容
// 参数：源切片
// 返回值：缩容后的切片
func Shrink[T any](src []T) []T {
	l, c := len(src), cap(src)

	// 根据缩容策略，获取缩容后的容量
	newCap, changed := getCap(c, l)
	if !changed {
		return src
	}

	s := make([]T, 0, newCap)
	s = append(s, src...)

	return s
}

// 功能：缩容策略
// 参数：c容量，l长度
// 返回值：缩容后的容量，是否缩容
func getCap(c, l int) (int, bool) {

	if c <= 64 {
		return c, false

	} else if c <= 2048 && (c/l) >= 4 {
		return c / 2, true

	} else if c <= 2048 {
		return c, false

	} else if (c / l) >= 2 {
		factor := 0.625 // 5/8
		return int(float32(c) * float32(factor)), true
	} else {
		return c, false
	}
}
