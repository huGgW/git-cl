package scanner

import (
	"bufio"
	"iter"
)

func LineSeq(sc *bufio.Scanner) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		for sc.Scan() {
			line := sc.Text()
			if !yield(line, nil) {
				return
			}
		}

		if err := sc.Err(); err != nil {
			yield("", err)
		}
	}
}
