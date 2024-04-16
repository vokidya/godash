package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	t.Parallel()

	result1 := Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)

	assert.Equal(t, 4, len(result1))
	assert.Equal(t, []int{0, 1}, result1[0])
	assert.Equal(t, []int{6}, result1[3])
}
