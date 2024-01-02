package cmd

import (
	"concur/pkg"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var api string
var key string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure settings",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		apis := pkg.ApiConfig{}
		err := viper.UnmarshalKey("apis", &apis)
		if err != nil {
			panic(err)
		}
		if api != "" {
			if _, ok := apis.Available[api]; ok {
				apis.Active = api
			} else {
				return fmt.Errorf("invalid api.")
			}
		}

		if key != "" {
			active := viper.GetString("apis.active")
			api := apis.Available[active]
			api.Apikey = key
			apis.Available[active] = api
		}

		viper.Set("apis", apis)
		viper.WriteConfig()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&api, "api", "a", "", "currency api used")
	configCmd.Flags().StringVarP(&key, "key", "k", "", "api key for the currency api")
}
