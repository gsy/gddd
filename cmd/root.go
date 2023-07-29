package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gddd",
	Short: "gddd is a scaffold for ddd project",
	Long:  `gddd is a scaffold for ddd project`,
	Run:   runHelp,
}

// func init() {
//	rootCmd.AddCommand(helpCmd)
// }

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func Execute() error {
	return rootCmd.Execute()
}
