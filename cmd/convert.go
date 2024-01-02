package cmd

import (
	"concur/pgk"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var verbose bool
var cache bool

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a currency in others",
	Long:  ``,
	Args:  validArgs,
	RunE: func(cmd *cobra.Command, args []string) error {

		value, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid value to be converted.")
		}
		baseCurrency := strings.ToUpper(args[1])
		currencies := []string{}
		for i := 2; i < len(args); i++ {
			currencies = append(currencies, strings.ToUpper(args[i]))
		}

		m := make(map[string]pgk.Currency)
		conversion := pgk.Conversion{BaseCurrency: baseCurrency, Data: m}

		if cache {
			pgk.SetCachedCurrencies(&conversion, baseCurrency, currencies)
		}

		notCachedCurrencies := []string{}
		for _, cur := range currencies {
			if conversion.Data[cur].Value == 0 {
				notCachedCurrencies = append(notCachedCurrencies, cur)
			}
		}

		if !cache || len(notCachedCurrencies) > 0 {

			currencies := strings.Builder{}
			for i := 2; i < len(args); i++ {
				currencies.WriteString(strings.ToUpper(args[i]))
				if i < len(args)-1 {
					currencies.WriteString("%2C")
				}
			}

			url := fmt.Sprintf("https://api.currencyapi.com/v3/latest?base_currency=%s&currencies=%s", conversion.BaseCurrency, currencies.String())
			client := http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("apikey", "cur_live_Y48ZlyKd8vtKY3nYeTGQYyJqNSbeUXE0jpTvE21d")
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("could not connect to currency api.")
			}
			defer resp.Body.Close()

			err = json.NewDecoder(resp.Body).Decode(&conversion)
			if err != nil {
				return fmt.Errorf("could not read currency response body.")
			}

			pgk.Upsert(conversion)
		}

		conversionMap := make(map[string]float32)
		for k, v := range conversion.Data {
			conversionMap[k] = v.Value * float32(value)
		}
		pgk.Print(baseCurrency, float32(value), conversionMap, verbose)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	convertCmd.Flags().BoolVarP(&cache, "cache", "c", false, "skip request and use cached conversion rates")
}

func validArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(3)(cmd, args); err != nil {
		return err
	}

	for _, c := range args[1:] {
		if slices.Contains(availableCurrencies, strings.ToUpper(c)) == false {
			return fmt.Errorf("invalid currency '%s' - available: %v", c, availableCurrencies)
		}
	}

	return nil
}

var availableCurrencies = []string{"BWP", "DJF", "SGD", "THB", "BBD", "NOK", "SLL", "IQD", "MGA", "MKD", "SDG", "TMT", "TZS", "UYU", "VND", "BSD", "ARB", "OP", "ZAR", "BDT", "KMF", "KRW", "MNT", "RSD", "RWF", "UGX", "AOA", "UZS", "LVL", "DZD", "FKP", "KGS", "TRY", "LTC", "ADA", "BUSD", "ANG", "KHR", "VEF", "XCD", "CVE", "COP", "WST", "ETH", "BNB", "BAM", "EGP", "XPD", "AWG", "HTG", "LBP", "MAD", "PGK", "XAU", "GYD", "CDF", "CUP", "PKR", "PLN", "RON", "SVC", "TND", "BRL", "BTC", "XDR", "IRR", "KPW", "LKR", "UAH", "AMD", "HKD", "KWD", "LRD", "MMK", "DKK", "SCR", "XPT", "SOL", "BOB", "MATIC", "CUC", "BND", "CAD", "GMD", "GNF", "HUF", "MRO", "MZN", "BHD", "SAR", "SBD", "AVAX", "PAB", "JOD", "ALL", "XOF", "CNY", "IDR", "ILS", "KYD", "OMR", "BGN", "CLF", "DOP", "GHS", "NGN", "BZD", "GIP", "MXN", "NPR", "PHP", "PYG", "TWD", "XRP", "ERN", "CHF", "CRC", "GEL", "ISK", "JMD", "SRD", "STD", "BYN", "XAF", "BMD", "BTN", "CLP", "JEP", "NAD", "VUV", "AZN", "INR", "JPY", "MUR", "MWK", "TJS", "TOP", "AUD", "HNL", "NZD", "BIF", "MOP", "USDT", "DAI", "AED", "HRK", "IMP", "SYP", "USD", "YER", "USDC", "FJD", "BYR", "SHP", "SZL", "ARS", "EUR", "GBP", "MVR", "PEN", "QAR", "SOS", "ZWL", "AFN", "KZT", "NIO", "SEK", "XPF", "ZMK", "DOT", "GTQ", "LAK", "LTL", "MDL", "KES", "ETB", "GGP", "LSL", "LYD", "MYR", "RUB", "TTD", "CZK", "ZMW", "XAG"}
