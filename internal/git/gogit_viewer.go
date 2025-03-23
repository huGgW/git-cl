package git

import (
	"context"
	"errors"
	"fmt"
	"io"
	"iter"
	"slices"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/huGgW/git-cl/pkg/iterator"
)

var (
	ErrFailedToOpenRepo = errors.New("failed to open repository")
	ErrFailedToGetHead  = errors.New("failed to get head")
)

type gogitViewer struct {
}

func NewGogitViewer() *gogitViewer {
	return &gogitViewer{}
}

func (g *gogitViewer) GetLocalBranches(ctx context.Context) (LocalBranches, error) {
	var localBranches LocalBranches

	repo, err := g.openRepo()
	if err != nil {
		return LocalBranches{}, fmt.Errorf("%w: %w", ErrFailedToGetLocalBranches, err)
	}

	head, err := g.getHead(repo)
	if err != nil {
		return LocalBranches{}, fmt.Errorf("%w: %w", ErrFailedToGetLocalBranches, err)
	}

	refIter, err := repo.Branches()
	if err != nil {
		return LocalBranches{}, fmt.Errorf("%w: %w", ErrFailedToGetLocalBranches, err)
	}
	for branchRef, err := range referenceIterToSeq(refIter) {
		if err != nil {
			return LocalBranches{}, fmt.Errorf("%w: %w", ErrFailedToGetLocalBranches, err)
		}

		branchName := branchRef.Name().Short()
		branch := Branch{Name: branchName}

		if eq, err := g.refEqual(head, branchRef); err != nil {
			return LocalBranches{}, fmt.Errorf("%w: %w", ErrFailedToGetLocalBranches, err)
		} else if eq {
			branch.IsCurrent = true
		}

		localBranches.Branches = append(localBranches.Branches, branch)
	}

	return localBranches, nil
}

func (g *gogitViewer) FilterLocalBranches(ctx context.Context, branches LocalBranches, filters ...Filter) []Branch {
	branchSeq := slices.Values(branches.Branches)

	// add current exclude filter because branch deletion cannot delete the current branch
	branchSeq = iterator.Filter(branchSeq, currentExcludeFilter)

	for _, filter := range filters {
		branchSeq = iterator.Filter(branchSeq, filter)
	}

	return slices.Collect(branchSeq)
}

func (g *gogitViewer) openRepo() (*gogit.Repository, error) {
	repo, err := gogit.PlainOpen(".") // TODO: Make path configurable
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToOpenRepo, err)
	}

	return repo, nil
}

func (g *gogitViewer) getHead(repo *gogit.Repository) (*plumbing.Reference, error) {
	head, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToGetHead, err)
	}

	return head, nil
}

func (g *gogitViewer) refEqual(r1, r2 *plumbing.Reference) (bool, error) {
	return r1.Hash() == r2.Hash(), nil
}

func referenceIterToSeq(iter storer.ReferenceIter) iter.Seq2[*plumbing.Reference, error] {
	return func(yield func(*plumbing.Reference, error) bool) {
		defer iter.Close()

		for {
			ref, err := iter.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}

				yield(nil, err)
				return
			}

			if !yield(ref, nil) {
				return
			}
		}
	}
}
