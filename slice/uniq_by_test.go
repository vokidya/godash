package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqBy(t *testing.T) {
	t.Parallel()

	result1 := UniqBy([]string{"test1", "aest1", "test2", "aest2"}, func(s1 string, s2 string) bool {
		return s1[0:1] == s2[0:1]
	})

	assert.Equal(t, []string{
		"test1",
		"aest1",
	}, result1)
}
