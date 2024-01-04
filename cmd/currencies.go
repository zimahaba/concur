package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var currenciesCmd = &cobra.Command{
	Use:   "currencies",
	Short: "Print available currencies",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(availableCurrencies)
	},
}

func init() {
	rootCmd.AddCommand(currenciesCmd)
}
