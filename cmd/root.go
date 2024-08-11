package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "git-cl",
	Short:   "git-cl is simple yet fast git branch cleaner",
	Long:    "git-cl is simple yet fast git branch cleaner",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Sweap~!")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
