package git

import (
	"context"
	"errors"
)

var (
	ErrFailedToFetch = errors.New("failed to fetch repository")
)

type Actor interface {
	FetchAll(ctx context.Context) error
}
