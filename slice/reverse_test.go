package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			want: []int{7, 5, 3, 1},
			src:  []int{1, 3, 5, 7},
			name: "normal test",
		},
		{
			src:  []int{},
			want: []int{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []int{},
			name: "length of src is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Reverse(tt.src)
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestReverseSelf(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			want: []int{7, 5, 3, 1},
			src:  []int{1, 3, 5, 7},
			name: "normal test",
		},
		{
			src:  []int{},
			want: []int{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []int{},
			name: "length of src is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSelf(tt.src)
			assert.ElementsMatch(t, tt.want, tt.src)
		})
	}
}
