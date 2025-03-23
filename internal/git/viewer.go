package git

import (
	"context"
	"errors"

	"github.com/huGgW/git-cl/pkg/set"
)

var (
	ErrFailedToGetLocalBranches = errors.New("failed to get local branches")
)

type Viewer interface {
	GetLocalBranches(context.Context) (LocalBranches, error)
	FilterLocalBranches(ctx context.Context, branches LocalBranches, filters ...Filter) []Branch
}

type LocalBranches struct {
	Branches []Branch
}

type Branch struct {
	Name      string
	IsCurrent bool
}

type Filter func(Branch) bool
type FilterProvider func([]string) Filter

var currentExcludeFilter Filter = func(branch Branch) bool {
	return !branch.IsCurrent
}

func BlacklistFilterProvider(blacklist []string) Filter {
	blacklistSet := set.SetOf(blacklist...)
	return func(branch Branch) bool {
		return !blacklistSet.Exists(branch.Name)
	}
}
