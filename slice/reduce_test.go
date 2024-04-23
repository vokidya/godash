package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Parallel()

	dataList := []string{"a", "b", "c", "d", "e"}

	result1 := Reduce(dataList, func(prev string, _ int, item string) string {
		return prev + item
	}, "")

	assert.Equal(t, "abcde", result1)
}
