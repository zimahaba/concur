package cmd

import (
	"concur/pgk"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "concur",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
			currencyapi := pgk.Api{Name: "Currency API", Url: "https://api.currencyapi.com/v3/latest?base_currency=%s&currencies=%s", Apikey: ""}
			exchangerateapi := pgk.Api{Name: "Exchange Rate API", Url: "exchangerateapi", Apikey: ""}

			pgk.Apis = pgk.ApiConfig{Active: "currencyapi", Available: map[string]pgk.Api{"currencyapi": currencyapi, "exchangerateapi": exchangerateapi}}

			viper.Set("apis", pgk.Apis)
			viper.SafeWriteConfigAs("config.yaml")
		} else {
			panic(err)
		}
	} else {
		err := viper.UnmarshalKey("apis", &pgk.Apis)
		if err != nil {
			panic(err)
		}
	}
}
