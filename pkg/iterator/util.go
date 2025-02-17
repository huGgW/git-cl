package iterator

import "iter"

func First[K, V any](it iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range it {
			if !yield(k) {
				return
			}
		}
	}
}

func Second[K, V any](it iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range it {
			if !yield(v) {
				return
			}
		}
	}
}

func Map[V1, V2 any](it iter.Seq[V1], mapper func(V1) V2) iter.Seq[V2] {
	return func(yield func(s V2) bool) {
		for t := range it {
			s := mapper(t)
			if !yield(s) {
				return
			}
		}
	}
}

func Map2[K1, V1, K2, V2 any](it iter.Seq2[K1, V1], mapper func(K1, V1) (K2, V2)) iter.Seq2[K2, V2] {
	return func(yield func(k2 K2, v2 V2) bool) {
		for k1, v1 := range it {
			k2, v2 := mapper(k1, v1)
			if !yield(k2, v2) {
				return
			}
		}
	}
}

func Filter[V any](it iter.Seq[V], cond func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range it {
			if !cond(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

func Filter2[K, V any](it iter.Seq2[K, V], cond func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range it {
			if !cond(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

func ForEach[V any](it iter.Seq[V], fn func(V)) {
	it(func(v V) bool {
		fn(v)
		return true
	})
}

func ForEach2[K, V any](it iter.Seq2[K, V], fn func(K, V)) {
	it(func(k K, v V) bool {
		fn(k, v)
		return true
	})
}
