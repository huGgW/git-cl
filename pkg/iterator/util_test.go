package iterator

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	// Test doubling integers in an array
	numbers := []int{1, 2, 3, 4, 5}
	var result []int

	mapper := Map(
		slices.Values(numbers),
		func(x int) int {
			return x * 2
		},
	)

	for v := range mapper {
		result = append(result, v)
	}

	assert.Equal(t, []int{2, 4, 6, 8, 10}, result, "mapped slice should match expected values")
}

func TestFilter(t *testing.T) {
	// Test filtering even numbers
	numbers := []int{1, 2, 3, 4, 5, 6}
	var result []int

	filtered := Filter(
		slices.Values(numbers),
		func(x int) bool {
			return x%2 == 0
		},
	)

	for v := range filtered {
		result = append(result, v)
	}

	assert.Equal(t, []int{2, 4, 6}, result, "filtered slice should contain only even numbers")
}

func TestMap2(t *testing.T) {
	// Test transforming key-value pairs
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	result := make(map[string]string)

	transformed := Map2(
		maps.All(input),
		func(k string, v int) (string, string) {
			return k + k, string(rune('A' + v - 1))
		},
	)

	for k, v := range transformed {
		result[k] = v
	}

	assert.Equal(t, map[string]string{
		"aa": "A",
		"bb": "B",
		"cc": "C",
	}, result, "transformed map should match expected key-value pairs")
}

func TestFilter2(t *testing.T) {
	// Test filtering key-value pairs with even values
	input := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	result := make(map[string]int)

	filtered := Filter2(
		maps.All(input),
		func(k string, v int) bool {
			return v%2 == 0
		},
	)

	for k, v := range filtered {
		result[k] = v
	}

	assert.Equal(t, map[string]int{
		"b": 2,
		"d": 4,
	}, result, "filtered map should contain only pairs with even values")
}
func TestForEach(t *testing.T) {
	// Test iterating over slice elements
	input := []int{1, 2, 3}
	result := make([]int, 0)

	ForEach(
		slices.Values(input),
		func(v int) {
			result = append(result, v)
		},
	)

	assert.Equal(t, []int{1, 2, 3}, result, "result slice should contain all elements in order")
}

func TestForEach2(t *testing.T) {
	// Test iterating over map key-value pairs
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	result := make(map[string]int)

	ForEach2(
		maps.All(input),
		func(k string, v int) {
			result[k] = v + 1
		},
	)

	assert.Equal(t, map[string]int{
		"a": 2,
		"b": 3,
		"c": 4,
	}, result, "result map should contain all key-value pairs")
}
