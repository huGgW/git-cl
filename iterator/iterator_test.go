package iterator

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	filterSeq := Filter(seq,
		func(v int) bool { return v%2 != 0 },
		func(v int) bool { return v%3 == 0 },
	)

	assert.Equal(t, []int{3, 9}, slices.Collect(filterSeq))
}

func TestMap(t *testing.T) {
    seq := slices.Values([]int{1, 2, 3, 4, 5})
    mapSeq1 := Map(seq, func(v int) int { return v * 2 })
    mapSeq2 := Map(mapSeq1, func(v int) string { return strconv.Itoa(v) })

    assert.Equal(t, []string{"2", "4", "6", "8", "10"}, slices.Collect(mapSeq2))
}
