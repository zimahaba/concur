package cmd

import (
	"concur/pgk"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show settings",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		apis := pgk.Apis{}
		err := viper.UnmarshalKey("apis", &apis)
		if err != nil {
			panic(err)
		}
		fmt.Println(apis.String())
		return nil
	},
}

func init() {
	configCmd.AddCommand(showCmd)
}
