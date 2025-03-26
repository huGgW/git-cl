package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/huGgW/git-cl/internal/config"
	"github.com/huGgW/git-cl/internal/git/actor"
	"github.com/huGgW/git-cl/internal/git/viewer"
	"github.com/spf13/cobra"
)

var (
	cfg  config.Config
	deps struct {
		viewer viewer.Viewer
		actor  actor.Actor
	}
)

var rootCmd = &cobra.Command{
	Use:               "git-cl",
	Short:             "git-cl is a tool for cleaning up local branches of git repository",
	ValidArgs:         []string{},
	PersistentPreRunE: persistentPreRunE,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cfg.Fetch, "fetch", "f", false, "set whether to fetch before cleaning up branches")
	deps.viewer = viewer.NewGogitViewer()
	deps.actor = actor.NewCliActor()
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
