package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	result1 := Filter([]int{0, 1, 2, 3, 4}, func(_ int, item int) bool {
		return item == 2
	})

	assert.Equal(t, 1, len(result1))
	assert.Equal(t, 2, result1[0])
}
