package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list branches",
	RunE:    listCmdRunE,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listCmdRunE(cmd *cobra.Command, args []string) error {
	const wrapErrMsg = "failed to list branches"

	ctx := cmd.Context()

	localBranches, err := deps.viewer.GetLocalBranches(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	for _, branch := range localBranches.Branches {
		if branch.IsCurrent {
			fmt.Printf("%s <- current\n", branch.Name)
		} else {
			fmt.Println(branch.Name)
		}
	}
	return nil
}
