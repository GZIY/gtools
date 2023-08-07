package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	type testCase struct {
		name string
		src  []int
		dst  int
		want bool
	}

	testCases := []testCase{
		{
			name: "dst exist",
			src:  []int{1, 4, 6, 2},
			dst:  4,
			want: true,
		},
		{
			name: "dst not exist",
			src:  []int{1, 3, 6, 2},
			dst:  4,
			want: false,
		},
		{
			name: "len src 0",
			src:  []int{},
			dst:  4,
			want: false,
		},
		{
			name: "src nil",
			src:  nil,
			dst:  4,
			want: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Contains(test.src, test.dst))
		})
	}
}

func TestContainsFunc(t *testing.T) {
	type testCase struct {
		name string
		src  []int
		dst  int
		want bool
	}

	testCases := []testCase{
		{
			name: "dst exist",
			src:  []int{1, 4, 6, 2},
			dst:  4,
			want: true,
		},
		{
			name: "dst not exist",
			src:  []int{1, 3, 6, 2},
			dst:  4,
			want: false,
		},
		{
			name: "len src 0",
			src:  []int{},
			dst:  4,
			want: false,
		},
		{
			name: "src nil",
			src:  nil,
			dst:  4,
			want: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsFunc(test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestContainsAny(t *testing.T) {
	type testCase struct {
		name string
		src  []int
		dst  []int
		want bool
	}

	testCases := []testCase{
		{
			name: "dst exist",
			src:  []int{1, 4, 6, 2},
			dst:  []int{1, 6},
			want: true,
		},
		{
			name: "dst not exist",
			src:  []int{1, 3, 6, 2},
			dst:  []int{5, 7},
			want: false,
		},
		{
			name: "len src 0",
			src:  []int{},
			dst:  []int{1},
			want: false,
		},
		{
			name: "src nil",
			src:  nil,
			dst:  []int{1},
			want: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAny(test.src, test.dst))
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	type testCase struct {
		name string
		src  []int
		dst  []int
		want bool
	}

	testCases := []testCase{
		{
			name: "dst exist",
			src:  []int{1, 4, 6, 2},
			dst:  []int{1, 6},
			want: true,
		},
		{
			name: "dst not exist",
			src:  []int{1, 3, 6, 2},
			dst:  []int{5, 7},
			want: false,
		},
		{
			name: "len src 0",
			src:  []int{},
			dst:  []int{1},
			want: false,
		},
		{
			name: "src nil",
			src:  nil,
			dst:  []int{1},
			want: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAnyFunc(test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestContainsAll(t *testing.T) {
	type testCase struct {
		name string
		src  []int
		dst  []int
		want bool
	}

	testCases := []testCase{
		{
			name: "dst all exist",
			src:  []int{1, 4, 6, 2},
			dst:  []int{1, 6},
			want: true,
		},
		{
			name: "dst not all exist",
			src:  []int{1, 3, 6, 2},
			dst:  []int{1, 6, 5, 7},
			want: false,
		},
		{
			name: "len src 0",
			src:  []int{},
			dst:  []int{1},
			want: false,
		},
		{
			name: "src nil dst empty",
			src:  nil,
			dst:  []int{},
			want: true,
		},
		{
			name: "src nil dst nil",
			src:  nil,
			dst:  nil,
			want: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAll(test.src, test.dst))
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	type testCase struct {
		name string
		src  []int
		dst  []int
		want bool
	}

	testCases := []testCase{
		{
			name: "dst all exist",
			src:  []int{1, 4, 6, 2},
			dst:  []int{1, 6},
			want: true,
		},
		{
			name: "dst not all exist",
			src:  []int{1, 3, 6, 2},
			dst:  []int{1, 6, 5, 7},
			want: false,
		},
		{
			name: "len src 0",
			src:  []int{},
			dst:  []int{1},
			want: false,
		},
		{
			name: "src nil dst empty",
			src:  nil,
			dst:  []int{},
			want: true,
		},
		{
			name: "src nil dst nil",
			src:  nil,
			dst:  nil,
			want: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAllFunc(test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}
