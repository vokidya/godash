package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOmit(t *testing.T) {
	t.Parallel()

	demo := map[string]any{
		"test1": "test1",
		"test2": "test2",
		"test3": "test3",
	}

	result1 := Omit(demo, []string{"test1"})

	assert.Equal(t, map[string]any{
		"test3": "test3",
		"test2": "test2",
	}, result1)
}
