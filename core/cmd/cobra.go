package cmd

import (
	"errors"
	"go-admin/core/cmd/api"
	"go-admin/core/config"
	"go-admin/core/utils/log"
	"go-admin/core/utils/textutils"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          config.ApplicationConfig.Name,
	Short:        config.ApplicationConfig.Name,
	SilenceUsage: true,
	Long:         config.ApplicationConfig.Name,
	Version:      config.ApplicationConfig.Version,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(textutils.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + textutils.Green(config.ApplicationConfig.Name+" "+config.ApplicationConfig.Version) + ` 可以使用 ` + textutils.Red(`-h`) + ` 查看命令`
	log.Infof("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
