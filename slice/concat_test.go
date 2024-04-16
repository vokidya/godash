package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	t.Parallel()

	result := Concat([]int{0, 1, 2}, []int{3, 4, 5, 6})

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6}, result)
}
