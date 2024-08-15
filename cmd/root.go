package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "mcenter-api",
	Short: "mcenter scaffold",
	Long:  "server scaffold written by qiaogy",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.AddCommand(startCmd)
}
