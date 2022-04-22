package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vending/app/types/constants"
)

var rootCmd = &cobra.Command{
	Short: "template for test",
	Long: `命令行工具，现在仅提供对程序启动的操作，后续
可添加对内存、日志等操作。`,
	CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
}
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "启动程序",
	Long:  `run 启动程序`,
	Args:  cobra.ExactArgs(1),
}

var runProCmd = &cobra.Command{
	Use:   "pro",
	Short: "run pro",
	Long:  `设置当前运行环境为正式环境`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		run(constants.ReleaseMode)
	},
}

var runDevCmd = &cobra.Command{
	Use:   "dev",
	Short: "run dev",
	Long:  `设置当前运行环境为开发环境`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		run(constants.DebugMode)
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止程序",
	Long:  ` stop 停止程序`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func Execute() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(stopCmd)
	runCmd.AddCommand(runProCmd)
	runCmd.AddCommand(runDevCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
