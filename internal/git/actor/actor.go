package actor

import (
	"context"
	"errors"

	"github.com/huGgW/git-cl/internal/git/model"
)

var (
	ErrFailedToFetch          = errors.New("failed to fetch repository")
	ErrFailedToDeleteBranches = errors.New("failed to delete branches")
)

type Actor interface {
	FetchAll(ctx context.Context) error
	DeleteBranches(ctx context.Context, branches []model.Branch) error
}
