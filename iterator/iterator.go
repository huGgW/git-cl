package iterator

import (
	"iter"
)

type Iterator[T any] interface {
	Next() (T, error)
	Close()
}

func SeqOf[T any](iterator Iterator[T]) iter.Seq[T] {
	return func(yield func(v T) bool) {
		defer iterator.Close()

		for {
			v, err := iterator.Next()
			if err != nil {
				return
			}

			if !yield(v) {
				return
			}
		}
	}
}

func Map[T any, U any](seq iter.Seq[T], mapper func(T) U) iter.Seq[U] {
	return func(yield func(v U) bool) {
		for t := range seq {
			v := mapper(t)
			if !yield(v) {
				return
			}
		}
	}
}

func Map2[K1, V1, K2, V2 any](seq iter.Seq2[K1, V1], mapper func(K1, V1) (K2, V2)) iter.Seq2[K2, V2] {
	return func(yield func(k K2, v V2) bool) {
		for k1, v1 := range seq {
			k2, v2 := mapper(k1, v1)
			if !yield(k2, v2) {
				break
			}
		}
	}
}

func Filter[T any](seq iter.Seq[T], conditions ...func(T) bool) iter.Seq[T] {
	return func(yield func(v T) bool) {
	OUTER:
		for v := range seq {
			for _, cond := range conditions {
				if !cond(v) {
					continue OUTER
				}
			}

			if !yield(v) {
				return
			}
		}
	}
}

func Filter2[K, V any](seq iter.Seq2[K, V], conditions ...func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(k K, v V) bool) {
	OUTER:
		for k, v := range seq {
			for _, cond := range conditions {
				if !cond(k, v) {
					continue OUTER
				}
			}

			if !yield(k, v) {
				return
			}
		}
	}
}
