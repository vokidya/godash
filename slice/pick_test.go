package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPick(t *testing.T) {
	t.Parallel()

	demo := map[string]any{
		"test1": "test1",
		"test2": "test2",
		"test3": "test3",
	}

	result1 := Pick(demo, []string{"test1", "test2"})

	assert.Equal(t, map[string]any{
		"test1": "test1",
		"test2": "test2",
	}, result1)
}
