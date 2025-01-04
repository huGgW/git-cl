package scanner

import (
	"bufio"
	"slices"
	"strings"
	"testing"

	"github.com/huGgW/git-cl/pkg/iterator"
	"github.com/stretchr/testify/assert"
)

func TestLineSeq(t *testing.T) {
	sc := bufio.NewScanner(strings.NewReader("a\nb\nc"))
	seq := LineSeq(sc)

	assert.Equal(t,
		slices.Collect(iterator.First(seq)),
		[]string{"a", "b", "c"},
	)
}
