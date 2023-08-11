package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymmetricDiffSet(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			src:  []int{1, 2, 4, 3},
			dst:  []int{4, 5, 6, 1},
			want: []int{2, 3, 5, 6},
			name: "normal test",
		},
		{
			src:  []int{1, 1, 2, 3, 4},
			dst:  []int{4, 5, 6, 1, 7, 6},
			want: []int{3, 6, 7, 5, 2},
			name: "deduplicate",
		},
		{
			src:  []int{},
			dst:  []int{1},
			want: []int{1},
			name: "src length is 0",
		},
		{
			src:  []int{1, 3, 5},
			dst:  []int{2, 4},
			want: []int{1, 3, 2, 4, 5},
			name: "not exist same ele",
		},
		{
			name: "both nil",
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := SymmetricDiffSet(tt.src, tt.dst)
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestSymmetricDiffSetFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			src:  []int{1, 2, 4, 3},
			dst:  []int{4, 5, 6, 1},
			want: []int{2, 3, 5, 6},
			name: "normal test",
		},
		{
			src:  []int{1, 1, 2, 3, 4},
			dst:  []int{4, 5, 6, 1, 7, 6},
			want: []int{3, 6, 7, 5, 2},
			name: "deduplicate",
		},
		{
			src:  []int{},
			dst:  []int{1},
			want: []int{1},
			name: "src length is 0",
		},
		{
			src:  []int{1, 3, 5},
			dst:  []int{2, 4},
			want: []int{1, 3, 2, 4, 5},
			name: "not exist same ele",
		},
		{
			name: "both nil",
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := SymmetricDiffSetFunc(tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}
