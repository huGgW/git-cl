package iterator

import (
	"errors"
	"io"
	"iter"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func ReferenceIterToSeq(iter storer.ReferenceIter) iter.Seq2[*plumbing.Reference, error] {
	return func(yield func(*plumbing.Reference, error) bool) {
		defer iter.Close()

		for {
			ref, err := iter.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}

				yield(nil, err)
				return
			}

			if !yield(ref, nil) {
				return
			}
		}
	}
}
