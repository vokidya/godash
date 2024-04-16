package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	t.Parallel()

	result := Difference([]int{0, 1, 2}, []int{1, 4})

	assert.Equal(t, []int{0, 2}, result)
}
