package git

import (
	"context"
	"fmt"
	"os/exec"
)

func FetchAll(ctx context.Context) error {
	const wrapErrMsg = "failed to fetch repository"
	cmd := exec.CommandContext(ctx, "git", "fetch", "--all")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	return nil
}
