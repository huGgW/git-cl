package viewer

import (
	"context"
	"errors"

	"github.com/huGgW/git-cl/internal/git/filter"
	"github.com/huGgW/git-cl/internal/git/model"
)

var (
	ErrFailedToGetLocalBranches = errors.New("failed to get local branches")
)

type Viewer interface {
	GetLocalBranches(context.Context) (model.LocalBranches, error)
	FilterLocalBranches(ctx context.Context, branches model.LocalBranches, filters ...filter.Filter) []model.Branch
}
