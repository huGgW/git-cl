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

	branches, err := deps.viewer.GetLocalBranches(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", wrapErrMsg, err)
	}

	for _, branch := range branches.All {
		if branches.Current != nil && *branches.Current == branch {
			fmt.Printf("%s <- current\n", branch)
		} else {
			fmt.Println(branch)
		}
	}
	return nil
}
