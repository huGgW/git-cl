package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dryCmd = &cobra.Command{
	Use:   "dry",
	Short: "dry run",
	RunE:  dryCmdRunE,
}

func init() {
	rootCmd.AddCommand(dryCmd)
}

func dryCmdRunE(cmd *cobra.Command, args []string) error {
	const wrapErrMsg = "failed to dry run"

	ctx := cmd.Context()
	localBranches, err := deps.viewer.GetLocalBranches(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	// TODO: add filters
	branches := deps.viewer.FilterLocalBranches(ctx, localBranches)
	fmt.Println("Branches to be deleted")
	for _, branch := range branches {
		fmt.Println(branch.Name)
	}

	return nil
}
