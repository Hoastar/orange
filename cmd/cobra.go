/*
@Time : 2020/10/28 下午10:01
@Author : hoastar
@File : cobra
@Software: GoLand
*/

package cmd

import (
	_ "errors"
	"github.com/hoastar/orange/cmd/api"
	"github.com/hoastar/orange/cmd/migrate"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "orange",
	Short: "-v",
	SilenceUsage: true,
	DisableAutoGenTag: true,
	Long: "orange",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},

	PersistentPostRunE: func(*cobra.Command, []string) error { return nil},
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := "欢迎使用 ferry，可以使用 -h 查看命令"
		logger.Info("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}


// Excute: apply commands
func Excute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}





























