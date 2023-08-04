// 泛型约束
package gtools

// 实数
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

// 包含了复数，大多数用不到
type Number interface {
	RealNumber | ~complex64 | ~complex128
}
