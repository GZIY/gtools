package slice

// 功能：反转一个切片，返回新切片，源切片不变
// 参数：源切片
// 返回值：反转后的新切片
func Reverse[T comparable](src []T) []T {
	var res = make([]T, len(src), len(src))

	for i := len(src) - 1; i >= 0; i-- {
		res[i] = src[i]
	}

	return res
}

// 功能：在源切片上反转
// 参数：源切片
// 返回值：无
func ReverseSelf[T comparable](src []T) {

	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}

}
