package git

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/huGgW/git-cl/pkg/iterator"
	"github.com/huGgW/git-cl/pkg/iterator/scanner"
)

type LocalBranches struct {
	Current string
	All     []string
}

func parseLocalBranches(reader io.Reader) (LocalBranches, error) {
	const wrapErrMsg = "failed to parse local branches"

	branches := LocalBranches{}
	sc := bufio.NewScanner(reader)
	scSeqErr := iterator.Map2(
		scanner.LineSeq(sc),
		func(line string, err error) (string, error) {
			return strings.TrimSpace(line), err
		},
	)

	for line, err := range scSeqErr {
		if err != nil {
			return LocalBranches{}, fmt.Errorf("%s: %w", wrapErrMsg, err)
		}

		if strings.HasPrefix(line, "* ") {
			line = strings.TrimPrefix(line, "* ")
			branches.Current = line
		}

		branches.All = append(branches.All, line)
	}

	return branches, nil
}

func FetchAll(ctx context.Context) error {
	const wrapErrMsg = "failed to fetch repository"
	cmd := exec.CommandContext(ctx, "git", "fetch", "--all")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	return nil
}

func GetLocalBranches(ctx context.Context) (LocalBranches, error) {
	const wrapErrMsg = "failed to get local branches"

	cmd := exec.CommandContext(ctx, "git", "branch", "--list")
	out, err := cmd.Output()
	if err != nil {
		return LocalBranches{}, fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	localBranches, err := parseLocalBranches(bytes.NewReader(out))
	if err != nil {
		return LocalBranches{}, fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	return localBranches, nil
}
