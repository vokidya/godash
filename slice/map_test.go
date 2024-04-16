package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()

	result1 := Map([]int{0, 1, 2, 3, 4}, func(_ int, item int) string {
		return strconv.Itoa(item)
	})

	assert.Equal(t, "0", result1[0])
}
