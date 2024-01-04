package cmd

import (
	"concur/pkg"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "concur",
	Short: "Concur is a command line currency converter",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			currencyapi := pkg.Api{Name: "Currency API", Url: "https://api.currencyapi.com/v3/latest?base_currency=%s&currencies=%s", Apikey: ""}
			exchangerateapi := pkg.Api{Name: "Exchange Rate API", Url: "exchangerateapi", Apikey: ""}

			pkg.Apis = pkg.ApiConfig{Active: "currencyapi", Available: map[string]pkg.Api{"currencyapi": currencyapi, "exchangerateapi": exchangerateapi}}

			viper.Set("apis", pkg.Apis)
			viper.SafeWriteConfigAs("config.yaml")
		} else {
			panic(err)
		}
	} else {
		err := viper.UnmarshalKey("apis", &pkg.Apis)
		if err != nil {
			panic(err)
		}
	}
}
