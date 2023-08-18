package list

type List[T any] interface {
	// 返回idx 位置的元素, 如果下标超出范围，应该返回错误
	Get(idx int) (T, error)

	// 末尾追加元素
	Append(ts ...T) error

	// 在 idx 位置处增加新元素,如果下标超出范围，应该返回错误
	Add(idx int, t T) error

	// Set 设置 index 位置的值，如果下标超出范围，应该返回错误
	Set(idx int, t T) error

	// Delete 删除目标元素的位置，并且返回该位置的值，如果 index 超出下标，应该返回错误
	Delete(idx int) (T, error)

	// 返回长度
	Len() int

	// 返回容量
	Cap() int

	// Range 遍历 List 的所有元素， f为对遍历的索引和值进行的 操作
	Range(fn func(idx int, t T) error) error

	// AsSlice 将 List 转化为一个切片,不允许返回nil，在没有元素的情况下，必须返回一个长度和容量都为 0 的切片
	// AsSlice 每次调用都必须返回一个全新的切片
	AsSlice() []T
}
