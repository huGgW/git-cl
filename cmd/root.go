package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/huGgW/git-cl/internal/config"
	"github.com/huGgW/git-cl/internal/git"
	"github.com/spf13/cobra"
)

var (
	cfg config.Config
)

var rootCmd = &cobra.Command{
	Use:               "gitcl",
	Short:             "gitcl is a tool for cleaning up local branches of git repository",
	ValidArgs:         []string{},
	PersistentPreRunE: persistentPreRunE,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cfg.Fetch, "fetch", "f", true, "set whether to fetch before cleaning up branches")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	if cfg.Fetch {
		if err := git.FetchAll(context.Background()); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
