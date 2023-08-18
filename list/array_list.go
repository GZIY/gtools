package list

import (
	"github.com/GZIY/gtools/errs"
	"github.com/GZIY/gtools/slice"
)

type ArrayList[T any] struct {
	vals []T
}

// NewArrayList 初始化一个len为0，cap为cap的 ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{vals: make([]T, 0, cap)}
}

// NewArrayListOf, 直接用ts 来初始化一个 ArrayList，不会赋值ts
func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: ts,
	}
}

func (a *ArrayList[T]) Get(idx int) (T, error) {
	l := a.Len()
	if idx < 0 || idx >= l {
		var t T
		return t, errs.NewErrIndexOutOfRange(l, idx)
	}
	return a.vals[idx], nil
}

func (a *ArrayList[T]) Append(ts ...T) error {
	a.vals = append(a.vals, ts...)
	return nil
}

func (a *ArrayList[T]) Add(idx int, t T) error {
	l := a.Len()
	if idx < 0 || idx > l {
		// 这里 idx 可以为长度，相当于 append
		return errs.NewErrIndexOutOfRange(l, idx)
	}

	// 如：vals=[1,2,3] 在 0 位置插入 100

	// 这里的追加元素主要为了 扩大 vals 长度，否则后面copy 不了
	// vals = [1,2,3,100]
	a.vals = append(a.vals, t)

	// 给idx挪位置，vals = [1,1,2,3], 注意这里就算 idx 为源vals长度3 时，也正常，vals[3+:] 返回空切片，copy内部不会执行
	copy(a.vals[idx+1:], a.vals[idx:])

	// 挪好的位置赋值 t, vals = [100,1,2,3]
	a.vals[idx] = t

	return nil
}

func (a *ArrayList[T]) Set(idx int, t T) error {
	l := a.Len()
	if idx < 0 || idx >= l {
		return errs.NewErrIndexOutOfRange(l, idx)
	}

	a.vals[idx] = t
	return nil
}

// Delete 方法会在必要的时候引起缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (a *ArrayList[T]) Delete(idx int) (T, error) {
	res, t, err := slice.Delete(a.vals, idx)
	if err != nil {
		return t, err
	}

	a.vals = res
	a.shrink()
	return t, nil
}

// 数组缩容
func (a *ArrayList[T]) shrink() {
	a.vals = slice.Shrink(a.vals)
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

func (a *ArrayList[T]) Range(fn func(idx int, t T) error) error {
	for i, v := range a.vals {
		err := fn(i, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.vals))

	copy(res, a.vals)

	return res
}
