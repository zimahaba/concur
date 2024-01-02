package cmd

import (
	"concur/pgk"
	"fmt"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show settings",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(pgk.Apis.String())
		return nil
	},
}

func init() {
	configCmd.AddCommand(showCmd)
}
