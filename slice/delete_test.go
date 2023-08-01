package slice

import (
	"github.com/GZIY/gtools/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

type typeCase struct {
	name      string
	slice     []int
	index     int
	wantSlice []int
	wantVal   int
	wantErr   error
}

func TestDelete(t *testing.T) {
	// 若干 case 例子
	testCases := []typeCase{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			index:     0,
			wantSlice: []int{100},
			wantVal:   123,
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			index:     1,
			wantSlice: []int{123, 125},
			wantVal:   124,
		},
		{
			name:    "index out of range",
			slice:   []int{123, 100},
			index:   12,
			wantErr: errs.NewErrIndexOutOfRange(2, 12),
		},
		{
			name:    "index less than 0",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:      "index last",
			slice:     []int{123, 100, 101, 102, 102, 102},
			index:     5,
			wantSlice: []int{123, 100, 101, 102, 102},
			wantVal:   102,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, val)
			assert.Equal(t, tc.wantSlice, res)
		})
	}

}