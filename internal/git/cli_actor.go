package git

import (
	"context"
	"fmt"
	"os/exec"
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
