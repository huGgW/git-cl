package set

import (
	"iter"
	"unique"
)

type Set[T comparable] map[unique.Handle[T]]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T](map[unique.Handle[T]]struct{}{})
}

func SetOf[T comparable](values ...T) Set[T] {
	s := NewSet[T]()
	for _, value := range values {
		s.Add(value)
	}
	return s
}

func SetFrom[T comparable](valueSeq iter.Seq[T]) Set[T] {
	s := NewSet[T]()
	for value := range valueSeq {
		s.Add(value)
	}
	return s
}

func (s Set[T]) Exists(value T) bool {
	_, ok := s[unique.Make(value)]
	return ok
}

func (s Set[T]) Add(value T) {
	s[unique.Make(value)] = struct{}{}
}

func (s Set[T]) Remove(value T) {
	delete(s, unique.Make(value))
}
