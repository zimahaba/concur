package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the app version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("V1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
