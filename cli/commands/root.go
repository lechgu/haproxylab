package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "haclient",
	Short: "haclient",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute ...
func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
