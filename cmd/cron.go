// Package cmd /*
package cmd

import (
	"github.com/pro911/golang-demo/internal/pkg/crontab"
	"github.com/spf13/cobra"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "单独调用 定时任务服务的。执行后常驻,会定时处理定时任务",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: handle,
}

func init() {
	rootCmd.AddCommand(cronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handle(cmd *cobra.Command, args []string) {
	crontab.NewCronTab()
}
