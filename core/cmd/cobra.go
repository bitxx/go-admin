package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"go-admin/core/cmd/api"
	"go-admin/core/utils/textutils"
)

var rootCmd = &cobra.Command{
	SilenceUsage: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(textutils.Red("requires at least one arg"))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
