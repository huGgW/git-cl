package actor

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/huGgW/git-cl/internal/git/model"
)

type cliActor struct {
}

func NewCliActor() *cliActor {
	return &cliActor{}
}

func (g *cliActor) FetchAll(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "git", "fetch", "--all")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	return nil
}

func (g *cliActor) DeleteBranches(ctx context.Context, branches []model.Branch) error {
	if len(branches) == 0 {
		return nil
	}

	args := []string{"branch", "-D"}
	for _, branch := range branches {
		args = append(args, branch.Name)
	}

	cmd := exec.CommandContext(ctx, "git", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailedToDeleteBranches, err)
	}

	return nil
}
