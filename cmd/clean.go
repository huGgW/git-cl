package cmd

import (
	"fmt"

	"github.com/huGgW/git-cl/internal/git/filter"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean branches",
	RunE:  cleanCmdRunE,
}

var (
	blackList []string
	dryRun    bool
)

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVarP(&dryRun, "dry", "d", false, "dry run the command (without actually deleting)")
	cleanCmd.Flags().StringSliceVarP(&blackList, "blacklist", "b", nil, "list of branches to be excluded from deletion")
}

func cleanCmdRunE(cmd *cobra.Command, args []string) error {
	const wrapErrMsg = "failed to clean branches"

	ctx := cmd.Context()

	localBranches, err := deps.viewer.GetLocalBranches(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	branchesToBeDeleted := deps.viewer.FilterLocalBranches(
		ctx,
		localBranches,
		filters()...,
	)

	fmt.Println("Branches to be deleted:")
	for _, branch := range branchesToBeDeleted {
		fmt.Println(branch.Name)
	}
	if dryRun { // Do not actually delete branches if dry run is true
		return nil
	}

	if err := deps.actor.DeleteBranches(ctx, branchesToBeDeleted); err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	return nil
}

func filters() []filter.Filter {
	var filters []filter.Filter

	if len(blackList) > 0 {
		filters = append(filters, filter.BlacklistFilterProvider(blackList))
	}

	return filters
}
