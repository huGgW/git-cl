package git

import (
	"context"
	"errors"
)

var (
	ErrFailedToFetch          = errors.New("failed to fetch repository")
	ErrFailedToDeleteBranches = errors.New("failed to delete branches")
)

type Actor interface {
	FetchAll(ctx context.Context) error
	DeleteBranches(ctx context.Context, branches []Branch) error
}
