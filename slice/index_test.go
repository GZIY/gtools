package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	type test struct {
		name string
		src  []int
		dst  int
		want int
	}

	tests := []test{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 0,
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: -1,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
			name: "last one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Index(tt.src, tt.dst))
		})
	}
}

func TestIndexFunc(t *testing.T) {
	type test struct {
		name string
		src  []int
		dst  int
		want int
	}

	tests := []test{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 0,
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: -1,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
			name: "last one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IndexFunc(tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestIndexLast(t *testing.T) {
	type test struct {
		name string
		src  []int
		dst  int
		want int
	}

	tests := []test{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 1,
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: -1,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: 5,
			name: "last one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IndexLast(tt.src, tt.dst))
		})
	}
}

func TestIndexLastFunc(t *testing.T) {
	type test struct {
		name string
		src  []int
		dst  int
		want int
	}

	tests := []test{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 1,
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: -1,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: 5,
			name: "last one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IndexLastFunc(tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestIndexAll(t *testing.T) {
	type test struct {
		name string
		src  []int
		dst  int
		want []int
	}

	tests := []test{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: []int{0, 1},
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: []int{},
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: []int{},
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: []int{},
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: []int{0, 5},
			name: "normal test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IndexAll(tt.src, tt.dst))
		})
	}
}

func TestIndexAllFunc(t *testing.T) {
	type test struct {
		name string
		src  []int
		dst  int
		want []int
	}

	tests := []test{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: []int{0, 1},
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: []int{},
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: []int{},
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: []int{},
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: []int{0, 5},
			name: "normal test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IndexAllFunc(tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}
