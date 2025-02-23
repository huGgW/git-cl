package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet[string]()
	s.Add("test")
	s.Add("burger")
	s.Add("chicken")
	assert.True(t, s.Exists("test"))
	assert.True(t, s.Exists("burger"))
	assert.False(t, s.Exists("pizza"))

	s.Remove("burger")
	assert.False(t, s.Exists("burger"))
}
