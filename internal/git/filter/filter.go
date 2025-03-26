package filter

import (
	"github.com/huGgW/git-cl/internal/git/model"
	"github.com/huGgW/git-cl/pkg/set"
)

type Filter func(model.Branch) bool
type FilterProvider func([]string) Filter

var CurrentExcludeFilter Filter = func(branch model.Branch) bool {
	return !branch.IsCurrent
}

func BlacklistFilterProvider(blacklist []string) Filter {
	blacklistSet := set.SetOf(blacklist...)
	return func(branch model.Branch) bool {
		return !blacklistSet.Exists(branch.Name)
	}
}
