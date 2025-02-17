package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean branches",
	RunE:  cleanCmdRunE,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func cleanCmdRunE(cmd *cobra.Command, args []string) error {
	const wrapErrMsg = "failed to clean branches"

	ctx := cmd.Context()

	localBranches, err := deps.viewer.GetLocalBranches(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	branches := deps.viewer.FilterLocalBranches(ctx, localBranches)
	if err := deps.actor.DeleteBranches(ctx, branches); err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	return nil
}
