package slice

import (
	"github.com/GZIY/gtools"
	"github.com/GZIY/gtools/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	type testCase[T gtools.RealNumber] struct {
		name    string
		input   []T
		wantVal T
		wantErr error
	}

	testCases := []testCase[int64]{
		{
			name:    "int64",
			input:   []int64{1, 2, 3, 4},
			wantVal: 4,
			wantErr: nil,
		},
		{
			name:    "empty",
			input:   []int64{},
			wantVal: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
		{
			name:    "nil",
			input:   nil,
			wantVal: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Max(tc.input)

			assert.Equal(t, res, tc.wantVal)
			assert.Equal(t, err, tc.wantErr)
		})
	}
}

func TestMin(t *testing.T) {
	type testCase[T gtools.RealNumber] struct {
		name    string
		input   []T
		wantVal T
		wantErr error
	}

	testCases := []testCase[int64]{
		{
			name:    "int64",
			input:   []int64{1, 2, 3, 4},
			wantVal: 1,
			wantErr: nil,
		},
		{
			name:    "empty",
			input:   []int64{},
			wantVal: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
		{
			name:    "nil",
			input:   nil,
			wantVal: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Min(tc.input)

			assert.Equal(t, res, tc.wantVal)
			assert.Equal(t, err, tc.wantErr)
		})
	}
}

func TestSum(t *testing.T) {
	type testCase[T gtools.RealNumber] struct {
		name    string
		input   []T
		wantVal T
		wantErr error
	}

	testCases := []testCase[int64]{
		{
			name:    "int64",
			input:   []int64{1, 2, 3, 4},
			wantVal: 10,
			wantErr: nil,
		},
		{
			name:    "empty",
			input:   []int64{},
			wantVal: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
		{
			name:    "nil",
			input:   nil,
			wantVal: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum(tc.input)

			assert.Equal(t, res, tc.wantVal)
		})
	}
}
