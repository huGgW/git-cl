package git

import (
	"context"
	"errors"
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
