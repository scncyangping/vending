package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "vending",
	Short: "简介",
	Long: `详细的介绍1
				详细的介绍2
					详细的介绍3
			详细的介绍4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd cmd")
	},
}

var child1 = &cobra.Command{
	Use:   "chi1",
	Short: "简介1",
	Long: `chi1详细的介绍1
				chi1详细的介绍2
					chi1详细的介绍3
			chi1详细的介绍4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("child1 cmd")
	},
}

var child2 = &cobra.Command{
	Use:   "chi2",
	Short: "简介2",
	Long: `chi2详细的介绍1
				chi2详细的介绍2
					chi2详细的介绍3
			chi2详细的介绍4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("child2 cmd")
	},
}

func Execute() {
	rootCmd.AddCommand(child1)
	rootCmd.AddCommand(child2)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
