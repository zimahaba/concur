package pgk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var Apis ApiConfig

type ApiClient interface {
	SetRemoteCurrencies(conversion *Conversion, currencies string) error
}

func GetApiClient() (ApiClient, error) {
	if Apis.Active == "currencyapi" {
		return CurrencyAPI{}, nil
	} else if Apis.Active == "exchangerateapi" {
		return ExchangeRateAPI{}, nil
	}
	return nil, fmt.Errorf("unable to get api client.")
}

type CurrencyAPI struct{}

func (c CurrencyAPI) SetRemoteCurrencies(conversion *Conversion, currencies string) error {
	url := fmt.Sprintf(Apis.Available[Apis.Active].Url, conversion.BaseCurrency, currencies)
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", Apis.Available[Apis.Active].Apikey)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not connect to currency api.")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&conversion)
	if err != nil {
		return fmt.Errorf("could not read currency response body.")
	}

	return nil
}

type ExchangeRateAPI struct{}

func (c ExchangeRateAPI) SetRemoteCurrencies(conversion *Conversion, currencies string) error {
	conversion.BaseCurrency = "USD"
	conversion.Data = map[string]Currency{"BRL": {Code: "BRL", Value: 4.354168}}
	return nil
}
