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
}

type LocalBranches struct {
	All     []string
	Current *string
}
