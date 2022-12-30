package slicex

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestArrayHas(t *testing.T) {
	tests := []struct {
		want  bool
		array []int
		has   int
		name  string
	}{
		{
			want:  false,
			array: []int{1, 3, 5, 7, 8},
			has:   2,
			name:  "not exist",
		},
		{
			want:  true,
			array: []int{1, 3, 5, 7, 8},
			has:   5,
			name:  "exist",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ArrayHas[int](test.array, test.has))
		})
	}
}
