package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	t.Parallel()

	result1 := Uniq([]string{"test1", "test2", "test2", "test1"})

	assert.Equal(t, []string{
		"test1",
		"test2",
	}, result1)
}
