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
	cfg  config.Config
	deps struct {
		viewer git.Viewer
		actor  git.Actor
	}
)

var rootCmd = &cobra.Command{
	Use:               "gitcl",
	Short:             "gitcl is a tool for cleaning up local branches of git repository",
	ValidArgs:         []string{},
	PersistentPreRunE: persistentPreRunE,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cfg.Fetch, "fetch", "f", true, "set whether to fetch before cleaning up branches")
	deps.viewer = git.NewGogitViewer()
	deps.actor = git.NewCliActor()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	if cfg.Fetch {
		if err := deps.actor.FetchAll(context.Background()); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
